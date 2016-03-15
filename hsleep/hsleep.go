package hsleep

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const chance_to_emit = 33
const sleep_for = 100 * time.Millisecond

func Hsleep(secs int) error {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	duration := time.Duration(secs) * time.Second
	deadline := time.Now().Add(duration)

	for now := time.Now(); deadline.After(now); now = time.Now() {
		if rng.Intn(100) <= chance_to_emit {
			time_left := deadline.Sub(now)
			fmt.Fprintf(os.Stderr, "\033[2K> %s\r", time_left.String())

			if time_left > sleep_for {
				time.Sleep(sleep_for)
			}
		}
	}
	fmt.Fprintf(os.Stderr, "\033[2K> Done!\n")
	return nil
}
