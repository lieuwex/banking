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
