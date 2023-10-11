from fastapi import FastAPI
from pydantic import BaseModel
from process_data import process


class CSV(BaseModel):
    content: str

app = FastAPI()

@app.get("/")
async def index():
    return {"message": "hello world"} 


@app.post("/process_data")
async def process_data(csv: CSV):
    inbounds, outbounds, inbounds_answered, outbounds_answered = process(csv.content)
    return {
        "inbounds": inbounds,
        "outbounds": outbounds,
        "inbounds_answered": inbounds_answered,
        "outbounds_answered": outbounds_answered
    }
