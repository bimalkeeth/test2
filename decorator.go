package main

import (
	"fmt"
	"math"
	"time"
)

type Function func(float642 float64) float64

func ProfileDecorator(fn Function) Function {
	return func(params float64) float64 {
		start := time.Now()
		result := fn(params)
		elapsed := time.Now().Sub(start)
		fmt.Println("Function completed with time: ", elapsed)
		return result
	}
}
func SquareRoot(val float64) float64 {
	return math.Sqrt(val)
}

func main() {

	decSquaroot := ProfileDecorator(SquareRoot)
	fmt.Println(decSquaroot(16))

}
