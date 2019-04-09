package main

import (
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func cleanup () {
	if r := recover(); r != nil {
		fmt.Println("Recovered cleanup =>", r)
	}
	wg2.Done()
}

func say2(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Microsecond*100)
		fmt.Println(s)
		if i == 2 {
			panic("panic alert for 2")
		}
	}
}

func main() {
	wg2.Add(1)
	go say2("hi")
	wg2.Add(1)
	go say2("hello")
	wg2.Wait()
}
