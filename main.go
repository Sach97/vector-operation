package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sub([]float64{0,1},[]float64{0,1}))
	// => [0 1]
}

// Sub - Vector arithmetic subtraction
func Sub(inReal0 []float64, inReal1 []float64) []float64 {
	outReal := make([]float64, len(inReal0))
	for i := 0; i < len(inReal0); i++ {
		outReal[i] = inReal0[i] - inReal1[i]
	}
	return outReal
}