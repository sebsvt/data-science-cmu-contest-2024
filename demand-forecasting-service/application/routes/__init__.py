from application.routes.forecasting import router as forecasting_api
from fastapi import APIRouter

router = APIRouter()

router.include_router(router=forecasting_api, prefix='/demand-forecasting', tags=['Demand Forecasting'])
