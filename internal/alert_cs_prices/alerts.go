package alert_cs_prices

import (
	"context"

	"github.com/bismastr/discord-bot/internal/messaging"
	"github.com/bismastr/discord-bot/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rabbitmq/amqp091-go"
)

type AlertPriceSertvice struct {
	consumer           *messaging.Consumer
	repositoryCsPrices *repository.Queries
}

type NotificationPriceSummary struct {
	ItemId       int     `json:"item_id"`
	AvgPrice     float64 `json:"avg_price"`
	MaxPrice     float64 `json:"max_price"`
	MinPrice     float64 `json:"min_price"`
	OpeningPrice float64 `json:"opening_price"`
	ClosingPrice float64 `json:"closing_price"`
	ChangePct    float64 `json:"change_pct"`
	DiscordId    int64   `json:"discord_id"`
}

func NewAlertPriceServcie(consumer *messaging.Consumer, repositoryCsPrices *repository.Queries) (*AlertPriceSertvice, error) {
	return &AlertPriceSertvice{
		consumer:           consumer,
		repositoryCsPrices: repositoryCsPrices,
	}, nil
}

func (a *AlertPriceSertvice) GetItemsContainsName(ctx context.Context, query string) (*[]repository.GetItemsContainsNameRow, error) {
	res, err := a.repositoryCsPrices.GetItemsContainsName(ctx, pgtype.Text{
		String: query,
		Valid:  true,
	})
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (a *AlertPriceSertvice) AddDailySchedule(ctx context.Context, param repository.InsertAlertDailyScheduleParams) error {
	err := a.repositoryCsPrices.InsertAlertDailySchedule(ctx, param)
	if err != nil {
		return err
	}

	return nil
}

func (a *AlertPriceSertvice) DailyReportSummary() (<-chan amqp091.Delivery, func(), error) {
	msgs, close, err := a.consumer.Consume("notification_price_alerts")

	return msgs, close, err
}
