package main

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

func indexOfFirstBadWord2(msg []string, badWords []string) (string, int) {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return "Bad word: " + word + ", index:", i
			}
		}
	}
	return "Bad word not found - exit code:", -1
}

// don't touch below this line

func test(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println(` -`, badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("====================================")
}

func test2(msg []string, badWords []string) {
	info, index := indexOfFirstBadWord2(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println(` -`, badWord)
	}
	fmt.Println(info, index)
	fmt.Println("====================================")
}

func main() {
	badWords := []string{"crap", "shoot", "dang", "frick"}
	message := []string{"hey", "there", "john"}
	test(message, badWords)
	test2(message, badWords)

	message = []string{"ugh", "oh", "my", "frick"}
	test(message, badWords)
	test2(message, badWords)

	message = []string{"what", "the", "shoot", "I", "hate", "that", "crap"}
	test(message, badWords)
	test2(message, badWords)
}
