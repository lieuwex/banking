package main

import (
	"banking/ing"
	"banking/types"
	"fmt"
	"os"
	"strconv"
	"time"
)

func extract(records []types.Entry) map[int64][]types.Entry {
	m := make(map[int64][]types.Entry)
	for _, record := range records {
		unix := record.Date.Unix()
		m[unix] = append(m[unix], record)
	}
	return m
}

func printUsage(cmd string) {
	fmt.Printf("usage: %s <current balance>\n", cmd)
	os.Exit(1)
}

type Day struct {
	Entries []types.Entry
	Date    time.Time

	DateBalance      float64
	BalanceAfterDate float64
}

func entriesToDays(balance float64, entries []types.Entry) []Day {
	m := extract(entries)

	days := make([]Day, 365)

	date := Today()
	for i := 0; i < 365; i++ {
		date = date.Add(-24 * time.Hour)
		unix := date.Unix()
		entries, has := m[unix]
		if !has {
			entries = []types.Entry{}
		}

		dateBalance := 0.0
		for _, entry := range entries {
			dateBalance += entry.Amount
		}

		day := Day{
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

func main() {
	if len(os.Args) < 2 {
		printUsage(os.Args[0])
	}
	currentBalance := os.Args[1]

	records, err := ing.ReadFromReader(os.Stdin)
	if err != nil {
		panic(err)
	}

	balance, err := strconv.ParseFloat(currentBalance, 64)
	if err != nil {
		panic(err)
	}

	days := entriesToDays(balance, records)

	if err := RunView(days); err != nil {
		panic(err)
	}
}
