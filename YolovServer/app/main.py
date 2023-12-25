from fastapi import Request, FastAPI

from app.model.model import ModelSchema

app = FastAPI()


unique_model = ModelSchema()


@app.post("/operation/{value}")
async def set_operation(model_operation: str):
    unique_model.operation = model_operation
    return {"message": model_operation}


@app.post("/type/{value}")
async def set_yolov8(model_yolov8: str):
    unique_model.yolov8 = model_yolov8
    return {"message": model_yolov8}


@app.post("/url")
async def set_dataset(request: Request):
    url = request.json()["url"]
    unique_model.dataset_url = url
    return {"message": url}


@app.get("/")
async def set_dataset():
    return {"message": unique_model}


@app.get("/check")
async def set_dataset():
    print("CHECK")
    return {"message": ""}


