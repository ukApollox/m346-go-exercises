package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(6) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "The dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stderr, "The dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	// m346-go-exercises\go-1-ex-3> go run .\main.go > eyes.txt 2> dice.log
}
