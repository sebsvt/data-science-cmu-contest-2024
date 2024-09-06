interface KPI {
  bias_rel: number;
  mape: number;
  mae: number;
  rmse: number;
}

interface Prediction {
  future_date: string; // ISO date string format
  demand: number;
}

interface DemandForecastItem {
  name: string;
  from_last_date: string; // ISO date string format
  predictions: Prediction[];
  kpi: KPI;
}

interface DemandForecast {
  forecast_id: string;
  title: string;
  description: string;
  demand_forecast_item: DemandForecastItem[];
  status: string;
  created_at: string; // ISO date string format
  partner_id: string;
  builder_id: string;
  number_of_item: number;
}
