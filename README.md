# mortgage-calculator

Calculates monthly mortgage payments, including property taxes and homeowner's
insurance, for configurable values of house prices, down payments, mortgage
durations, interest rates, property tax rates, and homeowner's insurance costs.

## Usage

```
$ go run main.go -help
  -down-payment-flat int
        down payment (flat dollar amount)
  -down-payment-pct float
        down payment (percent of house price), e.g. 0.2 for 20% (default 0.2)
  -homeowners-insurance int
        annual homeowner's insurance payment (default 1200)
  -house-price int
        house price (default 300000)
  -interest-rate float
        interest rate, e.g. 0.0275 for 2.75% (default 0.0275)
  -mortgage-insurance-rate float
        pmi insurance rate, e.g. 0.01 for 1% (default 0.01)
  -property-tax float
        property tax rate, e.g. 0.0079 for 0.79% (default 0.0079)
  -term-years int
        mortgage term in years, e.g. 30 or 15 (default 30)

$ go run main.go
House Price: $300000
Down Payment: $60000 (20.00%)
Mortage Amount: $240000
Mortgage Term (Years): 30
Interest Rate: 2.75%
Property Tax Rate: 0.79%
Annual Homeowner's Insurance Cost: $1200

Monthly Payment: $1277.28
- Principal & Interest: $979.78
- Homeowner's Insurance: $100.00
- Property Taxes: $197.50
```