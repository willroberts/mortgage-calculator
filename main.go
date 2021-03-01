package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	defaultHousePrice       = 300000 // Median U.S. home price Feb 2021.
	defaultDownPercent      = 0.2    // Standard 20% to avoid PMI.
	defaultTermYears        = 30     // Standard 30-year mortgage.
	defaultInterestRate     = 0.0275 // Feb 2021 interest rate.
	defaultPropertyTaxRate  = 0.0079 // Sample county property tax rate.
	defaultInsurancePayment = 1200   // Average annual homeowner's insurance costs.
)

var (
	housePrice       int
	downPayment      int
	downPercent      float64
	termYears        int
	interestRate     float64
	propertyTaxRate  float64
	insurancePayment int

	// todo: pmi, fha
)

func init() {
	flag.IntVar(&housePrice, "house-price", defaultHousePrice, "house price")
	flag.IntVar(&downPayment, "down-payment-flat", 0, "down payment (flat dollar amount)")
	flag.Float64Var(&downPercent, "down-payment-pct", defaultDownPercent, "down payment (percent of house price), e.g. 0.2 for 20%")
	flag.IntVar(&termYears, "term-years", defaultTermYears, "mortgage term in years, e.g. 30 or 15")
	flag.Float64Var(&interestRate, "interest-rate", defaultInterestRate, "interest rate, e.g. 0.0275 for 2.75%")
	flag.Float64Var(&propertyTaxRate, "property-tax", defaultPropertyTaxRate, "property tax rate, e.g. 0.0079 for 0.79%")
	flag.IntVar(&insurancePayment, "homeowners-insurance", defaultInsurancePayment, "annual homeowner's insurance payment")
	flag.Parse()
}

func main() {
	// Read back input.
	fmt.Println("House Price:", housePrice)
	fmt.Println("Mortgage Term (Years):", termYears)
	fmt.Printf("Interest Rate: %.2f%%\n", interestRate*100)
	fmt.Printf("Property Tax Rate: %.2f%%\n", propertyTaxRate*100)
	fmt.Printf("Annual Homeowner's Insurance Cost: $%d\n", insurancePayment)
	fmt.Println()

	// Calculate monthly payments.
	totalDown := downPayment + int((float64(housePrice) * downPercent))
	mortgageAmount := housePrice - totalDown
	principalAndInterest := computePrincipalAndInterest(mortgageAmount)
	insurance := float64(insurancePayment) / 12
	taxes := float64(housePrice) * propertyTaxRate / 12
	total := principalAndInterest + insurance + taxes

	// Show results.
	fmt.Printf("Down Payment: $%d\n", totalDown)
	fmt.Printf("Mortage Amount: $%d\n", mortgageAmount)
	fmt.Println()

	fmt.Println("Monthly Payment:")
	fmt.Printf("- Principal & Interest: $%.2f\n", principalAndInterest)
	fmt.Printf("- Homeowner's Insurance: $%.2f\n", insurance)
	fmt.Printf("- Property Taxes: $%.2f\n", taxes)
	fmt.Printf("- Total: $%.2f\n", total)
}

// Computes average monthly principal and interest payments over the lifetime of the loan.
func computePrincipalAndInterest(mortgageAmount int) float64 {
	loan := float64(mortgageAmount)     // Total mortgage amount.
	interest := interestRate / 12       // Monthly interest as percent.
	payments := float64(termYears * 12) // Total number of payments.

	return loan * (interest * math.Pow(1+interest, payments)) / (math.Pow(1+interest, payments) - 1)
}
