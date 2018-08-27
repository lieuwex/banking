package ing

import (
	"banking/types"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadFromReader(r io.Reader) ([]*types.Entry, error) {
	// REVIEW: close?
	reader := csv.NewReader(os.Stdin)
	records, err := reader.ReadAll()
	if err != nil {
		return []*types.Entry{}, err
	}

	var res []*types.Entry
	for _, fields := range records[1:] {
		date, err := time.Parse("20060102", fields[0])
		if err != nil {
			return []*types.Entry{}, err
		}

		var sign float64
		switch fields[5] {
		case "Bij":
			sign = 1
		case "Af":
			sign = -1

		default:
			return []*types.Entry{}, fmt.Errorf("unknown type '%s'", fields[5])
		}

		raw := strings.Replace(fields[6], ",", ".", 1)
		amount, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return []*types.Entry{}, err
		}

		entry := &types.Entry{
			Date:           date,
			Description:    fields[1],
			Account:        fields[2],
			CounterAccount: fields[3],
			Amount:         sign * amount,
			MutationType:   fields[7],
			Information:    fields[8],
		}
		res = append(res, entry)
	}
	return res, nil
}
