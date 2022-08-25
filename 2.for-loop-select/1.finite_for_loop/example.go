package finiteforloop

import "fmt"

func sendStringToChannel(ch chan string, message string) {
	ch <- message
}

func Example() {
	// finite loop means, the select statement will exucted as many as iteration
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendStringToChannel(ch1, "from channel one")
	go sendStringToChannel(ch2, "from channel two")

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
