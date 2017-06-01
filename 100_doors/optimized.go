package main

import "fmt"

func main() {
	for i := 0; i * i <= 100; i++ {
		fmt.Printf("door %d is open\n", i * i)
	}
}