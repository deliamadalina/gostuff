package main

import (  
    "fmt"
    "time"
)
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
func hello(done chan bool) {  
    fmt.Println("hello go routine is going to sleep")
    time.Sleep(4 * time.Second)
    fmt.Println("hello go routine awake and going to write to done")
    done <- true
}
func newHello(done chan bool) {
    fmt.Println("new hello go routine is going to sleep")
    time.Sleep(1 * time.Second)
    fmt.Println("new hello go routine awake and going to write to done")
    done <- true
}
func newerHello(done chan bool) {
    fmt.Println("newer hello go routine is going to sleep")
    for _, v := range pow {
        time.Sleep(1 * time.Second)
	fmt.Println(v)
    }
    fmt.Println("newer hello go routine awake and going to write to done")
    done <- true
}
func main() {  
    done := make(chan bool)
    newdone :=make(chan bool)
    newerdone :=make(chan bool)
    fmt.Println("Main going to call hello go goroutine")
    go hello(done)
    go newHello(newdone)
    go newerHello(newerdone)
    <-done
    <-newdone
    <-newerdone
    fmt.Println("Main received data")
}
