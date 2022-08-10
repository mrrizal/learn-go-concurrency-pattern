package main

import (
	examplelexical "github.com/mrrizal/learn-go-concurrency-pattern/1.confinement/example_lexical"
)

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
	examplelexical.Example()
}
