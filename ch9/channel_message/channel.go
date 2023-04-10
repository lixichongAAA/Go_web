package main

import (
	"fmt"
	"sync"
)

func thrower(c chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("Throw >> ", i)
	}
	wg.Done()
}

func catcher(c chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("Catch << ", num)
	}
	wg.Done()
}

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go thrower(c, &wg)
	go catcher(c, &wg)
	wg.Wait()
}
