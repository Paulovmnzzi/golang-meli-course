package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var bankBalance int
	var mutex sync.Mutex

	//Print out starting values
	fmt.Printf("Initial account balance: %d\n", bankBalance)
	fmt.Println()

	//Define weekly revenue
	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Freelancer", Amount: 100},
		{Source: "Investments", Amount: 50},
		{Source: "Romi's gifts", Amount: 10},
	}

	// wg.Add(len(incomes))

	//loop through 52 weeks and print out how much is made;
	for i, income := range incomes {
		wg.Add(1)

		go func(i int, income Income, mutex *sync.Mutex) {
			defer wg.Done()

			mutex.Lock()

			for week := 1; week <= 52; week++ {
				actualBalance := bankBalance
				actualBalance += income.Amount
				bankBalance = actualBalance

				// fmt.Printf("On week %d, you earned %d from %s \n", week, income.Amount, income.Source)
			}
			mutex.Unlock()
		}(i, income, &mutex)
	}

	wg.Wait()

	fmt.Printf("Final account balance: $%d", bankBalance)

}
