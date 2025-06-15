package main
import "fmt"
func input() (float64, float64, float64, float64, float64, float64, float64, float64) {
	var g, pr, pe, is, r, h, b, d float64
	fmt.Print("Gross income(monthly): ")
	fmt.Scan(&g)
	fmt.Print("Second pillar(percentage): ")
	fmt.Scan(&pr)
	fmt.Print("Third pillar: ")
	fmt.Scan(&pe)
	fmt.Print("Investments/Savings: ")
	fmt.Scan(&is)
	fmt.Print("Rent: ")
	fmt.Scan(&r)
	fmt.Print("Household expenses: ")
	fmt.Scan(&h)
	fmt.Print("Bills: ")
	fmt.Scan(&b)
	fmt.Print("Debts: ")
	fmt.Scan(&d)
	if g < 0 || pr < 0 || pe < 0 || is < 0 || r < 0 || h < 0 || b < 0 || d < 0{
		panic("Can't use negative numbers.")
	} else {
		return g, pr, pe, is, r, h, b, d
	}
}
func excal(g float64) float64 {
	if g <= 1200 {
		return 654
	} else if g > 1200 && g < 2100 {
		return 654-(0.72667*(g-1200))
	} else {
		return 0
	}
}
func netcal(e, g float64) float64 {
	return e+((g-e)*0.78)
}
func ptotal(pe, g, pr float64) float64 {
	return pe+(g*(pr/100))
}
func left(n, pt, i, r, h, b, d float64) float64 {
	return n-pt-i-r-h-b-d
}
func main() {
	gross, pension_rate, pension_extra, invsav, rent, household, bills, debt := input()
	exemption := excal(gross)
	net := netcal(exemption, gross)
	pension_total := ptotal(pension_extra, gross, pension_rate)
	remaining := left(net, pension_total, invsav, rent, household, bills, debt)
	fmt.Printf("------\nYour net income is %.2f€.\nAfter paying:\n%.2f€ into your pensions,\n%.2f€ in investments/savings\n%.2f€ in rent,\n%.2f€ for food,\n%.2f€ in bills,\n%.2f€ into debts.\nYou have %.2f€ left.\n", net, pension_total, invsav, rent, household, bills, debt, remaining)
}
