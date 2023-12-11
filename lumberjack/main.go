package main

import (
	"github.com/natefinch/lumberjack"
	"log"
)

func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./lumberjack/foo.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	for i := 0; i < 100000; i++ {
		log.Println("hello world")
	}
}
