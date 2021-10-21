package calculator

import (
	"math"
)

type ComputeData struct {
	TotalTarget         float64
	IncomeFirst         float64
	IncomeSecond        float64
	BasePercentage      float64
	FractionSize        uint32
	PercentageIncrement float64
}

// This calculates the amount of contribution an income is due, this depends on the
// income amount but the increment gets bigger the bigger the income, it's not linear.
func getContributionPercentage(income float64, data ComputeData) float64 {
	currentPercentage := data.BasePercentage
	total := 0.0
	fractionsAmount := int(math.Floor(income / float64(data.FractionSize)))

	for i := 1; i <= fractionsAmount; i++ {
		isLast := i == fractionsAmount

		total += currentPercentage
		if !isLast {
			currentPercentage += data.PercentageIncrement
		}
	}
	return total
}

// With both calculated percentages it gets absolute ones based on the sum of both
func getAbsPercentages(first float64, second float64) (float64, float64) {
	total := first + second
	firstAbs := first * 100 / total
	secondAbs := second * 100 / total

	return firstAbs, secondAbs
}

// Compute returns the amount to pay as taxes based on input amounts and a total target
func Compute(data ComputeData) (float64, float64) {
	firstBasePercentage := getContributionPercentage(data.IncomeFirst, data)
	secondBasePercentage := getContributionPercentage(data.IncomeSecond, data)

	firstPercentage, secondPercentage := getAbsPercentages(firstBasePercentage, secondBasePercentage)

	firstAmount := firstPercentage * data.TotalTarget / 100
	secondAmount := secondPercentage * data.TotalTarget / 100

	return firstAmount, secondAmount
}
