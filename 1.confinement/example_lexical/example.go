package examplelexical

import "fmt"

func Example() {
	/*
		as you can see, there is shared variable between function in this example
	*/
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
	}

	results := chanOwner()
	consumer(results)
}
