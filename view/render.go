package view

import (
	"banking/types"
	"banking/utils"
	"fmt"
	"time"
)

func createTable(timeDelta time.Duration, days []types.Day) *CustomTable {
	date := utils.Today()
	res := MakeCustomTable()

	for i := 0; i < 365; i++ {
		d := date.Add(-1 * timeDelta)
		balanceDifference, balance, entries := getBetween(days, d, date)
		date = d

		var lastCell string
		if balanceDifference != 0 {
			lastCell = formatPrice(balanceDifference, true)
		}
		res.AddRow(
			nil,
			date.Format("2006-01-02"),
			"",
			"",
			formatPrice(balance, false),
			lastCell,
		)

		for _, entry := range entries {
			entryCopy := entry

			res.AddRow(
				&entryCopy,
				"",
				formatPrice(entry.Amount, true),
				entry.Description,
			)
		}
	}

	return res
}

func infoBoxString(entry *types.Entry) string {
	var firstTag string
	if len(entry.Tags) >= 1 {
		firstTag = entry.Tags[0]
	}

	return fmt.Sprintf(
		"description: %s\n\naccount: %s\ncounter account: %s\n\n%s\n\nfirst tag: %s",
		entry.Description,
		entry.Account,
		entry.CounterAccount,
		entry.Information,
		firstTag,
	)
}
