package mdt

import (
	"bufio"
	"io"
)

func Convert(r io.Reader) (string, error) {

	table := newTable()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		table.addRow(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return table.toMarkdown(), nil
}
