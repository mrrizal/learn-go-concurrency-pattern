package exampleadhoc

import "fmt"

func Example() {
	/*
		this example achive confinment by ad hoc, while the "data" variable is only used at loopData function.
		but it's very possible to someone touch the variable data at other places and break down the confinement.
	*/

	// create empty list of int with size 4
	data := make([]int, 4)

	// create channel
	handleData := make(chan int)

	// this function will recive write only channel (int) as parameter
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}
