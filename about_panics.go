package go_koans

import "fmt"

func divideFourBy(i int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("HERE")
			fmt.Println(err)
		}
	}()
	return 4 / i
}

const __divisor__ = 0

func aboutPanics() {
	n := divideFourBy(__divisor__)
	fmt.Println(n)
	assert(n == 0) // panics are exceptional errors at runtime
}
