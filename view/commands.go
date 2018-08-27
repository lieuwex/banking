package view

import "strings"

func (s *ViewState) runCommand(cmd string) {
	args := strings.Split(cmd, " ")
	entry := s.table.GetSelected()

	switch args[0] {
	case "w", "write":
		panic("TODO: write")

	case "q", "quit":
		s.app.Stop()

	case "t", "tag":
		entry.Tags = append(entry.Tags, args[1])
	case "rt", "remove-tag":
		index := -1
		for i, tag := range entry.Tags {
			if tag == args[1] {
				index = i
				break
			}
		}

		if index > -1 {
			entry.Tags = append(entry.Tags[:index], entry.Tags[index+1:]...)
		}

	default:
		panic("unknown command " + cmd)
	}
}
