package view

import "github.com/gdamore/tcell"

func (s *ViewState) queryInput(event *tcell.EventKey, fn func()) {
	keycode := event.Key()
	query := s.model.query

	if keycode == 13 {
		fn()
	} else if keycode == 127 && len(query) > 0 {
		query = query[:len(query)-1]
	} else {
		query += string(event.Rune())
	}

	s.model.query = query
}

func (s *ViewState) getMainKeyHandler() func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		var fn func()
		if s.model.isSearching {
			fn = s.finishSearch
		} else if s.model.isCommanding {
			fn = s.finishCommand
		}

		if fn != nil {
			s.queryInput(event, s.finishCommand)
			s.redrawStuff()
			return nil
		}

		switch event.Rune() {
		case '/':
			s.startSearch()
		case ':':
			s.startCommand()

		case '{':
			panic("TODO {")
		case '}':
			panic("TODO }")

		default:
			if event.Key() == 13 {
				// TODO
			}
		}

		s.redrawStuff()
		return event
	}
}
