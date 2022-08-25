package main

import (
	finiteForLoop "github.com/mrrizal/learn-go-concurrency-pattern/2.for-loop-select/1.finite_for_loop"
	inifiniteForLoop "github.com/mrrizal/learn-go-concurrency-pattern/2.for-loop-select/2.infinite_for_loop"
)

func main() {
	// please take a note, when we close the channel, the channel will return empty data depens on what channel type is
	finiteForLoop.Example()
	inifiniteForLoop.Example()
}
