package main
import "fmt"
func excal(g float64) float64 {
	if g <= 1200 || 654-(0.72667*(g-1200)) >= g*0.22 {
		return g*0.22
	} else if g > 1200 && g < 2100 {
		return 654-(0.72667*(g-1200))
	} else {
		return 0
	}
}
func netcal(e, g float64) float64 {
	return e+(g*0.78)
}
func ptotal(pe, g, pr float64) float64 {
	return pe+(g*(pr/100))
}
func left(n, pt, i, r, h, b, d float64) float64 {
	return n-pt-i-r-h-b-d
}
func main() {
	var gross, exemption, pension_rate, pension_extra, invsav, rent, household, bills, debt, net, pension_total, remaining float64
	fmt.Print("Gross income(monthly): ")
	fmt.Scan(&gross)
	fmt.Print("Second pillar(percentage): ")
	fmt.Scan(&pension_rate)
	fmt.Print("Third pillar: ")
	fmt.Scan(&pension_extra)
	fmt.Print("Investments/Savings: ")
	fmt.Scan(&invsav)
	fmt.Print("Rent: ")
	fmt.Scan(&rent)
	fmt.Print("Household expenses: ")
	fmt.Scan(&household)
	fmt.Print("Bills: ")
	fmt.Scan(&bills)
	fmt.Print("Debts: ")
	fmt.Scan(&debt)
	exemption = excal(gross)
	net = netcal(exemption, gross)
	pension_total = ptotal(pension_extra, gross, pension_rate)
	remaining = left(net, pension_total, invsav, rent, household, bills, debt)
	fmt.Printf("------\nYour net income is %.2f€.\nAfter paying:\n%.2f€ into your pensions,\n%.2f€ in investments/savings\n%.2f€ in rent,\n%.2f€ for food,\n%.2f€ in bills,\n%.2f€ into debts.\nYou have %.2f€ left.\n", net, pension_total, invsav, rent, household, bills, debt, remaining)
}
