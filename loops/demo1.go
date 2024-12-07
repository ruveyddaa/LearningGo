package loops

import "fmt"

func Demo1() {
	var sayi1, sayi2 int = 0, 0

	fmt.Println("birinci sayiyi giriniz")
	fmt.Scanln(&sayi1)
	fmt.Println("ikinci sayiyi giriniz")
	fmt.Scanln(&sayi2)

	var totoal1 int = 0
	var totoal2 int = 0
	for i := 2; i <= sayi1; i++ {
		if sayi1%i == 0 {
			totoal1 = +i

		}

	}
	for a := 2; a <= sayi2; a++ {
		if sayi2%a == 0 {
			totoal2 = +a

		}

	}
	fmt.Println(totoal1)
	fmt.Println(totoal2)

	if totoal1 == sayi2 && totoal2 == sayi1 {
		fmt.Println("bu sayilar arakadaş sayidir")
	} else {
		fmt.Println("bu sayilar arkaşadaş değildir")
	}

}
