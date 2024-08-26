from fastapi import FastAPI

def get_application() -> FastAPI:
	app = FastAPI(title="Demand Forecasting Model", version="0.0.1")
	return app

app =get_application()
