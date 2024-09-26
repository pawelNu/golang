package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

type sender2 struct {
	rateLimit int
	user      user2
}

type user2 struct {
	name   string
	number int
}

// don't edit below this line

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func test2(s sender2) {
	fmt.Println("Sender name:", s.user.name)
	fmt.Println("Sender number:", s.user.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func main() {

	test(sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 18055558790,
		},
	})

	test(sender{
		rateLimit: 5000,
		user: user{
			name:   "Sarah",
			number: 19055558790,
		},
	})

	test(sender{
		rateLimit: 1000,
		user: user{
			name:   "Sally",
			number: 19055558790,
		},
	})

	test2(sender2{
		rateLimit: 9999,
		user: user2{
			name:   "test",
			number: 123455,
		},
	})
}
