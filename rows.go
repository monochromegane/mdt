package mdt

import "bytes"

type rows []row

func (r rows) colNum() int {
	var max int
	for _, row := range r {
		if n := row.colNum(); n > max {
			max = n
		}
	}
	return max
}

func (r rows) lengthAt(index int) int {
	var max int
	for _, row := range r {
		if n := row.lengthAt(index); n > max {
			max = n
		}
	}
	if max < 3 {
		max = 3
	}
	return max
}

func (r rows) toMarkdown() string {
	colNum := r.colNum()
	lens := []int{}
	for i := 0; i < colNum; i++ {
		lens = append(lens, r.lengthAt(i))
	}

	var buf bytes.Buffer
	for _, row := range r {
		buf.WriteString(row.toMarkdown(colNum, lens) + "\n")
	}
	return buf.String()
}
