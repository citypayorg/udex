package crons

import (
	"fmt"
	"log"

	"github.com/citypayorg/udex/tree/udex/backend/types"
	"github.com/citypayorg/udex/tree/udex/backend/utils"
	"github.com/citypayorg/udex/tree/udex/backend/ws"

	"github.com/citypayorg/udex/tree/udex/backend/app"
	"github.com/robfig/cron"
)

// tickStreamingCron takes instance of cron.Cron and adds tickStreaming
// crons according to the durations mentioned in config/app.yaml file
func (s *CronService) tickStreamingCron(c *cron.Cron) {
	for unit, durations := range app.Config.TickDuration {
		for _, duration := range durations {
			schedule := getCronScheduleString(unit, duration)
			c.AddFunc(schedule, s.tickStream(unit, duration))
		}
	}
}

// tickStream function fetches latest tick based on unit and duration for each pair
// and broadcasts the tick to the client subscribed to pair's respective channel
func (s *CronService) tickStream(unit string, duration int64) func() {
	return func() {
		p := make([]types.PairAssets, 0)
		ticks, err := s.ohlcvService.GetOHLCV(p, duration, unit)
		if err != nil {
			log.Printf("%s", err)
			return
		}

		for _, tick := range ticks {
			baseAsset := tick.Pair.BaseToken
			quoteAsset := tick.Pair.QuoteToken
			id := utils.GetTickChannelID(baseAsset, quoteAsset, unit, duration)
			ws.GetOHLCVSocket().BroadcastOHLCV(id, tick)
		}
	}
}

// getCronScheduleString converts unit and duration to schedule string used for
// cron addFunc to schedule crons
func getCronScheduleString(unit string, duration int64) string {
	switch unit {

	case "sec":
		return fmt.Sprintf("*/%d * * * * *", duration)

	case "min":
		return fmt.Sprintf("0 */%d * * * *", duration)

	case "hour":
		return fmt.Sprintf("0 0 %d * * *", duration)

	case "day":
		return fmt.Sprintf("@daily")

	case "week":
		return fmt.Sprintf("0 0 0 * * */%d", duration)

	case "month":
		return fmt.Sprintf("0 0 0 */%d * *", duration)

	case "year":
		return fmt.Sprintf("@yearly")

	default:
		panic(fmt.Errorf("Invalid unit please try again"))
	}
}
