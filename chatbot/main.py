import json
import random
import requests
from fastapi import FastAPI
from pydantic import BaseModel
from transformers import pipeline
from langdetect import detect
from googletrans import Translator

app = FastAPI()

class Message(BaseModel):
    message: str
    session_id: str = "default"


with open("intents.json", "r") as f:
    intents_data = json.load(f)

responses = {}
candidate_labels = []
for intent in intents_data["intents"]:
    tag = intent["tag"]
    candidate_labels.append(tag)
    responses[tag] = intent["responses"]

#user sesh storage
user_sessions = {}

@app.post("/chatbot")
async def chatbot_endpoint(message: Message):
    bot_response = get_response(message.message, message.session_id)
    return {"response": bot_response}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)
