package variance

import "math"

// Welford's online algorithm as described on Wikipedia:
// https://en.wikipedia.org/wiki/Algorithms_for_calculating_variance
func Welford(numbers []float64) (mean, variance, sampleVariance float64) {
	var a welfordAggregate
	for i := range numbers {
		a = welfordUpdate(a, numbers[i])
	}
	return welfordFinalize(a)
}

type welfordAggregate struct {
	count, mean, m2 float64
}

func welfordUpdate(a welfordAggregate, next float64) welfordAggregate {
	count := a.count + 1.0
	delta := next - a.mean
	mean := a.mean + (delta / count)
	delta2 := next - mean
	m2 := a.m2 + (delta * delta2)
	return welfordAggregate{count, mean, m2}
}

func welfordFinalize(a welfordAggregate) (mean, variance, sampleVariance float64) {
	if a.count < 2 {
		return a.mean, math.NaN(), math.NaN()
	}
	return a.mean, a.m2 / a.count, a.m2 / (a.count - 1)
}
