package types

import (
	"fmt"
	"time"
)

type Entry struct {
	Date           time.Time
	Description    string
	Account        string
	CounterAccount string
	Amount         float64
	MutationType   string // REVIEW
	Information    string
}

func (e Entry) String() string {
	return fmt.Sprintf(
		"%s\t%vEUR\t%s<->%s (%s %s)",
		e.Date.Format("2006-01-02"),
		e.Amount,
		e.Account,
		e.CounterAccount,
		e.Description,
		e.Information,
	)
}
