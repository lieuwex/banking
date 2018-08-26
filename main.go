package main

import (
	"banking/ing"
	"fmt"
	"os"
)

func main() {
	records, err := ing.ReadFromReader(os.Stdin)
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
}
