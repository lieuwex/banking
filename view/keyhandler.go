package view

import "github.com/gdamore/tcell"

func (s *ViewState) getMainKeyHandler() func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		if s.model.isSearching {
			if event.Key() == 13 {
				s.finishSearch()
			} else {
				s.model.query += string(event.Rune())
			}

			s.redrawStuff()
			return nil
		} else if s.model.isCommanding {
			if event.Key() == 13 {
				s.finishCommand()
			} else {
				s.model.query += string(event.Rune())
			}

			s.redrawStuff()
			return nil
		}

		switch event.Rune() {
		case 'q':
			s.app.Stop()

		case '/':
			s.startSearch()
		case ':':
			s.startCommand()

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

		s.redrawStuff()
		return event
	}
}
