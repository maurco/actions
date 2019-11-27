package main

import (
	"os"
	"time"

	"github.com/maurerlabs/actions/toolkit"
)

func main() {
	name := os.Getenv("INPUT_WHO_TO_GREET")

	toolkit.Info("Hello, %s!", name)
	toolkit.SetOutput("time", time.Now())
}
