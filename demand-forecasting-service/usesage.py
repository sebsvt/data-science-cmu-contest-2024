from eclair.demand_forecasting.models.statistics import simple_exponential_smoothing_model
from eclair.demand_forecasting.models.matchine_learning import prophet_forecasting_model
from eclair.demand_forecasting.kpi import forecast_kpi
from eclair.demand_forecasting.processing import extract_data_by_item_name_and_group_by, detection_and_delete_outlier_by_quatile, detection_and_delete_outlier_by_std
import pandas as pd

def main():
	df = pd.read_csv('./pure_data.csv')
	document_df = extract_data_by_item_name_and_group_by(data=df, datetime_column='ds', column='item_name', filter_value='เอกสาร', y='qty')
	document_df = detection_and_delete_outlier_by_std(document_df, 'y')
	result = simple_exponential_smoothing_model(demand=document_df['y'])
	result_ml = prophet_forecasting_model(data=document_df, periods=1)
	# print(result_ml)
	forecast_kpi(result_ml)

main()
