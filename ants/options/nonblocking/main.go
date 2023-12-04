package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
)

func main() {
	p, _ := ants.NewPool(2, ants.WithNonblocking(true))
	defer p.Release()

	for i := 0; i < 3; i++ {
		pos := i
		err := p.Submit(func() {
			fmt.Printf("hello world: %v\n", pos)
		})
		if err != nil {
			fmt.Println(err)
		}
	}

	// 会有任务因为协程池满了而提交失败
	select {}
}
