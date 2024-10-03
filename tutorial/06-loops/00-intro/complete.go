package main

import (
	"fmt"
)

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1 + (float64(i) * 0.01)
	}
	return totalCost
}

// don't edit below this line

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func simpleForLoop() {
	fmt.Println("simpleForLoop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func simpleForLoop2() {
	fmt.Println("simpleForLoop2")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func infiniteForLoop() {
	fmt.Println("infiniteForLoop")
	for {
		// This will run indefinitely
		fmt.Println("This will never end!")
		break // Use break to exit the loop
	}
}

func forLoopWithCondition() {
	// similar to while loop
	fmt.Println("forLoopWithCondition")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func forLoopWithRangeOverSlice() {
	fmt.Println("forLoopWithRangeOverSlice")
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

func forLoopWithRangeOverMap() {
	fmt.Println("forLoopWithRangeOverMap")
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

func forLoopWithRangeOverString() {
	fmt.Println("forLoopWithRangeOverString")
	str := "hello"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}

func forLoopWithContinue() {
	fmt.Println("forLoopWithContinue")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println(i) // Print only odd numbers
	}
}

func forLoopWithBreak() {
	fmt.Println("forLoopWithBreak")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // Exit the loop when i equals 5
		}
		fmt.Println(i)
	}
}

func main() {
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)
	simpleForLoop()
	simpleForLoop2()
	infiniteForLoop()
	forLoopWithCondition()
	forLoopWithRangeOverSlice()
	forLoopWithRangeOverMap()
	forLoopWithRangeOverString()
	forLoopWithContinue()
	forLoopWithBreak()
}
