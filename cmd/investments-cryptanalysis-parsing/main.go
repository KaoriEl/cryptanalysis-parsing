package main

import (
	"main/internal/Server"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		Server.StartServer()
	}()
	wg.Wait()
}
