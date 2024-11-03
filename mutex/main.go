package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(message string, mut *sync.Mutex) {
	defer wg.Done()
	defer mut.Unlock()
	mut.Lock()
	msg = message
}

func main() {

	var mutex sync.Mutex
	msg = "hello world"

	wg.Add(2)
	go updateMessage("Hello universe!", &mutex)
	go updateMessage("Hello cosmos!", &mutex)
	wg.Wait()

	fmt.Println(msg)

}
