package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	taskNum   = 10000000
	workerNum = 10000
)

func timeWrapper(name string, f func()) {
	start := time.Now()
	f()
	fmt.Printf("[%s] cost time: %d ms\n", name, time.Since(start).Milliseconds())
}

func main() {
	var resultForGoroutine int32 = 0
	var resultForAnts int32 = 0
	var expect int32 = 0
	tasks := make([]int32, taskNum)
	for i, _ := range tasks {
		tasks[i] = rand.Int31n(10)
	}

	timeWrapper("for range", func() {
		for _, v := range tasks {
			expect += v
		}
	})

	waitGroupForGoroutine := &sync.WaitGroup{}
	timeWrapper("goroutine", func() {
		for i := 0; i < taskNum; i++ {
			waitGroupForGoroutine.Add(1)
			pos := i // 此处有大坑
			go func() {
				atomic.AddInt32(&resultForGoroutine, tasks[pos])
				waitGroupForGoroutine.Done()
			}()
		}
		waitGroupForGoroutine.Wait()
	})

	waitGroupForAnts := &sync.WaitGroup{}
	pool, _ := ants.NewPool(workerNum)
	defer pool.Release()
	timeWrapper("ants pool", func() {
		for i := 0; i < taskNum; i++ {
			waitGroupForAnts.Add(1)
			pos := i // 此处有大坑
			_ = pool.Submit(func() {
				atomic.AddInt32(&resultForAnts, tasks[pos])
				waitGroupForAnts.Done()
			})
		}
		waitGroupForAnts.Wait()
	})

	fmt.Printf("resultForAnts: %d, resultForGoroutine: %d, expect: %d\n", resultForAnts, resultForGoroutine, expect)
}
