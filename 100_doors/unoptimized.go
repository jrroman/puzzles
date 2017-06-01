package main

import "fmt"

func main() {
	var doors []bool = make([]bool, 100)

	for pass := 1; pass <= 100; pass++ {
		for door := pass - 1; door < 100; door += pass {
			doors[door] = !doors[door]
		}
	}

	for idx, _ := range doors {
		if !doors[idx] {
			fmt.Printf("door %d is closed\n", idx + 1)
		} else {
			fmt.Printf("door %d is open\n", idx + 1)
		}
	}
}