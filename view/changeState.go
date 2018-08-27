package view

func (s *ViewState) startSearch() {
	s.model.isSearching = true
}

func (s *ViewState) finishSearch() {
	s.model.isSearching = false
	s.model.filter = s.model.query
}

func (s *ViewState) startCommand() {
	s.model.isCommanding = true
}

func (s *ViewState) finishCommand() {
	s.model.isCommanding = false
	s.runCommand(s.model.query)
}
