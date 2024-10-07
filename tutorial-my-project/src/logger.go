package main

import (
	"log"
	"path/filepath"
	"runtime"
)

func logger(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fileName := filepath.Base(file)
		log.Printf("[%s:%d] %s", fileName, line, msg)
	} else {
		log.Printf("%s", msg)
	}
}
