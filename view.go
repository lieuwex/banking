package main

import (
	"banking/types"
	"fmt"
	"time"

	"github.com/rivo/tview"
)

type GroupBy int

const (
	ByDay GroupBy = iota
	ByWeek
	ByMonth
	ByYear
)

type Row struct {
	Selectable bool
	Words      []string
}
type CustomTable struct {
	Table *tview.Table

	rows []Row

	selected int
}

func MakeCustomTable() *CustomTable {
	res := &CustomTable{
		Table:    tview.NewTable(),
		selected: 0,
	}

	res.Table.SetSelectable(true, false)

	res.Table.SetSelectionChangedFunc(func(row, column int) {
		if res.rows[row].Selectable {
			res.selected = row
			return
		}

		var delta int
		if row < res.selected { // going up
			delta = -1
		} else {
			delta = 1
		}

		for {
			res.selected += delta
			if res.selected < 0 ||
				res.selected >= len(res.rows) ||
				res.rows[res.selected].Selectable {
				break
			}
		}
		if res.selected < 0 {
			res.selected = 0
		} else if res.selected >= len(res.rows) {
			res.selected = len(res.rows) - 1
		}

		res.Table.Select(res.selected, 0)
	})

	return res
}
func (t *CustomTable) AddRow(selectable bool, items ...string) {
	row := Row{Selectable: selectable}
	index := len(t.rows)

	for i, item := range items {
		if item == "" {
			continue
		}

		t.Table.SetCellSimple(index, i, item)
		row.Words = append(row.Words, item)
	}

	t.rows = append(t.rows, row)
}

func formatPrice(amount float64, useColor bool) string {
	color := "green"
	if amount < 0 {
		color = "red"
	}

	if useColor {
		return fmt.Sprintf("[:%s]%.2f EUR", color, amount)
	}
	return fmt.Sprintf("%.2f EUR", amount)
}

func getTimeByOption(option GroupBy) time.Duration {
	day := 24 * time.Hour

	switch option {
	case ByDay:
		return day
	case ByWeek:
		return 7 * day
	case ByMonth:
		return 30 * day
	case ByYear:
		return 365 * day
	}

	return -1
}

func getBetween(days []Day, start, end time.Time) (balanceDifference, balance float64, entries []types.Entry) {
	balanceDifference = 0

	for _, day := range days {
		if day.Date.Unix() < start.Unix() ||
			day.Date.Unix() >= end.Unix() {
			continue
		}

		for _, entry := range day.Entries {
			entries = append(entries, entry)
		}

		balance = day.BalanceAfterDate
		balanceDifference += day.DateBalance
	}

	return balanceDifference, balance, entries
}

func RunView(days []Day) error {
	date := Today()
	timeDelta := getTimeByOption(ByDay)

	table := MakeCustomTable()

	for i := 0; i < 365; i++ {
		d := date.Add(-1 * timeDelta)
		balanceDifference, balance, entries := getBetween(days, d, date)
		date = d

		var lastCell string
		if balanceDifference != 0 {
			lastCell = formatPrice(balanceDifference, true)
		}
		table.AddRow(
			false,
			date.Format("2006-01-02"),
			"",
			"",
			formatPrice(balance, false),
			lastCell,
		)

		for _, entry := range entries {
			table.AddRow(
				true,
				"",
				formatPrice(entry.Amount, true),
				entry.Description,
			)
		}
	}

	return tview.NewApplication().SetRoot(table.Table, true).Run()
}
