package mdt

import (
	"regexp"
	"strings"
)

type table struct {
	rows rows
}

func newTable() table {
	return table{}
}

func (t *table) addRow(s string) {
	s = strings.TrimSpace(s)
	// skip blank row.
	if s == "" {
		return
	}

	if len(t.rows) == 0 {
		// header and division rows.
		row, divRow := newHeaderRow(s)
		t.rows = append(t.rows, row)
		t.rows = append(t.rows, divRow)
	} else {
		if isDivRow(s) {
			// over write division row.
			t.rows[1] = newDivRow(s)
		} else {
			// content row.
			t.rows = append(t.rows, newRow(s))
		}
	}
}

func (t *table) toMarkdown() string {
	return t.rows.toMarkdown()
}

var regDivRow = regexp.MustCompile("^[\\s]*\\|[\\s]*:*--+[\\s]*:*\\|")

func isDivRow(s string) bool {
	return regDivRow.MatchString(s)
}
