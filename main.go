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

func main() {
	if len(os.Args) < 2 {
		printUsage(os.Args[0])
	}
	currentBalance := os.Args[1]

	records, err := ing.ReadFromReader(os.Stdin)
	if err != nil {
		panic(err)
	}
	m := extract(records)

	balance, err := strconv.ParseFloat(currentBalance, 64)
	if err != nil {
		panic(err)
	}

	date := today()

	for i := 0; i < 365; i++ {
		date = date.Add(-24 * time.Hour)
		unix := date.Unix()

		entriesStr := ""
		dateBalance := 0.0

		entries, has := m[unix]
		for _, entry := range entries {
			dateBalance += entry.Amount
			entriesStr += fmt.Sprintf("\n   %.2fEUR\t\t%s", entry.Amount, entry.Description)
		}

		if has {
			fmt.Printf("%s\t\t%.2f\t%.2f", date.Format("2006-01-02"), balance, dateBalance)
		} else {
			fmt.Printf("%s\t\t%.2f", date.Format("2006-01-02"), balance)
		}
		fmt.Printf("%s\n\n", entriesStr)

		balance -= dateBalance
	}
}
