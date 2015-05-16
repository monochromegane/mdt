package mdt

import "strings"

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
	// add row.
	t.rows = append(t.rows, newRow(s))

	// add division row.
	if len(t.rows) == 1 {
		t.rows = append(t.rows, newDivRow(""))
	}
}

func (t *table) toMarkdown() (string, error) {
	return t.rows.toMarkdown(), nil
}
