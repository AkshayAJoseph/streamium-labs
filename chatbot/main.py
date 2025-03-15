import json
import random
import requests
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
#import fasttext
from googletrans import Translator
from transformers import pipeline
import os

app = FastAPI()

class Message(BaseModel):
    message: str
    session_id: str = "default"

classifier = pipeline("zero-shot-classification", model="facebook/bart-large-mnli")
translator = Translator()

#lang_detector = fasttext.load_model("lid.176.bin")  

file_path = os.path.join(os.path.dirname(__file__), "intents.json")
with open(file_path, "r", encoding="utf-8") as f:
    try:
        intents_data = json.load(f)
    except json.JSONDecodeError as e:
        print(f"JSON Decode Error: {e}")
        intents_data = {"intents": []}

responses = {}
candidate_labels = []
for intent in intents_data["intents"]:
    tag = intent["tag"]
    candidate_labels.append(tag)
    responses[tag] = intent["responses"]

user_sessions = {}

def fetch_esports_data():
    try:
        mock_data = {"match_name": "Championship Finals", "time": "3 PM", "team1": "Team A", "team2": "Team B"}
        return f"Upcoming match: {mock_data['match_name']} at {mock_data['time']} between {mock_data['team1']} and {mock_data['team2']}."
    except Exception as e:
        print(f"Esports API error: {e}")
        return "Sorry, I couldn't retrieve esports match details at the moment."

def fetch_job_status():
    try:
        mock_data = {"job_title": "Software Engineer", "status": "under review"}
        return f"Your job application for {mock_data['job_title']} is currently {mock_data['status']}."
    except Exception as e:
        print(f"Jobs API error: {e}")
        return "Sorry, I couldn't retrieve job application status at the moment."

def get_response(user_message: str, session_id: str) -> str:
    """Process user input and generate chatbot response."""
    if not user_message.strip():
        return "Please say something so I can assist you!"

    try:
        lang_prediction = lang_detector.predict(user_message.replace('\n', ' '))
        lang = lang_prediction[0][0].replace('__label__', '')  
        confidence = lang_prediction[1][0]
        print(f"Detected language: {lang} (confidence: {confidence:.2f})")
        if confidence < 0.9:  
            print("Low confidence in language detection, defaulting to English")
            lang = "en"
    except Exception as e:
        print(f"Language detection error: {e}")
        lang = "en"  

    try:
        user_message_en = translator.translate(user_message, dest="en").text if lang != "en" else user_message
    except Exception as e:
        print(f"Translation error: {e}")
        user_message_en = user_message

    try:
        result = classifier(user_message_en, candidate_labels)
        tag = result["labels"][0]
        score = result["scores"][0]
    except Exception as e:
        print(f"Classifier error: {e}")
        return "Sorry, something went wrong on my end. Please try again."

    if score < 0.5:
        response_text = random.choice(responses.get("fallback", ["I'm sorry, I didn't quite understand that. Could you please rephrase?"]))
    else:
        if tag == "esports":
            response_text = fetch_esports_data()
        elif tag == "jobs":
            response_text = fetch_job_status()
        else:
            response_text = random.choice(responses[tag])

    if session_id not in user_sessions:
        user_sessions[session_id] = []
    user_sessions[session_id].append({"user": user_message, "bot": response_text})

    try:
        #if lang != "en":
            #return translator.translate(response_text, dest=lang).text
        return response_text  
    except Exception as e:
        print(f"Response translation error: {e}")
        return response_text

@app.post("/chatbot")
async def chatbot_endpoint(message: Message):
    try:
        bot_response = get_response(message.message, message.session_id)
        return {"response": bot_response}
    except Exception as e:
        print(f"Endpoint error: {e}")
        raise HTTPException(status_code=500, detail="Internal server error")

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)