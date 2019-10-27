package main

import (
       "fmt"
       "stringutil"
       "time"
       "math"
)

func main() {
	fmt.Printf("hello, world\n")
        fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	fmt.Println("The time is", time.Now())
	fmt.Println(stringutil.Reverse(time.Now().String()))
        fmt.Println(math.Pi)
	piStr := fmt.Sprintf("%g", math.Pi)
	fmt.Println(stringutil.Reverse(piStr))
}

