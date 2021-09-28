package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/miguel-pm/go-levy/calculator"
	"github.com/miguel-pm/go-levy/values"
)

func main() {
	fmt.Println("Launching Process.")

	if len(os.Args[1:]) == 0 {
		fmt.Println("ERROR: no cli arguments provided, Income Amount required")
		os.Exit(1)
	}

	incomeAmount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("ERROR: provided value for Income Amount is invalid:")
		fmt.Println(err)
		os.Exit(1)
	}

	calcData := calculator.ComputeData{
		IncomeAmount:        incomeAmount,
		FractionSize:        values.FRACTION_SIZE,
		BasePercentage:      values.BASE_PERCENTAGE,
		PercentageIncrement: values.PERCENTAGE_INCREMENT,
	}
	fmt.Println("Income Amount -------->", calcData.IncomeAmount)
	fmt.Println("Fraction Size -------->", calcData.FractionSize)
	fmt.Println("Base Percentage ------>", calcData.BasePercentage)
	fmt.Println("Percentage Increment ->", calcData.PercentageIncrement)

	result := calculator.Compute(calcData)

	fmt.Println("Finished, result: ", result)
}
