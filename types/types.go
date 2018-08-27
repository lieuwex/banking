package types

import (
	"fmt"
	"strings"
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

	Tags []string
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

func (e Entry) Matches(str string) bool {
	base := strings.ToLower(e.Description)
	str = strings.ToLower(str)

	return strings.Contains(base, str)
}

type Day struct {
	Entries []*Entry
	Date    time.Time

	DateBalance      float64
	BalanceAfterDate float64
}
