package main

import (
	"banking/ing"
	"banking/view"
	"fmt"
	"os"
	"strconv"
)

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

	balance, err := strconv.ParseFloat(currentBalance, 64)
	if err != nil {
		panic(err)
	}

	state := view.MakeViewState(balance, records)
	if err := state.Run(); err != nil {
		panic(err)
	}
}
