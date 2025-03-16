import json
import random
import numpy as np
import nltk
import pickle
import tensorflow as tf
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from nltk.stem import WordNetLemmatizer
from tensorflow.keras.models import Sequential, load_model
from tensorflow.keras.layers import Dense, Dropout

nltk.download("punkt", quiet=True)
nltk.download("wordnet", quiet=True)

app = FastAPI()

with open("intents.json", "r", encoding="utf-8") as file:
    data = json.load(file)

lemmatizer = WordNetLemmatizer()
ignore_letters = ["?", "!", ".", ","]

words, classes, documents = [], [], []
MODEL_FILE = "chatbot_model.h5"
WORDS_FILE = "words.pkl"
CLASSES_FILE = "classes.pkl"

def prepare_and_train_model():
    global words, classes, documents
    for intent in data["intents"]:
        for pattern in intent["patterns"]:
            word_list = nltk.word_tokenize(pattern)
            words.extend(word_list)
            documents.append((word_list, intent["tag"]))
            if intent["tag"] not in classes:
                classes.append(intent["tag"])

    words = sorted(set(lemmatizer.lemmatize(w.lower()) for w in words if w not in ignore_letters))
    classes = sorted(set(classes))

    training = []
    output_empty = [0] * len(classes)

    for doc in documents:
        bag = [1 if w in [lemmatizer.lemmatize(w.lower()) for w in doc[0]] else 0 for w in words]
        output_row = list(output_empty)
        output_row[classes.index(doc[1])] = 1
        training.append([bag, output_row])

    random.shuffle(training)
    training = np.array(training, dtype=object)
    train_x, train_y = np.array(list(training[:, 0])), np.array(list(training[:, 1]))

    model = Sequential([
        Dense(128, activation="relu", input_shape=(len(train_x[0]),)),
        Dropout(0.5),
        Dense(64, activation="relu"),
        Dropout(0.5),
        Dense(len(classes), activation="softmax")
    ])

    model.compile(loss="categorical_crossentropy", optimizer="adam", metrics=["accuracy"])
    model.fit(train_x, train_y, epochs=200, batch_size=5, verbose=1)

    model.save(MODEL_FILE)
    with open(WORDS_FILE, "wb") as f:
        pickle.dump(words, f)
    with open(CLASSES_FILE, "wb") as f:
        pickle.dump(classes, f)
    return model

try:
    model = load_model(MODEL_FILE)
    with open(WORDS_FILE, "rb") as f:
        words = pickle.load(f)
    with open(CLASSES_FILE, "rb") as f:
        classes = pickle.load(f)
    print("Model loaded successfully!")
except:
    print("Training new model...")
    model = prepare_and_train_model()

def clean_sentence(sentence):
    return [lemmatizer.lemmatize(word.lower()) for word in nltk.word_tokenize(sentence)]

def bag_of_words(sentence):
    sentence_words = clean_sentence(sentence)
    return np.array([1 if word in sentence_words else 0 for word in words])

def predict_class(sentence):
    bow = bag_of_words(sentence)
    res = model.predict(np.array([bow]))[0]
    return classes[np.argmax(res)] if np.max(res) > 0.6 else "fallback"

def chatbot_response(msg):
    tag = predict_class(msg)
    for intent in data["intents"]:
        if intent["tag"] == tag:
            return random.choice(intent["responses"])
    return "I'm not sure how to respond to that."

class ChatMessage(BaseModel):
    text: str

@app.post("/chatbot/")
async def chatbot(message: ChatMessage):
    try:
        user_message = message.text
        response = chatbot_response(user_message)
        return {"response": response}
    except Exception as e:
        print(f"Endpoint error: {e}")
        raise HTTPException(status_code=500, detail="Internal server error")

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)