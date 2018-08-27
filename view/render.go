package view

import (
	"banking/types"
	"banking/utils"
	"bytes"
	"strings"
	"time"

	"text/template"
)

func createTable(timeDelta time.Duration, days []types.Day) *CustomTable {
	date := utils.Today()
	res := MakeCustomTable()

	for i := 0; i < 365; i++ {
		d := date.Add(-1 * timeDelta)
		balanceDifference, balance, entries := getBetween(days, d, date)
		date = d

		var lastCell string
		if balanceDifference != 0 {
			lastCell = formatPrice(balanceDifference, true)
		}
		res.AddRow(
			nil,
			date.Format("2006-01-02"),
			"",
			"",
			formatPrice(balance, false),
			lastCell,
		)

		for _, entry := range entries {
			entryCopy := entry

			res.AddRow(
				entryCopy,
				"",
				formatPrice(entry.Amount, true),
				entry.Description,
			)
		}
	}

	return res
}

const templateStr = `
description: {{.Description}}

account: {{.Account}}
counter account: {{.CounterAccount}}

{{.Information}}

{{if .Tags}}
tags: {{join .Tags ", "}}
{{end}}
`

var infoBoxTemplate = template.Must(
	template.New("infoBox").
		Funcs(templateFuncs).
		Parse(templateStr),
)

func infoBoxString(entry *types.Entry) string {
	var buf bytes.Buffer
	infoBoxTemplate.Execute(&buf, entry)
	return strings.TrimSpace(buf.String())
}
