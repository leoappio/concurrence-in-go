package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go imprimir("teste", canal)

	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
}

func imprimir(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
