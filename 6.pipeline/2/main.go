package main

import "fmt"

func main() {
	/*
		contoh pipeline stream processing

		kelemahan dari contoh dibawah adalah, kita akan memanggil fungsi sebanyak perulangan.
		kelemahan yang lainnya adalah, dengan cara dibawah kita melimitasi scaling dan reusabilitas
		dari fungsi" dibawah
	*/

	multiply := func(value int, multiplier int) int {
		return value * multiplier
	}

	add := func(value int, additive int) int {
		return value + additive
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	}
}
