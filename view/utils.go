package view

import (
	"banking/types"
	"fmt"
	"strings"
	"time"
)

func getBetween(days []types.Day, start, end time.Time) (balanceDifference, balance float64, entries []*types.Entry) {
	balanceDifference = 0

	for _, day := range days {
		if day.Date.Unix() < start.Unix() ||
			day.Date.Unix() >= end.Unix() {
			continue
		}

		for _, entry := range day.Entries {
			entries = append(entries, entry)
		}

		balance = day.BalanceAfterDate
		balanceDifference += day.DateBalance
	}

	return balanceDifference, balance, entries
}

func formatPrice(amount float64, useColor bool) string {
	if useColor {
		color := "green"
		if amount < 0 {
			color = "red"
		}

		return fmt.Sprintf("[:%s]%.2f EUR", color, amount)
	}

	return fmt.Sprintf("%.2f EUR", amount)
}

func getTimeByOption(option GroupBy) time.Duration {
	day := 24 * time.Hour

	switch option {
	case ByDay:
		return day
	case ByWeek:
		return 7 * day
	case ByMonth:
		return 30 * day
	case ByYear:
		return 365 * day
	}

	return -1
}

var templateFuncs = map[string]interface{}{
	"join": func(args ...interface{}) string {
		arr := args[0].([]string)
		delim := args[1].(string)
		return strings.Join(arr, delim)
	},
}
