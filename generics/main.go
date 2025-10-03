package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumNumber[k comparable, v Number](m map[k]v) v {
	var s v
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

	fmt.Printf("Non-generic sums : %v and %v\n", SumNumber(ints), SumNumber(floats))
}
