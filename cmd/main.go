package main

import (
	"fmt"
	"log"
	"os"

	"github.com/monochromegane/mdt"
)

func main() {
	out, err := mdt.Convert(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}
