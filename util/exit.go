package util

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

func Exit() {
	spinner := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	spinner.Suffix = "Exiting..."
	spinner.Start()

	time.Sleep(3 * time.Second)

	spinner.Stop()
	fmt.Println("Bye! 👋")
}
