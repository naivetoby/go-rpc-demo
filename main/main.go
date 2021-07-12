package main

import (
	"github.com/thinktkj/go-rpc-demo/internal/api"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		api.GrpcEnable(wg)
	}()

	wg.Wait()
}
