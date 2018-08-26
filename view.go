package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
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
	} else {
		return fmt.Sprintf("%.2f EUR", amount)
	}
}

func RunView(days []Day) error {
	table := MakeCustomTable()

	date := Today()
	for i := 0; i < 365; i++ {
		date = date.Add(-24 * time.Hour)

		index := -1
		for i, day := range days {
			if Date(day.Date) == Date(date) {
				index = i
				break
			}
		}
		if index == -1 {
			continue
		}
		day := days[index]

		var lastCell string
		if day.DateBalance != 0 {
			lastCell = formatPrice(day.DateBalance, true)
		}
		table.AddRow(
			false,
			date.Format("2006-01-02"),
			"",
			"",
			formatPrice(day.BalanceAfterDate, false),
			lastCell,
		)

		for _, entry := range day.Entries {
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
