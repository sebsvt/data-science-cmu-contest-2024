from sqlalchemy import Column, Integer, Float, String, Date, ForeignKey, TIMESTAMP, Text
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.orm import relationship
from sqlalchemy.ext.declarative import declarative_base
import uuid

Base = declarative_base()

class Forecasting(Base):
	__tablename__ = 'forecasting'

	id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
	message = Column(Text, nullable=True)
	created_at = Column(TIMESTAMP, nullable=False)
	partner_id = Column(UUID(as_uuid=True), nullable=False)

	results = relationship("ForecastResult", back_populates="forecasting")


class ForecastResult(Base):
	__tablename__ = 'forecast_results'

	id = Column(Integer, primary_key=True, autoincrement=True)
	forecasting_id = Column(UUID(as_uuid=True), ForeignKey('forecasting.id'), nullable=False)
	name = Column(String, nullable=False)
	from_last_date = Column(Date, nullable=False)

	kpi_id = Column(Integer, ForeignKey('kpis.id'), nullable=False)
	forecasting = relationship("Forecasting", back_populates="results")
	demand_forecastings = relationship("DemandForecasting", back_populates="forecast_result")
	kpi = relationship("KPI", back_populates="forecast_result")


class DemandForecasting(Base):
	__tablename__ = 'demand_forecastings'

	id = Column(Integer, primary_key=True, autoincrement=True)
	forecast_result_id = Column(Integer, ForeignKey('forecast_results.id'), nullable=False)
	date = Column(Date, nullable=False)
	demand = Column(Float, nullable=False)

	forecast_result = relationship("ForecastResult", back_populates="demand_forecastings")


class KPI(Base):
	__tablename__ = 'kpis'

	id = Column(Integer, primary_key=True, autoincrement=True)
	bias_rel = Column(Float, nullable=False)
	mape = Column(Float, nullable=False)
	mae = Column(Float, nullable=False)
	rmse = Column(Float, nullable=False)

	forecast_result = relationship("ForecastResult", back_populates="kpi")
