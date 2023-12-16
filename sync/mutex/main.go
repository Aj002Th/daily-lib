package main

import (
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	cnt := 0
	for i := 0; i < 10; i++ {
		go func() {
			m.Lock()
			cnt++
			m.Unlock()
		}()
	}

	time.Sleep(time.Second) // 保证所有协程执行完	fmt.Println(cnt)
}
