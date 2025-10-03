package main

import "fmt"

func SumInt(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloat(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func main() {

	ints := map[string]int64{
		"First":  145,
		"Second": 145,
	}

	floats := map[string]float64{
		"First":  145.155,
		"Second": 145.155,
	}

	fmt.Printf("Non-generic sums : %v and %v\n", SumInt(ints), SumFloat(floats))
}
