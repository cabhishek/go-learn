package main

import (
    "fmt"
)

func fib(c chan int, n int) {

    a, b := 0, 1

    for i:= 0; i < n; i++ {
        a, b = b, b+a

        c <- a
    }

    close(c)
}

func main(){

    c := make(chan int)

    go fib(c, 10)

    for n := range c {
        fmt.Println(n)
    }
}
