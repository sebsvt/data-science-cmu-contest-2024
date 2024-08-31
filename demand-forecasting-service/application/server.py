from application.routes import router as api_router
from fastapi import FastAPI

# How many product demands for each provinces

class ForecastingDemand:
	forecast_id: int
	title: str
	description: str
	partner_id: int

def get_application() -> FastAPI:
	app = FastAPI(title="Demand Forecasting Model", version="0.0.1")
	app.include_router(router=api_router, prefix='/api')
	return app

app = get_application()
