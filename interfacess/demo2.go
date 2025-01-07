package interfacess

import "fmt"

type Mortage struct {
	creditPaymentTotal float64
	adreess            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculater interface {
	Calculate() float64
}

func (m Mortage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 12

}
func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 12
}

func CalculateMontlyPayment(credits []CreditCalculater) float64 { // aray tipiş olarak interface atadık

	paymenttotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymenttotal = paymenttotal + credits[i].Calculate()

	}
	return paymenttotal

}

func Demo2() {
	credit1 := Mortage{rate: 10, creditPaymentTotal: 100000, adreess: "ankara"}
	credit2 := Mortage{rate: 12, creditPaymentTotal: 50000, adreess: "istanbul"}
	credit3 := Car{rate: 15, creditPaymentTotal: 60000, carInfo: "polo"}

	credits := []CreditCalculater{credit1, credit2, credit3}
	total := CalculateMontlyPayment(credits)

	fmt.Println("toplam ödeme", total)

}
