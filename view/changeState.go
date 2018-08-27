package view

import (
	"banking/types"
	"fmt"
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
	s.setCommandBarText("/", "")
}

func (s *ViewState) finishSearch(query string) {
	s.model.isSearching = false
	s.model.filter = query
}

func (s *ViewState) startCommand() {
	s.model.isCommanding = true
	s.setCommandBarText(":", "")
}

func (s *ViewState) finishCommand(cmd string) {
	s.model.isCommanding = false
	// TODO: execute
}

func (s *ViewState) setCommandBarText(prefix, str string) {
	str = fmt.Sprintf("\n%s%s", prefix, str)
	s.commandBar.SetText(str)
}
