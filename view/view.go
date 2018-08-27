package view

import (
	"banking/calc"
	"banking/types"
	"fmt"

	"github.com/rivo/tview"
)

type ViewState struct {
	model model

	table      *CustomTable
	infoBox    *tview.TextView
	tagModal   tview.Primitive
	commandBar *tview.TextView
	app        *tview.Application
}

func MakeViewState(balance float64, entries []types.Entry) *ViewState {
	return &ViewState{
		model: model{
			entries: entries,
			balance: balance,
		},
	}
}

// REVIEW
func (s *ViewState) redrawStuff() {
	setCommandBarText := func(prefix, str string) {
		str = fmt.Sprintf("\n%s%s", prefix, str)
		s.commandBar.SetText(str)
	}

	if s.model.isSearching {
		setCommandBarText("/", s.model.query)
	} else if s.model.isCommanding {
		setCommandBarText(":", s.model.query)
	} else {
		setCommandBarText("", "")
	}

	s.app.Draw()
}

func (state *ViewState) Run() error {
	// prepare
	days := calc.EntriesToDays(state.model.balance, state.model.entries)
	state.model.timeDelta = getTimeByOption(ByDay)

	// create table
	state.table = createTable(state.model.timeDelta, days)
	state.table.AddSelectionHandler(func(entry *types.Entry) {
		state.updateInfoBox(entry)
	})

	// create infoBox
	state.infoBox = tview.NewTextView()
	state.infoBox.SetBorder(true).SetTitle("info")

	// combine into topRow
	topRow := tview.NewFlex().
		AddItem(state.table.Table, 0, 2, true).
		AddItem(state.infoBox, 0, 1, false)

	// create commandBar
	state.commandBar = tview.NewTextView()

	// combine into background
	background := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(topRow, 0, 1, true).
		AddItem(state.commandBar, 2, 1, false)

	// create application
	state.app = tview.NewApplication()
	state.app.SetRoot(background, true)
	state.app.SetInputCapture(state.getMainKeyHandler())

	// run
	return state.app.Run()
}
