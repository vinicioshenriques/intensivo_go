package main

import (
	"fmt"
	"time"
)

// Processo -> Alocar um bloco de memória
// Thread -> Acessar o bloco de memória
// T1 e T2 -> Acessam o mesmo bloco de memória
// Race Condition -> Condição de corrida

// GO ROUTINE -> CANAL -> GO ROUTINE 2
// INPUT -> OUTPUT

func WorkerPost(workerID int, canal1 chan int, canal2 chan int) {
	for x := range canal1 {
		fmt.Println("WorkerPost ID:", workerID, "Data:", x)
		canal2 <- x * 2
		time.Sleep(time.Second)
	}
}

func WorkerReceiver(workerID int, canal1 chan int) {
	for x := range canal1 {
		fmt.Println("WorkerReceiver ID:", workerID, "Data:", x)
	}
}

func main() { // T1
	canal := make(chan int)
	canal2 := make(chan int)

	qtdWorkers := 100

	for i := 0; i < qtdWorkers; i++ {
		go WorkerPost(i, canal, canal2)
		go WorkerReceiver(i, canal2)
	}

	for i := 0; i < 100; i++ {
		canal <- i
	}
}
