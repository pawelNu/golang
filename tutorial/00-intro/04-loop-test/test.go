package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var counter int = 0
	for i := 0; i <= 1_000_000_000; i++ {
		// loop from 0 to 1 000 000 000
		counter++
	}

	duration := time.Since(start)
	fmt.Println("Counter: ", counter)
	fmt.Printf("Duration time: %.2f seconds\n", duration.Seconds())
}
