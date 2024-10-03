package main

import (
	"fmt"
)

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) (int, []float64) {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0
	costs := []float64{}
	for actualCostInPennies <= float64(maxCostInPennies) {
		costs = append(costs, actualCostInPennies)
		maxMessagesToSend++
		actualCostInPennies *= costMultiplier
	}
	return maxMessagesToSend, costs
}

// don't touch below this line

func test(costMultiplier float64, maxCostInPennies int) {
	maxMessagesToSend, costs := getMaxMessagesToSend(costMultiplier, maxCostInPennies)
	fmt.Printf("Multiplier: %v\n",
		costMultiplier,
	)
	fmt.Printf("Max cost: %v\n",
		maxCostInPennies,
	)
	fmt.Printf("Max messages you can send: %v\n",
		maxMessagesToSend,
	)
	fmt.Println("Costs of each message:")
	for i, cost := range costs {
		fmt.Printf("Message %d: %.2f\n", i+1, cost)
	}
	fmt.Println("====================================")
}

func main() {
	test(1.1, 5)
	test(1.3, 10)
	test(1.35, 25)
}
