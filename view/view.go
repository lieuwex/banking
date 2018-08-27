package view

import (
	"banking/types"
	"banking/utils"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

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
