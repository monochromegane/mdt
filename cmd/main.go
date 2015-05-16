package main

import (
	"fmt"
	"os"

	"github.com/monochromegane/mdt"
)

func main() {
	out, _ := mdt.Convert(os.Stdin)
	fmt.Printf("%s\n", out)
}
