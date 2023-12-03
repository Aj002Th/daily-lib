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
	var result int32 = 0
	var expect int32 = 0
	tasks := make([]int32, taskNum)
	for i, _ := range tasks {
		tasks[i] = rand.Int31n(100)
	}
	for _, v := range tasks {
		expect += v
	}

	waitGroupForAnts := &sync.WaitGroup{}
	pool, _ := ants.NewPool(workerNum)
	defer pool.Release()
	timeWrapper("ants pool", func() {
		for i := 0; i < taskNum; i++ {
			waitGroupForAnts.Add(1)
			pos := i // 此处有大坑
			_ = pool.Submit(func() {
				atomic.AddInt32(&result, tasks[pos])
				waitGroupForAnts.Done()
			})
		}
		waitGroupForAnts.Wait()
	})

	fmt.Printf("result: %d, expect: %d\n", result, expect)
}
