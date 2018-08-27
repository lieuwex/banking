package view

import (
	"banking/types"
)

func (s *ViewState) startTagging() {
	s.model.isTagging = true
	s.pages.ShowPage("tag-modal")
}

func (s *ViewState) finishTagging(tag string) {
	entry := s.table.GetSelected()
	entry.Tags = append(entry.Tags, tag)
	s.model.isTagging = false
	s.pages.HidePage("tag-modal")
}

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
	s.model.query = ""
	// TODO: execute
}
