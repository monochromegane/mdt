# mdt [![Build Status](https://travis-ci.org/monochromegane/mdt.svg?branch=master)](https://travis-ci.org/monochromegane/mdt)

A markdown table generation tool from CSV/TSV.

## Usage

```sh
$ mdt < hoge.csv
| headerA | headerB                |
| ------- | ---------------------- |
| short   | very very long content |
```

## Features

- Convert from CSV/TSV.
- Auto formatting.
- Support multibyte contents.
- Define align at CSV/TSV header ex. `:headerA:, headerB:`
- Repeating execution.

See also [examples](https://godoc.org/github.com/monochromegane/mdt#pkg-examples).

## Installation

```sh
$ go get github.com/monochromegane/mdt/...
```

## Contribution

1. Fork it
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request

## License

[MIT](https://github.com/monochromegane/mdt/blob/master/LICENSE)

## Author

[monochromegane](https://github.com/monochromegane)

