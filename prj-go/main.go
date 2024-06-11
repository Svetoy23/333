package main

import (
	"fmt"
	// "math"
	// "math/rand"
)

// var c, b, a bool

// func add(x int, y int) (q1, q2, q3, q4 int) {
// 	q1 = x + y
// 	q2 = x - y
// 	q3 = x * y
// 	q4 = x / y
// 	return
// }

func main() {
	// fmt.Println("Ваш приз:", rand.Intn(1001), "₴!")
	// fmt.Println("Число Pi:", math.Pi)
	// fmt.Printf("Тепер у вас є %g друзів.\n", math.Sqrt(7))
	// fmt.Println(add(23, 77))
	// k := 3
	sum := 1
	count := 0
	for sum < 1000 {
		sum += sum
		count++
	}
	fmt.Println(sum, count)
}
