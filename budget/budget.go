package main
import "fmt"
func main() {
	var gross, tax, pension_rate, pension_extra, investments, rent, household, bills, debt, net, pension_total, left float32
	fmt.Print("Gross salary: ")
	fmt.Scan(&gross)
	fmt.Print("Tax rate(percentage): ")
	fmt.Scan(&tax)
	fmt.Print("Second pillar(percentage): ")
	fmt.Scan(&pension_rate)
	fmt.Print("Third pillar: ")
	fmt.Scan(&pension_extra)
	fmt.Print("Investments: ")
	fmt.Scan(&investments)
	fmt.Print("Rent: ")
	fmt.Scan(&rent)
	fmt.Print("Household expenses: ")
	fmt.Scan(&household)
	fmt.Print("Bills: ")
	fmt.Scan(&bills)
	fmt.Print("Debts: ")
	fmt.Scan(&debt)
	net = gross * (1 - (tax / 100))
	pension_total = pension_extra + (gross * (pension_rate / 100))
	left = net - pension_total - investments - rent - food - bills - debt
	fmt.Printf("Your net income is %.2f€.\nAfter paying:\n%.2f€ into your pensions,\n%.2f€ in investments\n%.2f€ in rent,\n%.2f€ for food,\n%.2f€ in bills,\n%.2f€ into debts.\nYou have %.2f€ left.\n", net, pension_total, investments, rent, household, bills, debt, left)
}
