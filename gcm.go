package main

import (
	"flag"
	"fmt"
	"os"
)

const gcmVersion = "0.0.1"

func main() {
	var version = flag.Bool("v", false, "print the current version of gcm")
	flag.Parse()

	if *version {
		fmt.Println(gcmVersion)
		os.Exit(0)
	}
}
