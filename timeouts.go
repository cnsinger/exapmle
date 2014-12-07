package main

import "fmt"
import "time"

func f(c1 chan<- string) {
    time.Sleep(time.Second * 2)
    c1 <- "result 1"
}

func main() {
    c1 := make(chan string, 1)
    go f(c1)

    select {
    case res := <- c1:
        fmt.Println(res)
    case <- time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }

    select {
    case res := <- c1:
        fmt.Println(res)
    case <- time.After(time.Second * 3):
        fmt.Println("timeout 3")
    }
}
