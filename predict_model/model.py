from collections import Counter

import pandas as pd
from httpx import AsyncClient
from sklearn.ensemble import RandomForestClassifier
from sklearn.model_selection import train_test_split


async def get_blocks(blocks_url: str, block_count: int) -> list[list[dict]]:
    """ Получение блоков из API """
    async with AsyncClient() as client:
        response = await client.get(blocks_url, params={'count': block_count})
        response.raise_for_status()
        return response.json()


def prepare_data(blockchain_blocks: list[list[dict]]) -> pd.DataFrame:
    """ Подготовка данных для модели """
    block_counts = [Counter([transaction['transactionType'] for transaction in block]) for block in blockchain_blocks]
    return pd.DataFrame(block_counts).fillna(0)


def train_model(data_frame: pd.DataFrame) -> RandomForestClassifier:
    """ Обучение модели """
    x = data_frame
    y = data_frame.shift(-1).ffill()  # Сдвигаем данные для создания целевой переменной
    x_train, x_test, y_train, y_test = train_test_split(x, y, test_size=0.2, random_state=42)

    classifier = RandomForestClassifier(n_estimators=100, random_state=42)
    classifier.fit(x_train, y_train)
    return classifier


def predict_next_block(classifier: RandomForestClassifier, data_frame: pd.DataFrame) -> pd.DataFrame:
    """ Предсказание для следующего блока """
    last_block = data_frame.iloc[-1]
    # Преобразование последнего блока в DataFrame с теми же названиями столбцов
    last_block_df = pd.DataFrame([last_block], columns=data_frame.columns)
    classifier_prediction = classifier.predict(last_block_df)
    return classifier_prediction
