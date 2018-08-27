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
	days      []types.Day
	timeDelta time.Duration

	isTagging bool

	table    *CustomTable
	infoBox  *tview.TextView
	tagModal tview.Primitive
	pages    *tview.Pages
	app      *tview.Application
}

func MakeViewState(days []types.Day) *ViewState {
	return &ViewState{
		days: days,
	}
}

func (s *ViewState) getKeyHandler() func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			s.app.Stop()

		case '{':
			panic("TODO {")
		case '}':
			panic("TODO }")

		case 't':
			s.isTagging = true
			s.pages.ShowPage("tag-modal")

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
		var firstTag string
		if len(entry.Tags) >= 1 {
			firstTag = entry.Tags[0]
		}

		infoBoxStr := fmt.Sprintf(
			"description: %s\n\naccount: %s\ncounter account: %s\n\n%s\n\nfirst tag: %s",
			entry.Description,
			entry.Account,
			entry.CounterAccount,
			entry.Information,
			firstTag,
		)
		s.infoBox.SetText(infoBoxStr)
	}
}

func modal(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewGrid().
		SetColumns(0, width, 0).
		SetRows(0, height, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)
}

func (s *ViewState) createTagModal() tview.Primitive {
	field := tview.NewInputField().SetLabel("add tag")
	field.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}

		tag := field.GetText()
		entry := s.table.GetSelected()
		entry.Tags = append(entry.Tags, tag)
		s.isTagging = false
		s.pages.HidePage("tag-modal")
	})

	return modal(field, 40, 3)
}

func (state *ViewState) Run() error {
	// prepare
	state.timeDelta = getTimeByOption(ByDay)
	state.table = createTable(state.timeDelta, state.days)
	state.table.AddSelectionHandler(state.getSelectedHandler())

	state.infoBox = tview.NewTextView()
	state.infoBox.SetBorder(true).SetTitle("info")

	flex := tview.NewFlex().
		AddItem(state.table.Table, 0, 2, true).
		AddItem(state.infoBox, 0, 1, false)

	state.tagModal = state.createTagModal()

	state.pages = tview.NewPages().
		AddPage("background", flex, true, true).
		AddPage("tag-modal", state.tagModal, true, true).
		HidePage("tag-modal")

	// make real view
	state.app = tview.NewApplication()
	state.app.SetRoot(state.pages, true)
	state.app.SetInputCapture(state.getKeyHandler())

	// run
	return state.app.Run()
}
