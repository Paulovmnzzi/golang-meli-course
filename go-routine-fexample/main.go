package main

import (
	"fmt"
	"sync"
)

func printSomething(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(msg)
}

func main() {

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
		"zeta",
		"eta",
	}

	wg.Add(len(words))

	for indice, variable := range words {
		go printSomething(fmt.Sprintf("%d:%s", indice, variable), &wg)
	}

	wg.Wait()

}
