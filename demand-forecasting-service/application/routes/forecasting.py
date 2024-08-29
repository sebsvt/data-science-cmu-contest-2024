from application.schemas import ForecastingDemandRequest, ForecastingDemand, ForecastingDemandStatusEnum
from fastapi import APIRouter
from datetime import datetime

router = APIRouter()

@router.post("/")
async def demand_forecasting(request: ForecastingDemandRequest):
	return request
