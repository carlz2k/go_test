package main

import (
	log "github.com/sirupsen/logrus"
	"go_test/com/carlz/helper"
	"go_test/com/carlz/lib"
)

func main() {
	a := []float64{1, 2, 3}
	b := lib.Sum(a)

	helper.Print(b)
	helper.Print(lib.Average(a))
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
