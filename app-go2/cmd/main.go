package main

import (
	"fmt"
	"math"
)

func calculaDelta(a float64, b float64, c float64) float64 {
	var delta float64
	delta = (math.Pow(b, 2) - (4 * a * c))
	return delta

}

func bhaskara(a float64, b float64, c float64) (bool, float64, float64) {
	var delta float64 = calculaDelta(a, b, c)
	fmt.Println(delta)
	if delta < 0 {
		return false, math.NaN(), math.NaN()
	} else if delta == 0 {
		x := -b / (2 * a * c)
		return true, x, x
	} else if delta > 0 {
		x1 := -b + math.Sqrt(delta)
		x2 := -b - math.Sqrt(delta)
		return true, x1, x2
	}
	return false, math.NaN(), math.NaN()
}

func main() {
	var root, x1, x2 = bhaskara(7, 30, 4)
	fmt.Println(root, x1, x2)
}
