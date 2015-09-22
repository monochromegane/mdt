package mdt_test

import (
	"fmt"
	"strings"

	"github.com/monochromegane/mdt"
)

func ExampleConvert_csv() {
	r := strings.NewReader(`
headerA, headerB
content, content
`)
	result, _ := mdt.Convert("", r)
	fmt.Printf("%s", result)

	// Output:
	// | headerA | headerB |
	// | ------- | ------- |
	// | content | content |
}

func ExampleConvert_tsv() {
	r := strings.NewReader(`
headerA	headerB
content	content
`)
	result, _ := mdt.Convert("", r)
	fmt.Printf("%s", result)

	// Output:
	// | headerA | headerB |
	// | ------- | ------- |
	// | content | content |
}

func ExampleConvert_format() {
	r := strings.NewReader(`
headerA, headerB
short, very very long content
`)
	result, _ := mdt.Convert("", r)
	fmt.Printf("%s", result)

	// Output:
	// | headerA | headerB                |
	// | ------- | ---------------------- |
	// | short   | very very long content |
}

func ExampleConvert_multibyte() {
	r := strings.NewReader(`
headerA, headerB
マルチバイト文字, content
ﾏﾙﾁﾊﾞｲﾄ文字, content
`)
	result, _ := mdt.Convert("", r)
	fmt.Printf("%s", result)

	// Output:
	// | headerA          | headerB |
	// | ---------------- | ------- |
	// | マルチバイト文字 | content |
	// | ﾏﾙﾁﾊﾞｲﾄ文字      | content |
}

func ExampleConvert_align() {
	r := strings.NewReader(`
headerA:, :headerB:
content, content
`)
	result, _ := mdt.Convert("", r)
	fmt.Printf("%s", result)

	// Output:
	// | headerA | headerB |
	// | -------:|:-------:|
	// | content | content |
}

func ExampleConvert_repeat() {
	r := strings.NewReader(`
| headerA | headerB |
| -------:|:-------:|
| content | content |
next content, next content
`)
	result, _ := mdt.Convert("", r)
	fmt.Printf("%s", result)

	// Output:
	// | headerA      | headerB      |
	// | ------------:|:------------:|
	// | content      | content      |
	// | next content | next content |
}

func ExampleConvert_short() {
	r := strings.NewReader(`
#,A
1,B
`)
	result, _ := mdt.Convert("", r)
	fmt.Printf("%s", result)

	// Output:
	// | #   | A   |
	// | --- | --- |
	// | 1   | B   |
}

func ExampleConvert_short_align() {
	r := strings.NewReader(`
#:,:A:
1,B
`)
	result, _ := mdt.Convert("", r)
	fmt.Printf("%s", result)

	// Output:
	// | #   | A   |
	// | ---:|:---:|
	// | 1   | B   |
}

func ExampleConvert_with_header() {
	r := strings.NewReader(`
content, content
`)
	result, _ := mdt.Convert("headerA,headerB", r)
	fmt.Printf("%s", result)

	// Output:
	// | headerA | headerB |
	// | ------- | ------- |
	// | content | content |
}

func ExampleConvert_with_header_align() {
	r := strings.NewReader(`
content, content
`)
	result, _ := mdt.Convert("headerA:,:headerB", r)
	fmt.Printf("%s", result)

	// Output:
	// | headerA | headerB |
	// | -------:|:------- |
	// | content | content |
}
