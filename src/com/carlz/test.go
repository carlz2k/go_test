package main

import (
	"com/carlz/lib"
	"com/carlz/helper"
)

func main() {
	a := []float64{1, 2, 3}
	b := lib.Sum(a)
	helper.Print(b)
	helper.Print(lib.Average(a))

}
