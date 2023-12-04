package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
)

func PanicFunc() {
	panic("panic in PanicFunc")
}

func PanicHandler(i interface{}) {
	fmt.Printf("custom handler is handling: %v\n", i)
}

func main() {
	p, _ := ants.NewPool(10, ants.WithPanicHandler(PanicHandler))
	defer p.Release()

	for i := 0; i < 3; i++ {
		_ = p.Submit(PanicFunc)
	}

	// 打印三条 panic 但是程序不会退出
	// 会使用自定义的 panic handler
	select {}
}
