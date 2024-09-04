package aggregate

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ForecastingDemandStatusEnum string

const (
	Pending   ForecastingDemandStatusEnum = "pending"
	Failed    ForecastingDemandStatusEnum = "failed"
	Completed ForecastingDemandStatusEnum = "completed"
)

type DemandForecast struct {
	ForecastID      uuid.UUID                   `bson:"forecast_id" json:"forecast_id"`
	Title           string                      `bson:"title" json:"title"`
	Description     string                      `bson:"description" json:"description"`
	PredictedDemand []PredictedDemand           `bson:"demand_forecast_item" json:"demand_forecast_item"`
	Status          ForecastingDemandStatusEnum `bson:"status" json:"status"`
	CreatedAt       time.Time                   `bson:"created_at" json:"created_at"`
	PartnerID       uuid.UUID                   `bson:"partner_id" json:"partner_id"`
	BuilderID       uuid.UUID                   `bson:"builder_id" json:"builder_id"`
}

type PredictedDemand struct {
	Name         string         `bson:"name" json:"name"` // (aka: Y, label)
	FromLastDate string         `bson:"from_last_date" json:"from_last_date"`
	Predictions  []FutureDemand `bson:"predictions" json:"predictions"`
	Kpi          Kpi            `bson:"kip" json:"kpi"`
}

type FutureDemand struct {
	FutureDate string  `bson:"future_date" json:"future_date"`
	Demand     float64 `bson:"demand" json:"demand"`
}

type Kpi struct {
	Name  string  `bson:"name" json:"name"`
	Value float64 `bson:"value" json:"value"`
}

type DemandForecastRepository interface {
	NextIdentity() uuid.UUID
	FromForecastID(ctx context.Context, forecast_id uuid.UUID) (*DemandForecast, error)
	Save(ctx context.Context, entity DemandForecast) error
}
