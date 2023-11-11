from fastapi import FastAPI
from pydantic import BaseModel
from starlette.middleware.cors import CORSMiddleware

from predict_model.model import get_blocks, prepare_data, train_model, predict_next_block

BLOCK_CHAIN_URL = "http://localhost:8000"

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=[BLOCK_CHAIN_URL],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


class Block(BaseModel):
    transactionType: str
    predictionPercent: float


class Response(BaseModel):
    prediction: list[Block]


@app.get("/predict")
async def get_prediction(amount: int) -> Response:
    blocks = await get_blocks(f"{BLOCK_CHAIN_URL}/get-blocks", amount)

    # Генерация предсказаний
    df = prepare_data(blocks)
    model = train_model(df)
    prediction = predict_next_block(model, df)

    # Подсчет процентного распределения типов транзакций
    prediction_percentages = {df.columns[i]: prediction[0][i] / 10 for i in range(len(df.columns))}

    # Форматирование предсказаний
    prediction_blocks = []
    for transaction_type, percentage in prediction_percentages.items():
        prediction_blocks.append(Block(transactionType=transaction_type, predictionPercent=percentage))
    return Response(prediction=prediction_blocks)


if __name__ == '__main__':
    import uvicorn

    uvicorn.run(app, host='localhost', port=8001)
