# mdt [![Build Status](https://travis-ci.org/monochromegane/mdt.svg?branch=master)](https://travis-ci.org/monochromegane/mdt)

A markdown table generation tool from CSV/TSV.

## Usage

```sh
$ mdt < hoge.csv
| headerA | headerB                |
| ------- | ---------------------- |
| short   | very very long content |
```

`mdt` is very simple command line tool, so you can integrate with other tools like `pbpaste`, `Automator`, etc...

For example, you can use `mdt` with `Automator Service` on GitHub PR/Issue page.

![mdt](https://cloud.githubusercontent.com/assets/1845486/7668803/cc0a9178-fc87-11e4-9d0e-9fd32ea3c1fc.gif)

## Features

- Convert from CSV/TSV.
- Auto formatting.
- Support multibyte contents.
- Define align at CSV/TSV header ex. `:headerA:, headerB:`
- Repeating execution.
- Specify a header line (`-H` option).

See also [examples](https://godoc.org/github.com/monochromegane/mdt#pkg-examples).

## Installation

```sh
$ go get github.com/monochromegane/mdt/...
```

## Integration

### Automator

![mdt\_with\_automator](https://cloud.githubusercontent.com/assets/1845486/7668851/5d833f84-fc8c-11e4-8787-aa39ce6ab300.png)

1. Create new Automator Service.
2. Select `Run Shell Script` action.
3. `Service receives selected` -> `text`
4. Check `Output replaces selected text`
5. Input `/path/to/mdt` at `Run Shell Script` (`pass input` is `to stdin`).

You can call the service by shortcut key.

`System Preferences` > `Keyboard` > `Shortcuts` tab > `Services` > Select your service, and set shortcut.

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

