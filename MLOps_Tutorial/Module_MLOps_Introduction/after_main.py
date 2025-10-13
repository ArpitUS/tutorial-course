import mlflow
import mlflow.sklearn
from fastapi import FastAPI
from pydantic import BaseModel
from sklearn.linear_model import LinearRegression
import numpy as np

# Train simple model
X = np.array([[1], [2], [3], [4]])
y = np.array([2, 4, 6, 8])
model = LinearRegression()
model.fit(X, y)

# Log model
mlflow.sklearn.log_model(model, "linear_model")
mlflow.log_param("input_feature", "x")
mlflow.log_metric("r2_score", model.score(X, y))

# FastAPI app
app = FastAPI()

class InputData(BaseModel):
    x: float

@app.post("/predict")
def predict(data: InputData):
    result = model.predict([[data.x]])[0]
    return {"input": data.x, "prediction": float(result)}
