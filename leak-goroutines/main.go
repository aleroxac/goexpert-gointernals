package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var leakedGoroutines int32

func leakedGoroutine() {
	atomic.AddInt32(&leakedGoroutines, 1)
	for {
		// loop infinito representando o vazamento
	}
}

func leakHendler(w http.ResponseWriter, r *http.Request) {
	go leakedGoroutine() // Inicia uma nova goroutine que vai vazar
	fmt.Fprintf(w, "Leaked a goroutine!\n")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Retorna o número atual de goroutines vazadas
	fmt.Fprintf(w, "Current number of leaked goroutines: %d\n", atomic.LoadInt32(&leakedGoroutines))
}

func main() {
	http.HandleFunc("/leak", leakHendler)     // Endpoint para causar vazamento
	http.HandleFunc("/status", statusHandler) // Endpoint para verificar o estado do vazamento

	fmt.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
