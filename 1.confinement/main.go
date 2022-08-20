package main

import "fmt"

func main() {
	/*
		the goal of confinement is to make safe operation fo concurrency code.
		this goal can be achived by mutex and channels too. but, using confinement is implicit way to achive,
		so we no longer need to syncronized the concureny operations (mutex, channels for example).

		note: immutable data is also implicit way to achive safe concurrency operations.

		there is two type of confinement:
		- ad hoc, this one achive confinement by convention wheter is set by languages community, group u work within
		  or the codebase you work within. (hard to achive)

		- lexical
	*/

	// exampleadhoc.Example()
	// examplelexical.Example()

	// another example that proof, not using confinement is bad idea, the result will be inconsisstent
	// var counter int
	// incrementFunc := func() {
	// 	counter += 1
	// }

	// for i := 0; i < 100; i++ {
	// 	go incrementFunc()
	// }

	// fmt.Println(counter)

	// you may think, not all goroutines finished,
	// even when we try to use wait group to make sure all goroutines is done, the result still will be inconsistent
	// var counter int
	// var wg sync.WaitGroup

	// incrementFunc := func() {
	// 	defer wg.Done()
	// 	counter += 1
	// }

	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)
	// 	go incrementFunc()
	// }

	// wg.Wait()
	// fmt.Println(counter)

	// let's try using mutex to handle this and use wait group to make sure all goroutines is called.
	// but using this approach, is same as sequantial, it's because only 1 gourtine can access the acouter at one time,
	// and the rest will be wait until the current goroutine is done.
	// type SafeCounter struct {
	// 	counter int
	// 	mu      sync.Mutex
	// }

	// var safeCounter SafeCounter
	// var wg sync.WaitGroup

	// incrementFunc := func() {
	// 	defer wg.Done()
	// 	safeCounter.mu.Lock()
	// 	safeCounter.counter += 1
	// 	safeCounter.mu.Unlock()
	// }

	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)
	// 	go incrementFunc()
	// }

	// wg.Wait()
	// fmt.Println(safeCounter.counter)

	// confinement applied, as you can see, counter variable only used when read value from channel
	// the channel itself, is can only write by incrementFunc
	var counter int

	incrementFunc := func() <-chan int {
		result := make(chan int)
		go func() {
			defer close(result)
			for i := 0; i < 100; i++ {
				result <- 1
			}
		}()
		return result
	}

	result := incrementFunc()
	for val := range result {
		counter += val
	}
	fmt.Println(counter)
}
