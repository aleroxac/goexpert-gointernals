package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2) // Cria um channel bufferizado com capacidade de 2

	go func() {
		ch <- 1
		fmt.Println("Sent 1")
		ch <- 2
		fmt.Println("Sent 2")
		ch <- 3
		fmt.Println("Sent 3") // Esta operação vai bloquear até que haja espaço no buffer
	}()
	time.Sleep(time.Second)

	go func() {
		val := <-ch
		fmt.Println("Received", val)
	}()
	time.Sleep(2 * time.Second)
}
