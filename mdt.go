package mdt

import (
	"bufio"
	"io"
)

func Convert(header string, r io.Reader) (string, error) {

	table := newTable()

	if header != "" {
		table.addRow(header)
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		table.addRow(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return table.toMarkdown(), nil
}
