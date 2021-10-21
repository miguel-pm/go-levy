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

	executionArgs := os.Args[1:]
	if len(executionArgs) == 0 {
		fmt.Println("ERROR: no cli arguments provided")
		os.Exit(1)
	}
	if len(executionArgs) < 3 {
		fmt.Println("ERROR: not enough arguments provided. Required: TOTAL_TARGET, INCOME_1, INCOME_2")
		os.Exit(1)
	}

	totalTarget, err := strconv.ParseFloat(executionArgs[0], 64)
	if err != nil {
		fmt.Println("ERROR: provided value for TOTAL_TARGET is invalid:")
		fmt.Println(err)
		os.Exit(1)
	}
	incomeFirst, err := strconv.ParseFloat(executionArgs[1], 64)
	if err != nil {
		fmt.Println("ERROR: provided value for INCOME_1 is invalid:")
		fmt.Println(err)
		os.Exit(1)
	}
	incomeSecond, err := strconv.ParseFloat(executionArgs[2], 64)
	if err != nil {
		fmt.Println("ERROR: provided value for INCOME_2 is invalid:")
		fmt.Println(err)
		os.Exit(1)
	}

	calcData := calculator.ComputeData{
		TotalTarget:         totalTarget,
		IncomeFirst:         incomeFirst,
		IncomeSecond:        incomeSecond,
		BasePercentage:      values.BASE_PERCENTAGE,
		FractionSize:        values.FRACTION_SIZE,
		PercentageIncrement: values.PERCENTAGE_INCREMENT,
	}
	fmt.Println("Total Target --------->", calcData.TotalTarget)
	fmt.Println("Income First Amount -->", calcData.IncomeFirst)
	fmt.Println("Income Second Amount ->", calcData.IncomeSecond)
	fmt.Println("Base Percentage ------>", calcData.BasePercentage)
	fmt.Println("Fraction Size -------->", calcData.FractionSize)
	fmt.Println("Percentage Increment ->", calcData.PercentageIncrement)

	first, second := calculator.Compute(calcData)

	fmt.Printf("Finished, result. First: %f, Second: %f", first, second)
}
