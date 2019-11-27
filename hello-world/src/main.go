package main

import (
	"time"

	"github.com/maurerlabs/actions/toolkit"
)

func main() {
	name := toolkit.GetInput("who_to_greet")

	toolkit.Info("Hello, %s!", name)
	toolkit.SetOutput("time", time.Now())
}
