package main

import "fmt"

func main() {
	/*
		pipeline merupakan serangkaian function yang memiliki tipe data yang sama
		antara inputan/parameter dengan keluaran/hasil

		dibawah ini merupakan contoh function nya,
		dapat dilihat bahwa functin multiply dan add sama sama menerima dan menghasilkan list of intger

		kelemahan pada contoh function dibawah yaitu, pada penggunaan memory nya,
		dimana setiap kita memanggil fungsi add ataupun multiply, kita akan membuat list kosong
		dengan daya tampung sebanyak list inputan

		satu hal lagi, pada contoh dibawah fungsi" dibawah memproses data secara batch processing
		yaitu, memproses beberapa sekaligus, biasanya dalam bentuk list/array
	*/
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for counter, v := range values {
			multipliedValues[counter] = v * multiplier
		}
		return multipliedValues
	}

	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for counter, v := range values {
			addedValues[counter] = v + additive
		}
		return addedValues
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range multiply(add(multiply(ints, 2), 1), 2) {
		fmt.Println(v)
	}
}
