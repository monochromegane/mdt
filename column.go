package mdt

import "fmt"

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
