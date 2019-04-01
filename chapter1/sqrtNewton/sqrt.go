package sqrtNewton

import l "log"

func Sqrt(number float64, iteration uint, log bool) (out float64) {
	if iteration == 0 {
		if log {
			l.Printf("iteration: %d, out: %e", 0, number)
		}
		return number
	}
	out = number / 2
	if iteration == 1 {
		if log {
			l.Printf("iteration: %d, out: %e", 1, out)
		}
		return
	}
	for i := uint(2); i <= iteration; i++ {
		out = (number/out + out) / number
		if log {
			l.Printf("iteration: %d, out: %.40f", i, out)
		}
	}
	return
}
