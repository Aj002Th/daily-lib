package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"time"
)

func main() {
	p, _ := ants.NewPool(3, ants.WithMaxBlockingTasks(2))
	defer p.Release()

	for i := 0; i < 10; i++ {
		// 这里为什么要启动新协程来 submit 才能触发 WithMaxBlockingTasks 的限制呢?
		// 因为容量已满但是等待执行的任务未到达 MaxBlockingTasks 限制时,
		// 在阻塞模式下, submit 会阻塞住
		// 导致后面的任务都无法继续 submit,
		// 所以需要启动新协程来 submit 才能触发 WithMaxBlockingTasks 的限制

		pos := i
		go func() {
			err := p.Submit(func() {
				fmt.Printf("hello world: %v\n", pos)
				time.Sleep(time.Second)
			})
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	// 会有任务因为等待执行的任务过多而提交失败
	select {}
}
