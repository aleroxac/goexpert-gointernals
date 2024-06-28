package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {
	// Habilitar o tracing do GC
	// debug.SetGCPercent(-1) // Desativa o GC automático

	// Ajustar o percentual do GC
	// debug.SetGCPercent(300)

	// Definindo um limite de memória de 10MB
	debug.SetMemoryLimit(10 * 1024 * 1024)

	// Função para alocar memória
	allocateMemory := func(size int) []byte {
		return make([]byte, size)
	}

	// Allocando memória para observar o comportamento do GC
	for i := 0; i < 10; i++ {
		_ = allocateMemory(20 * 1024 * 1024) // Alocando 20MB
		time.Sleep(100 * time.Millisecond)

		// Exibindo o uso de memória
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc - %v MiB\n", m.Alloc/1024/1024)               // memoria atualmente aalocada no heap
		fmt.Printf("TotalAlloc - %v MiB\n", m.TotalAlloc/1024/1024)     // total acumulado desde o inicio do programa
		fmt.Printf("Sys - %v MiB\n", m.Sys/1024/1024)                   // total de memoria do sistema asolicitada pelo runtime do go
		fmt.Printf("Lookups - %v\n", m.Lookups)                         // numero de buscas de ponteiros de referência(geralmente usado com cgo)
		fmt.Printf("Mallocs - %v\n", m.Mallocs)                         // numeo de chamadas de alocacoes de memoria desde o inicio do programa
		fmt.Printf("Frees - %v\n", m.Frees)                             // numero de chamadas de liberacao de memoria desde o inicio do programa
		fmt.Printf("HeapAlloc - %v MiB\n", m.HeapAlloc/1024/1024)       // memoria do heap atualmente em uso
		fmt.Printf("HeapSys - %v MiB\n", m.HeapSys/1024/1024)           // memoria do heap solicitada ao sistema operacional
		fmt.Printf("HeapIdle - %v MiB\n", m.HeapIdle/1024/1024)         // memoria do heap alocada mas ociosa
		fmt.Printf("HeapInuse - %v MiB\n", m.HeapInuse/1024/1024)       // memoria do heap ativamente usada
		fmt.Printf("HeapReleased - %v MiB\n", m.HeapReleased/1024/1024) // memoria do heap liberada de volta para o sistema operacional
		fmt.Printf("HeapObjects - %v MiB\n", m.HeapObjects/1024/1024)   // numero de objetos atualmente alocados no heap
		fmt.Printf("StackInuse - %v MiB\n", m.StackInuse/1024/1024)     // memoria usada pelas pulhas de goroutines
		fmt.Printf("StackSys - %v MiB\n", m.StackSys/1024/1024)         // memoria solicitada pelas pulhas de goroutines
		fmt.Printf("MSpanInuse - %v MiB\n", m.MSpanInuse/1024/1024)     // memoria usada para descritores de spans atualmente em uso
		fmt.Printf("MSpanSys - %v MiB\n", m.MSpanSys/1024/1024)         // memoria solicitada para descritores de spans
		fmt.Printf("MCacheInuse - %v MiB\n", m.MCacheInuse/1024/1024)   // memoria usada para caches de objetos atualmente em uso
		fmt.Printf("MCacheSys - %v MiB\n", m.MCacheSys/1024/1024)       // memoria solicitada para caches de objetos
		fmt.Printf("BuckHashSys - %v MiB\n", m.BuckHashSys/1024/1024)   // memoria usada para tabelas de hash internas
		fmt.Printf("GCSys - %v MiB\n", m.GCSys/1024/1024)               // memoria usada pelo garbace collector
		fmt.Printf("OtherSys - %v MiB\n", m.OtherSys/1024/1024)         // memoria usada para outros propositos pelo runtime
		fmt.Println("-------------------------------------------")
	}
}

// -----HOW TO RUN
// uso do GC com alocação padrão
// GODEBUG=gctrace=1 go run main.go

// uso do GC em 300% de alocação
// GODEBUG=gctrace=1 GOGC=300 go run main.go

// sem uso do GC
// GODEBUG=gctrace=1 GOGC=-1 go run main.go
