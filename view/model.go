package view

import (
	"banking/types"
	"time"
)

type model struct {
	// data
	entries []types.Entry
	balance float64

	timeDelta time.Duration // REVIEW: config struct?
	isTagging bool
}
