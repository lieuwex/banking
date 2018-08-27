package calc

import (
	"banking/types"
	"banking/utils"
	"time"
)

func extract(records []*types.Entry) map[int64][]*types.Entry {
	m := make(map[int64][]*types.Entry)
	for _, record := range records {
		unix := record.Date.Unix()
		m[unix] = append(m[unix], record)
	}
	return m
}

func EntriesToDays(balance float64, entries []*types.Entry) []types.Day {
	m := extract(entries)

	days := make([]types.Day, 365)

	date := utils.Today()
	for i := 0; i < 365; i++ {
		date = date.Add(-24 * time.Hour)
		unix := date.Unix()
		entries, has := m[unix]
		if !has {
			entries = []*types.Entry{}
		}

		dateBalance := 0.0
		for _, entry := range entries {
			dateBalance += entry.Amount
		}

		day := types.Day{
			Entries: entries,
			Date:    date,

			DateBalance:      dateBalance,
			BalanceAfterDate: balance,
		}
		days[len(days)-i-1] = day

		balance -= dateBalance
	}

	return days
}
