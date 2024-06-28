package main

import "fmt"

func main() {
	ch := make(chan int, 2) // Cria um channel bufferizado com capacidade de 2

	ch <- 42 // Envia o primeiro valor para o channel
	ch <- 43 // Envia o segundo valor para o channel

	fmt.Println(<-ch) // Recebe e imprime o primeiro valor
	fmt.Println(<-ch) // Recebe e imprime o segundo valor
}
