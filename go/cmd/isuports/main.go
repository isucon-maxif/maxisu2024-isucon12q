package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	isuports "github.com/isucon/isucon12-qualify/webapp/go"
)

func main() {
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()
	isuports.Run()
}
