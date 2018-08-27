package view

import "strings"

func (s *ViewState) runCommand(cmd string) {
	if cmd == "" {
		return
	}

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
	case "ta", "tag-all":
		query := args[1]
		tag := args[2]
		for _, entry := range s.model.entries {
			if entry.Matches(query) {
				entry.Tags = append(entry.Tags, tag)
			}
		}

	default:
		panic("unknown command " + cmd)
	}
}
