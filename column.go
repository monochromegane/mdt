package mdt

import (
	"fmt"
	"strings"
	"unicode/utf8"
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
	var length int
	for _, s := range c.content {
		if l := utf8.RuneLen(s); l > 1 {
			length += 2
		} else {
			length++
		}
	}
	return length
}

func (c column) multiByteNum() int {
	var num int
	for _, s := range c.content {
		if l := utf8.RuneLen(s); l > 1 {
			num++
		}
	}
	return num
}

func (c column) toMarkdown(length int) string {
	format := fmt.Sprintf(" %%- %ds ", length-c.multiByteNum())
	return fmt.Sprintf(format, c.content)
}

type align int

const (
	left align = iota
	center
	right
	none
)

type divColumn struct {
	column
	align align
}

func newDivColumn(s string) divColumn {
	s = strings.TrimSpace(s)
	align := none
	switch {
	case strings.HasPrefix(s, ":") && strings.HasSuffix(s, ":"):
		align = center
	case strings.HasPrefix(s, ":"):
		align = left
	case strings.HasSuffix(s, ":"):
		align = right
	}
	return divColumn{
		column: newColumn(strings.Trim(s, ":")),
		align:  align,
	}
}

func (d divColumn) toMarkdown(length int) string {
	prefix := " "
	suffix := " "
	switch d.align {
	case left:
		prefix = ":"
	case center:
		prefix = ":"
		suffix = ":"
	case right:
		suffix = ":"
	}
	return fmt.Sprintf("%s%s%s", prefix, strings.Repeat("-", length), suffix)
}
