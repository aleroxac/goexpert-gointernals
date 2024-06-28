package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42 // Envia um valor para o channel
	}()

	value := <-ch // Recebe o valor do channel
	fmt.Println(value)
}
