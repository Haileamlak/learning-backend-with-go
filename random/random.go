package main

import (
	"fmt"
	// "math"
	// "slices"
)

var x int = 10

func main() {
	// var i int = 20
	// var f float64 = float64(i)
	// fmt.Print(i, f)
	// const value = 10
	// var i = value
	// var f = value
	// fmt.Print(i, f)
	// var b byte = 255
	// b++
	// var smallI int32 = int32(math.Pow(2, 31) + 1)
	// smallI++
	// var bigI uint64 = uint64(math.Pow(2, 644) - 1)

	// fmt.Print(b, smallI, bigI)
	// var x = []int {1, 2, 4, 4, 5}
	// var y = []int {1, 2, 4, 4, 5}
	// fmt.Println(slices.Equal(x, y))
	// var x []int
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 10)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 20)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 30)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 40)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 50)
	// clear(x)
	// fmt.Println(x, len(x), cap(x))
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

}
