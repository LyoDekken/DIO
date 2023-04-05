package main

import (
	"fmt"
	"sync"
	"time"
)

func ping(c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		c <- "ping"
		time.Sleep(time.Second)
	}
}

func pong(c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		c <- "pong"
		time.Sleep(time.Second)
	}
}

func imprimir(c <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	c := make(chan string)

	go ping(c, &wg)
	go pong(c, &wg)
	go imprimir(c, &wg)

	wg.Wait()
}
