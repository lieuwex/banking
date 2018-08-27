package view

import (
	"banking/types"
	"time"
)

type model struct {
	// data
	entries []types.Entry
	balance float64

	filter string

	timeDelta time.Duration // REVIEW: config struct?

	isTagging bool

	isSearching  bool
	isCommanding bool
	query        string
}
