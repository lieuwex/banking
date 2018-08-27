package view

import (
	"banking/calc"
	"banking/types"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type ViewState struct {
	model model

	table      *CustomTable
	infoBox    *tview.TextView
	tagModal   tview.Primitive
	commandBar *tview.TextView
	pages      *tview.Pages
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

func (s *ViewState) getMainKeyHandler() func(event *tcell.EventKey) *tcell.EventKey {
	query := ""

	return func(event *tcell.EventKey) *tcell.EventKey {
		if s.model.isSearching {
			if event.Key() == 13 {
				s.finishSearch(query)
				query = ""
			} else {
				query += string(event.Rune())
				s.setCommandBarText("/", query)
			}
			s.app.Draw()
			return nil
		} else if s.model.isCommanding {
			if event.Key() == 13 {
				s.finishSearch(query)
				query = ""
			} else {
				query += string(event.Rune())
				s.setCommandBarText(":", query)
			}
			s.app.Draw()
			return nil
		}

		switch event.Rune() {
		case 'q':
			s.app.Stop()

		case '/':
			s.startSearch()

		case '{':
			panic("TODO {")
		case '}':
			panic("TODO }")

		case 't':
			s.startTagging()

		default:
			if event.Key() == 13 {
				// TODO
			}
		}

		return event
	}
}

func (s *ViewState) createTagModal() tview.Primitive {
	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewGrid().
			SetColumns(0, width, 0).
			SetRows(0, height, 0).
			AddItem(p, 1, 1, 1, 1, 0, 0, true)
	}

	field := tview.NewInputField().SetLabel("add tag")
	field.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}

		s.finishTagging(field.GetText())
	})

	return modal(field, 40, 3)
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

	// create tagModal
	state.tagModal = state.createTagModal()

	// create pages (root view)
	state.pages = tview.NewPages().
		AddPage("background", background, true, true).
		AddPage("tag-modal", state.tagModal, true, true).
		HidePage("tag-modal")

	// create application
	state.app = tview.NewApplication()
	state.app.SetRoot(state.pages, true)
	state.app.SetInputCapture(state.getMainKeyHandler())

	// run
	return state.app.Run()
}
