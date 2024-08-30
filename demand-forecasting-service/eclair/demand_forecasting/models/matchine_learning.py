from prophet import Prophet
import pandas as pd

def prophet_forecasting_model(data: pd.DataFrame, periods: int = 4, freq: str = 'D') -> pd.DataFrame:
	model = Prophet()

	model.fit(data)

	future = model.make_future_dataframe(periods=periods, freq=freq)

	forecast = model.predict(future)

	return pd.DataFrame.from_dict({'Date': forecast['ds'], 'Demand': data['y'], 'Forecast': forecast['yhat'], 'ForecastLower': forecast['yhat_lower'], 'ForecastUpper': forecast['yhat_upper'], 'Error': data['y'] - forecast['yhat']})
