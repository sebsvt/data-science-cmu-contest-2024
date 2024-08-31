from application.schemas import ForecastingDemandRequest, ForecastingDemand, ForecastingDemandStatusEnum
from eclair.usesage import demand_forecaste_next_day_from_the_last_date
from fastapi import APIRouter
from datetime import datetime
import pandas as pd

router = APIRouter()

@router.post("/")
async def demand_forecasting(request: ForecastingDemandRequest):
	try:
		df = pd.read_csv("./pure_data.csv")
		demand_forecast = demand_forecaste_next_day_from_the_last_date(
			data=df,
			datetime_col=request.timeseries_name,
			demand_type_col=request.demand_type_column,
			filter_by=request.items[0],
			demand=request.predicted_column_name
		)
		return demand_forecast
	except Exception as error:
		return {"error": error}
