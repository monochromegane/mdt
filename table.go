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
	t.rows = append(t.rows, newRow(s))
}

func (t *table) toMarkdown() (string, error) {
	return t.rows.toMarkdown(), nil
}
