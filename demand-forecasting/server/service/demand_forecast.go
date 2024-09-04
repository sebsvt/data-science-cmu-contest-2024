package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sebsvt/message-broker/aggregate"
)

type DemandForecastCreated struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	BuilderID   uuid.UUID `json:"builder_id"`
	PartnerID   uuid.UUID `json:"partner_id"`
}

type DemandForecast struct {
	Title       string                                `json:"title"`
	Description string                                `json:"description"`
	Status      aggregate.ForecastingDemandStatusEnum `json:"status"`
	PartnerID   uuid.UUID                             `json:"partner_id"`
	CreatedAt   time.Time                             `json:"created_at"`
}

type DemandForecastService interface {
	CreateNewDemandForecast(ctx context.Context, demand_forecast_created DemandForecastCreated) (uuid.UUID, error)
	GetDemandForecastByID(ctx context.Context, forecast_id uuid.UUID) (*aggregate.DemandForecast, error)
	UpdateDemandForecast() error
}
