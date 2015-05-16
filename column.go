package mdt

import (
	"fmt"
	"strings"
)

type columner interface {
	length() int
	toMarkdown(length int) string
}

type column struct {
	content string
}

func newColumn(s string) column {
	return column{
		content: s,
	}
}

func (c column) length() int {
	return len(c.content)
}

func (c column) toMarkdown(length int) string {
	format := fmt.Sprintf(" %%- %ds ", length)
	return fmt.Sprintf(format, c.content)
}

type divColumn struct {
	column
}

func newDivColumn(s string) divColumn {
	return divColumn{newColumn(s)}
}

func (d divColumn) toMarkdown(length int) string {
	return fmt.Sprintf(" %s ", strings.Repeat("-", length))
}
