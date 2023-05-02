package variance

import (
	"math"
	"testing"
)

func TestWelfordBasicExample(t *testing.T) {
	mean, variance, sampleVariance := Welford([]float64{15, 65, 55, 35, 45, 25})
	expectedMean, expectedVariance, expectedSampleVariance := 40.0, 291.66666, 350.0
	if math.Abs(mean-expectedMean) > 0.001 {
		t.Fatalf("mean incorrect: expected %f, got %f", expectedMean, mean)
	}
	if math.Abs(variance-expectedVariance) > 0.001 {
		t.Fatalf("variance incorrect: expected %f, got %f", expectedVariance, variance)
	}
	if math.Abs(sampleVariance-expectedSampleVariance) > 0.001 {
		t.Fatalf("sample variance incorrect: expected %f, got %f", expectedSampleVariance, sampleVariance)
	}
}
