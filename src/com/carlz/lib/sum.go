package lib

func Sum(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total
}

func Average(xs []float64) float64 {
	return Sum(xs) / float64(len(xs))
}
