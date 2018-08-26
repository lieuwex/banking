package main

import (
	"banking/ing"
	"banking/types"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func extract(records []types.Entry) (keys []int, values map[int][]float64) {
	// fill map
	m := make(map[int][]float64)
	for _, record := range records {
		unix := int(record.Date.Unix())
		m[unix] = append(m[unix], record.Amount)
	}

	// get sorted keys
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return keys, m
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
	keys, values := extract(records)

	balance, err := strconv.ParseFloat(currentBalance, 64)
	if err != nil {
		panic(err)
	}

	for i := len(keys) - 1; i >= 0; i-- {
		date := keys[i]
		dateBalance := 0.0
		for _, amount := range values[date] {
			dateBalance += amount
		}

		balance -= dateBalance

		fmt.Printf(
			"%s\t\t%.2f\t\t%.2f\n",
			time.Unix(int64(date), 0).Format("2006-01-02"),
			dateBalance,
			balance,
		)
	}
}
