package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startTime := time.Now()
	ch := make(chan string)

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://graph.microsoft.com",
		"https://random-api.romi.com",
	}

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for _, api := range apis {
		fmt.Printf(<-ch)
		fmt.Printf("API executed: %v", api)
	}

	elapsed := time.Since(startTime).Seconds()

	fmt.Printf("\nTotal time taken: %.2fs\n", elapsed)
}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("\nError gettin API: %v\n", api)
		return
	}
	ch <- fmt.Sprintf("\nAPI %s is reachable\n", api)
}
