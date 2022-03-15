package math

func Ramp(s float64, a, b float64) float64 {
	return (s - a) / (b - a)
}

func RampSat(s float64, a, b float64) float64 {
	return Saturate((s - a) / (b - a))
}

func Saturate(x float64) float64 {
	return Clampf(x, 0, 1)
}

func SmoothStep(s float64, a, b float64) float64 {
	x := RampSat(s, a, b)
	return x * x * (3 - 2*x)
}

func Clamp(x, min, max int) int {
	switch {
	case x < min:
		return min
	case x > max:
		return max
	default:
		return x
	}
}

func Clampf(x, min, max float64) float64 {
	switch {
	case x < min:
		return min
	case x > max:
		return max
	default:
		return x
	}
}
