package conditionals

import "fmt"

func Demo2() {

	var sayi1, sayi2, sayi3 int = 10, 20, 30

	var buyuksayi int = sayi3

	if sayi2 > buyuksayi {
		buyuksayi = sayi2
	}

	if sayi1 > buyuksayi {
		buyuksayi = sayi3
	}

	fmt.Printf("en büyük sayı : %v", buyuksayi) // %v value of buyuk sayı yani %v v,rgülden somnra konan şeyin valusu yazar

}
