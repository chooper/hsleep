package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/chooper/hsleep/hsleep"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()

	if flag.Arg(0) == "" {
		return errors.New("Missing required argument!\n\nUsage: hsleep <secs>")
	}

	secs, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		return err
	}

	return hsleep.Hsleep(secs)
}
