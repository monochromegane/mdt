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
		t.rows = append(t.rows, newRow(s))
		t.rows = append(t.rows, newDivRow(""))
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

func (t *table) toMarkdown() (string, error) {
	return t.rows.toMarkdown(), nil
}

var regDivRow = regexp.MustCompile("^[\\s]*\\|[\\s]*:*--+[\\s]*:*\\|")

func isDivRow(s string) bool {
	return regDivRow.MatchString(s)
}
