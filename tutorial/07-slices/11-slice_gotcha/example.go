package main

import "fmt"

func example1() {
	// it works because initial capacity of slice a is the same as length = 3
	// and adding 4th element will create new underlying array
	a := make([]int, 3)
	fmt.Println("len of a:", len(a))
	// len of a: 3
	fmt.Println("cap of a:", cap(a))
	// cap of a: 3
	fmt.Println("appending 4 to b from a")
	// appending 4 to b from a
	b := append(a, 4)
	fmt.Println("b:", b)
	// b: [0 0 0 4]
	fmt.Println("addr of b:", &b[0])
	// addr of b: 0x44a0c0
	fmt.Println("appending 5 to c from a")
	// appending 5 to c from a
	c := append(a, 5)
	fmt.Println("addr of c:", &c[0])
	// addr of c: 0x44a180
	fmt.Println("a:", a)
	// a: [0 0 0]
	fmt.Println("b:", b)
	// b: [0 0 0 4]
	fmt.Println("c:", c)
	// c: [0 0 0 5]
}

func example2() {
	// appending 5 and initializing slice g cause mutation of j
	// because slice i has length 3 and capacity 8 and the underlying array can grow in size
	i := make([]int, 3, 8)
	fmt.Println("len of i:", len(i))
	// len of i: 3
	fmt.Println("cap of i:", cap(i))
	// cap of i: 8
	fmt.Println("appending 4 to j from i")
	// appending 4 to j from i
	j := append(i, 4)
	fmt.Println("j:", j)
	// j: [0 0 0 4]
	fmt.Println("addr of j:", &j[0])
	// addr of j: 0x454000
	fmt.Println("appending 5 to g from i")
	// appending 5 to g from i
	g := append(i, 5)
	fmt.Println("addr of g:", &g[0])
	// addr of g: 0x454000
	fmt.Println("i:", i)
	// i: [0 0 0]
	fmt.Println("j:", j)
	// j: [0 0 0 5]
	fmt.Println("g:", g)
	// g: [0 0 0 5]
}

func main() {
	fmt.Println("Example 1")
	example1()
	fmt.Println()
	fmt.Println("Example 2")
	example2()
}
