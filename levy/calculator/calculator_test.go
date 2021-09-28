package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	tt := []struct {
		name     string
		given    ComputeData
		expected float64
	}{
		{
			name: "Iterates only once when the IncomeAmount is the same as the FractionSize",
			given: ComputeData{
				IncomeAmount:        100,
				FractionSize:        100,
				BasePercentage:      0.1,
				PercentageIncrement: 0.1,
			},
			expected: 10.0,
		},
	}
	for _, testData := range tt {
		t.Run(testData.name, func(t *testing.T) {
			result := Compute(testData.given)

			assert.Equal(t, testData.expected, result)
		})
	}
}
