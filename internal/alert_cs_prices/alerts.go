package alert_cs_prices

import (
	"encoding/json"
	"fmt"
	"log"
	"math"

	"github.com/bismastr/discord-bot/internal/bot"
	"github.com/bismastr/discord-bot/internal/messaging"
)

type AlertPriceSertvice struct {
	consumer *messaging.Consumer
	bot      *bot.BotService
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

func NewAlertPriceServcie(consumer *messaging.Consumer, bot *bot.BotService) (*AlertPriceSertvice, error) {
	return &AlertPriceSertvice{
		consumer: consumer,
		bot:      bot,
	}, nil
}

func (a *AlertPriceSertvice) DailyReportSummary() error {
	msgs, close, err := a.consumer.Consume("notification_price_alerts")
	if err != nil {
		log.Printf("Error decoding message: %v", err)
		return err
	}
	defer close()

	for d := range msgs {
		var dailySummary NotificationPriceSummary
		err := json.Unmarshal(d.Body, &dailySummary)
		if err != nil {
			return err
		}

		report := fmt.Sprintf("📊 **DAILY SUMMARY** <@%d> 📊 FOR \n", dailySummary.DiscordId, dailySummary.ItemId)
		report += "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"
		report += fmt.Sprintf("🟢 **Open**:   $%.2f\n", dailySummary.OpeningPrice)
		report += fmt.Sprintf("🔴 **Close**:  $%.2f\n", dailySummary.ClosingPrice)
		report += fmt.Sprintf("🔺 **High**:    $%.2f\n", dailySummary.MaxPrice)
		report += fmt.Sprintf("🔻 **Low**:     $%.2f\n", dailySummary.MinPrice)
		report += fmt.Sprintf("📌 **Avg**:     $%.2f\n", dailySummary.AvgPrice)
		report += fmt.Sprintf("📈 **Change**: %s%.2f%%\n", getChangeEmoji(dailySummary.ChangePct), math.Abs(dailySummary.ChangePct))
		report += "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"

		a.bot.SendMessageToChannel("1276782792876888075", report)
	}

	return nil
}

// Helper function for change emoji
func getChangeEmoji(change float64) string {
	if change >= 0 {
		return "⬆️ "
	}
	return "⬇️ "
}
