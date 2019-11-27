package main

import (
	"fmt"
	"os"
	"time"

	"github.com/maurerlabs/github-actions/toolkit"
)

func main() {
	name := os.Getenv("INPUT_WHO_TO_GREET")

	fmt.Printf("Hello, %s!\n", name)
	toolkit.Output("time", time.Now())
}
