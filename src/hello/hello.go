package main

import (
       "fmt"
       "stringutil"
       "time"
       "math"
       "math/cmplx"

)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("hello, world\n")
        fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	fmt.Println("The time is", time.Now())
	fmt.Println(stringutil.Reverse(time.Now().String()))
        fmt.Println(math.Pi)
	piStr := fmt.Sprintf("%g", math.Pi)
	fmt.Println(stringutil.Reverse(piStr))
        fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	g := 0.867 + 0.5i // complex128
        fmt.Println(x, y, z, g)


}

