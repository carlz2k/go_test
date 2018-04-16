package lib_test

import (
	"../lib"
	"testing"
	. "github.com/onsi/gomega"
)

func TestSwap(t *testing.T) {
	var testSuite = NewGomegaWithT(t)
	a := []float64{1, 2, 3}
	testSuite.Expect(lib.Sum(a)).To(Equal(6.0))
	b, c, d := lib.Swap(1, "2", "3")
	testSuite.Expect(b).To(Equal("3"))
	testSuite.Expect(c).To(Equal("2"))
	testSuite.Expect(d).To(Equal(1))
}
