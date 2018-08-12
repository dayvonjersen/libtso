import "math"

func round(val float64) int {
	var ret float64
	pow := math.Pow(10, 14)
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= 0.5 {
		ret = math.Ceil(digit)
	} else {
		ret = math.Floor(digit)
	}
	return int(ret / pow)
}
