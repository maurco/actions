package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	name := os.Getenv("INPUT_WHO_TO_GREET")

	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("::set-output name=time::%s\n", time.Now())
}
