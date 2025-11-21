package alert_cs_prices

import (
	"github.com/bismastr/discord-bot/internal/messaging"
	"github.com/bismastr/discord-bot/internal/repository"
	"github.com/rabbitmq/amqp091-go"
)

type AlertPriceSertvice struct {
	consumer *messaging.Consumer
}

type NotificationPriceSummary struct {
	ItemId          int32   `json:"item_id"`
	ChangePct       float64 `json:"change_pct"`
	Name            string  `json:"name"`
	AlertType       string  `json:"alert_type"`
	LatestSellPrice int32   `json:"latest_price"`
	OldSellPrice    int32   `json:"old_price"`
}

func NewAlertPriceServcie(consumer *messaging.Consumer, repositoryCsPrices *repository.Queries) (*AlertPriceSertvice, error) {
	return &AlertPriceSertvice{
		consumer: consumer,
	}, nil
}

func (a *AlertPriceSertvice) DailyReportSummary() (<-chan amqp091.Delivery, func(), error) {
	msgs, close, err := a.consumer.Consume("discord_notifications")

	return msgs, close, err
}
