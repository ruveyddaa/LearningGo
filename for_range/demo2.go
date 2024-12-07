package for_range

import "fmt"

// 1-10 arasındaki sayılardan tek olanları toplayan program
func Demo2() {

	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sayilar)

	total := 0

	for _, sayi := range sayilar {
		if sayi%2 != 0 {
			total = total + sayi

		}

	}

	fmt.Println("toplam :", total)

}
