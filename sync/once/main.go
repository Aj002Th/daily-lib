package main

import (
	"sync"
	"time"
)

func doSomething() {
	println("do something")
}

func main() {
	once := sync.Once{}
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(doSomething)
		}()
	}
	time.Sleep(time.Second * 3)
}
