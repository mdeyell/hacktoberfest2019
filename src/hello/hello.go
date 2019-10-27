package main

import (
       "fmt"
       "stringutil"
       "time"

)

func main() {
	fmt.Printf("hello, world\n")
        fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	fmt.Println("The time is", time.Now())
	fmt.Println(stringutil.Reverse(time.Now().String()))

}

