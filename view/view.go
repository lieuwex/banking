package view

import (
	"banking/types"
	"banking/utils"
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func createTable(timeDelta time.Duration, days []types.Day) *CustomTable {
	date := utils.Today()
	res := MakeCustomTable()

	for i := 0; i < 365; i++ {
		d := date.Add(-1 * timeDelta)
		balanceDifference, balance, entries := getBetween(days, d, date)
		date = d

		var lastCell string
		if balanceDifference != 0 {
			lastCell = formatPrice(balanceDifference, true)
		}
		res.AddRow(
			nil,
			date.Format("2006-01-02"),
			"",
			"",
			formatPrice(balance, false),
			lastCell,
		)

		for _, entry := range entries {
			entryCopy := entry

			res.AddRow(
				&entryCopy,
				"",
				formatPrice(entry.Amount, true),
				entry.Description,
			)
		}
	}

	return res
}

type ViewState struct {
	openID    string
	days      []types.Day
	timeDelta time.Duration

	table   *CustomTable
	infoBox *tview.TextView
	app     *tview.Application
}

func MakeViewState(days []types.Day) *ViewState {
	return &ViewState{
		openID: "",
		days:   days,
	}
}

func (s *ViewState) getKeyHandler() func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			s.app.Stop()

		default:
			if event.Key() == 13 {
				// TODO
			}
		}

		return event
	}
}

func (s *ViewState) getSelectedHandler() func(entry *types.Entry) {
	return func(entry *types.Entry) {
		infoBoxStr := fmt.Sprintf(
			"description: %s\n\naccount: %s\ncounter account: %s\n\n%s",
			entry.Description,
			entry.Account,
			entry.CounterAccount,
			entry.Information,
		)
		s.infoBox.SetText(infoBoxStr)
	}
}

func (state *ViewState) Run() error {
	// prepare
	state.timeDelta = getTimeByOption(ByWeek)
	state.table = createTable(state.timeDelta, state.days)
	state.table.AddSelectionHandler(state.getSelectedHandler())

	state.infoBox = tview.NewTextView()
	state.infoBox.SetBorder(true).SetTitle("info")

	flex := tview.NewFlex().
		AddItem(state.table.Table, 0, 2, true).
		AddItem(state.infoBox, 0, 1, false)

	// make real view
	state.app = tview.NewApplication()
	state.app.SetRoot(flex, true)
	state.app.SetInputCapture(state.getKeyHandler())

	// run
	return state.app.Run()
}
