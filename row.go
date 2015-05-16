package mdt

import (
	"fmt"
	"strings"
)

type row struct {
	columns   []columner
	newColumn newColumnFunc
}

type newColumnFunc func(s string) columner

func newRow(s string) row {
	row := row{newColumn: func(s string) columner { return newColumn(s) }}
	row.setRow(s)
	return row
}

func newDivRow(s string) row {
	row := row{newColumn: func(s string) columner { return newDivColumn(s) }}
	row.setRow(s)
	return row
}

func (r *row) setRow(s string) {
	s = strings.Trim(s, "|")
	s = strings.Replace(s, "|", "\t", -1)
	s = strings.Replace(s, ",", "\t", -1)

	for _, col := range strings.Split(s, "\t") {
		r.columns = append(r.columns, r.newColumn(col))
	}
}

func (r row) colNum() int {
	return len(r.columns)
}

func (r row) lengthAt(index int) int {
	if l := len(r.columns) - 1; index > l {
		return 0
	} else {
		return r.columns[index].length()
	}
}

func (r *row) fillColumn(colNum int) {
	short := colNum - len(r.columns)
	for i := 0; i < short; i++ {
		r.columns = append(r.columns, r.newColumn(""))
	}
}

func (r row) toMarkdown(colNum int, lens []int) string {
	r.fillColumn(colNum)
	columns := []string{}
	for i, c := range r.columns {
		columns = append(columns, c.toMarkdown(lens[i]))
	}

	return fmt.Sprintf("|%s|", strings.Join(columns, "|"))
}
