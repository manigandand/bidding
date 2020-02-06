package main

import (
	"fmt"
	"sync"
)

func main() {
	input := []int{1, 2, 3, 4, 5}
	in := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go write(input, in, &wg)
	go read(in, &wg)
	wg.Wait()

	fmt.Println("Done..")
}

func write(input []int, in chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, i := range input {
		in <- i
	}
	close(in)
	return
}

func read(in chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for out := range in {
		fmt.Println(out)
	}
}
