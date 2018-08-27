package view

import (
	"banking/types"
	"banking/utils"
	"fmt"
	"time"

	"github.com/gdamore/tcell"
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

	selected         int
	rowNumberToEntry map[int]*types.Entry
}

func MakeCustomTable() *CustomTable {
	res := &CustomTable{
		Table: tview.NewTable(),

		selected:         0,
		rowNumberToEntry: make(map[int]*types.Entry),
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
func (t *CustomTable) AddRow(entry *types.Entry, items ...string) {
	row := Row{
		Selectable: entry != nil,
	}
	index := len(t.rows)

	for i, item := range items {
		if item == "" {
			continue
		}

		t.Table.SetCellSimple(index, i, item)
		row.Words = append(row.Words, item)
	}

	t.rowNumberToEntry[index] = entry
	t.rows = append(t.rows, row)
}
func (t *CustomTable) GetSelected() *types.Entry {
	return t.rowNumberToEntry[t.selected]
}

func formatPrice(amount float64, useColor bool) string {
	if useColor {
		color := "green"
		if amount < 0 {
			color = "red"
		}

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

func getBetween(days []types.Day, start, end time.Time) (balanceDifference, balance float64, entries []types.Entry) {
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

type ViewState struct {
	openID string
	days   []types.Day
}

func MakeViewState(days []types.Day) *ViewState {
	return &ViewState{
		openID: "",
		days:   days,
	}
}

func (state *ViewState) Run() error {
	date := utils.Today()
	timeDelta := getTimeByOption(ByWeek)

	table := MakeCustomTable()

	for i := 0; i < 365; i++ {
		d := date.Add(-1 * timeDelta)
		balanceDifference, balance, entries := getBetween(state.days, d, date)
		date = d

		var lastCell string
		if balanceDifference != 0 {
			lastCell = formatPrice(balanceDifference, true)
		}
		table.AddRow(
			nil,
			date.Format("2006-01-02"),
			"",
			"",
			formatPrice(balance, false),
			lastCell,
		)

		for _, entry := range entries {
			entryCopy := entry

			table.AddRow(
				&entryCopy,
				"",
				formatPrice(entry.Amount, true),
				entry.Description,
			)
		}
	}

	app := tview.NewApplication()
	app.SetRoot(table.Table, true)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			app.Stop()

		default:
			if event.Key() == 13 {
				entry := table.GetSelected()
				println(entry.Description)
			}
		}

		return event
	})
	return app.Run()
}
