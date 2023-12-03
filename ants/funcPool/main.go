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
	waitGroup := &sync.WaitGroup{}
	tasks := make([]int32, taskNum)
	for i, _ := range tasks {
		tasks[i] = rand.Int31n(100)
	}
	var result int32 = 0
	var expect int32 = 0
	for _, v := range tasks {
		expect += v
	}

	fp, _ := ants.NewPoolWithFunc(workerNum, func(i interface{}) {
		i.(func())() // 这种写法也是挺骚的
	})
	defer fp.Release()
	timeWrapper("ants funcPool", func() {
		for i := 0; i < taskNum; i++ {
			waitGroup.Add(1)
			pos := i // 此处有大坑
			_ = fp.Invoke(func() {
				atomic.AddInt32(&result, tasks[pos])
				waitGroup.Done()
			})
		}
		waitGroup.Wait()
	})
	fmt.Printf("result: %d, expect: %d\n", result, expect)
}
