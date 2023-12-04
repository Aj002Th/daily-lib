package main

import (
	"github.com/panjf2000/ants/v2"
)

func PanicFunc() {
	panic("panic in PanicFunc")
}

func main() {
	p, _ := ants.NewPool(10)
	defer p.Release()

	for i := 0; i < 3; i++ {
		_ = p.Submit(PanicFunc)
	}

	// 打印三条 panic 但是程序不会退出
	select {}
}
