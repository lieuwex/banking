package view

import (
	"banking/types"
)

func (s *ViewState) updateInfoBox(entry *types.Entry) {
	str := infoBoxString(entry)
	s.infoBox.SetText(str)
}

func (s *ViewState) startSearch() {
	s.model.isSearching = true
}

func (s *ViewState) finishSearch() {
	s.model.isSearching = false
	s.model.filter = s.model.query
	s.model.query = ""
}

func (s *ViewState) startCommand() {
	s.model.isCommanding = true
}

func (s *ViewState) finishCommand() {
	s.model.isCommanding = false
	s.runCommand(s.model.query)
	s.model.query = ""
}
