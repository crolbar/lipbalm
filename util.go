package main

func clamp[T int | float64](v, low, high T) T {
	return min(max(v, low), high)
}
