from application.schemas import ForecastingDemandRequest, ForecastingDemand, ForecastingDemandStatusEnum
from eclair.usesage import demand_forecaste_next_day_from_the_last_date
from fastapi import APIRouter
from datetime import datetime
import pandas as pd

router = APIRouter()

@router.post("/")
async def demand_forecasting(request: ForecastingDemandRequest):
	try:
		print(request)
		demand_forecast = []
		df = pd.read_csv("./pure_data.csv")
		for i in request.items:
			result = demand_forecaste_next_day_from_the_last_date(
				data=df,
				datetime_col=request.timeseries_name,
				demand_type_col=request.demand_type_column,
				filter_by=i,
				demand=request.predicted_column_name
			)
			demand_forecast.append(result)

		return demand_forecast
	except Exception as error:
		print(error)
		return {"error": error}
