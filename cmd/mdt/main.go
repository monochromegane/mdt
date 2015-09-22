package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/monochromegane/mdt"
)

var header string

func init() {
	flag.StringVar(&header, "H", "", "Specify a header line. If none specified, then first line is used as a header line.")
	flag.Parse()
}

func main() {
	out, err := mdt.Convert(header, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}
