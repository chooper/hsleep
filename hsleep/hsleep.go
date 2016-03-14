package hsleep

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Hsleep(secs int) error {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	duration := time.Duration(secs) * time.Second
	deadline := time.Now().Add(duration)

	for deadline.After(time.Now()) {
		if rng.Intn(100) > 33 {
			time_left := deadline.Sub(time.Now())
			fmt.Fprintf(os.Stderr, "> %s\r", time_left.String())
		}
	}
	// FIXME(charles) Find a better way to clear the line
	// Terminal escape \033[2K ???
	fmt.Fprintf(os.Stderr, "                  \r")
	fmt.Fprintf(os.Stderr, "> Done!\n")
	return nil
}
