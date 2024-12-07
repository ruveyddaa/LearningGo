package pointers

import "fmt"

func Demo2(sayi2 *int) {

	*sayi2 = *sayi2 + 1
	fmt.Println("demodaki sayi", *sayi2)

}
