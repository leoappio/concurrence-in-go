package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		imprimir("goroutine 1")
		waitGroup.Done()
	}()

	go func() {
		imprimir("goroutine 2")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func imprimir(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
