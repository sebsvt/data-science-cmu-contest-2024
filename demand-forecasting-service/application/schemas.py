from enum import Enum
from pydantic import BaseModel
from datetime import datetime

from pydantic import BaseModel
from datetime import datetime
from enum import Enum

class DateRange(BaseModel):
	start_date: str
	end_date: str

class ForecastingDemandStatusEnum(str, Enum):
	pending = "pending"
	failed = "failed"
	completed = "completed"

class ForecastingDemandRequest(BaseModel):
	title: str
	description: str
	timeseries_name: str
	demand_type_column: str
	items: list[str]
	predicted_column_name: str #(aka: Y, label)

class ForecastingDemand(BaseModel):
	title: str
	description: str
	predicted_column_name: str #(aka: Y, label)
	features: list[str]
	model: str
	forecasting_result: dict
	accuracy: float
	status: ForecastingDemandStatusEnum = ForecastingDemandStatusEnum.pending
	created_at: datetime
	partner_id: str
	product: str

	class Config:
		from_attributes=True
