package main

import (
	"flag"
	"github.com/chooper/hsleep/hsleep"
	"strconv"
)

func main() {
	flag.Parse()

	if flag.Arg(0) == "" {
		panic("you must pass an argument!")
	}

	secs, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	hsleep.Hsleep(secs)
}
