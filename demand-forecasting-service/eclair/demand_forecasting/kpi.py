import pandas as pd
import numpy as np

def forecast_kpi(df: pd.DataFrame):
	demand_avg = df.loc[df['Error'].notnull(), 'Demand'].mean()
	bias_abs = df['Error'].mean()
	bias_rel = bias_abs/demand_avg
	print('Bias: {:0.2f}, {:.2%}'.format(bias_abs,bias_rel))
	MAPE = (df['Error'].abs()/df['Demand']).mean()
	print('MAPE: {:.2%}'.format(MAPE))
	MAE_abs = df['Error'].abs().mean()
	MAE_rel = MAE_abs / demand_avg
	print('MAE: {:0.2f}, {:.2%}'.format(MAE_abs, MAE_rel))
	RMSE_abs = np.sqrt((df["Error"] ** 2).mean())
	RMSE_rel = RMSE_abs / demand_avg
	print('RMSE: {:0.2f}, {:.2%}'.format(RMSE_abs, RMSE_rel))
