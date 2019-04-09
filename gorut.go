package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Microsecond * 100)
	}
	//wg.Done() // what if something goes wrong and the routine doesnt return a done, use defer statement
}

func main() {
	wg.Add(1)
	go say("hi")
	wg.Add(1)
	go say("hello")
	wg.Wait()
}
