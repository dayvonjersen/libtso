package main

import (
	"flag"
)

var (
	some_flag bool
)

func main() {
	flag.BoolVar(
		&some_flag,
		"flag-name",
		false, // default value
		"A short description",
	)

	flag.Parse()
}
