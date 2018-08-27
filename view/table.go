package view

import (
	"banking/types"

	"github.com/rivo/tview"
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

	selectedHandler func(entry *types.Entry)
}

func MakeCustomTable() *CustomTable {
	res := &CustomTable{
		Table: tview.NewTable(),

		selected:         0,
		rowNumberToEntry: make(map[int]*types.Entry),
	}

	res.Table.SetSelectable(true, false)

	onSelectionDone := func() {
		if res.selectedHandler != nil {
			entry := res.GetSelected()
			if entry != nil {
				res.selectedHandler(entry)
			}
		}
	}

	res.Table.SetSelectionChangedFunc(func(row, column int) {
		if res.rows[row].Selectable {
			res.selected = row
			onSelectionDone()
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
		onSelectionDone()
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

func (t *CustomTable) AddSelectionHandler(fn func(entry *types.Entry)) {
	t.selectedHandler = fn
}
