package calculator

import (
	"math"
)

type ComputeData struct {
	IncomeAmount        float64
	FractionSize        uint32
	BasePercentage      float64
	PercentageIncrement float64
}

// Compute returns the amount to pay as taxes based on an input amount
func Compute(data ComputeData) float64 {
	totalResult := 0.0
	currentPercentage := data.BasePercentage
	fractionsAmount := int(math.Floor(data.IncomeAmount / float64(data.FractionSize)))

	for i := 1; i <= fractionsAmount; i++ {
		isLast := i == fractionsAmount

		var fractionQuantity float64
		if isLast {
			// if it is the last section add only the remainder
			fractionQuantity = math.Mod(data.IncomeAmount, float64(data.FractionSize))
		} else {
			fractionQuantity = float64(data.FractionSize)
		}

		totalResult += fractionQuantity * currentPercentage
		if !isLast {
			currentPercentage += data.PercentageIncrement
		}
	}
	return totalResult
}
