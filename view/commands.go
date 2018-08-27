package view

import "strings"

func (s *ViewState) runCommand(cmd string) {
	args := strings.Split(cmd, " ")

	switch args[0] {
	case "w", "write":
		panic("TODO: write")

	case "q", "quit":
		s.app.Stop()

	case "t", "tag":
		entry := s.table.GetSelected()
		entry.Tags = append(entry.Tags, args[1])

	default:
		panic("unknown command " + cmd)
	}
}
