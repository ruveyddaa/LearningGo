package arrays

import "fmt"

func Demo1() {
	sayilar := [5]int{20, 30, 40, 80, 60} // bu dizizdeki elemnalarunver, tipini belirti
	fmt.Println(sayilar)

	//length = uzunluk  yani len(sayilar) sayilar arrayinin uzunluğu kadar demek
	var enbüyük int = 0
	for i := 0; i < len(sayilar); i++ {
		if enbüyük < sayilar[i] {
			enbüyük = sayilar[i]
		}
	}

	fmt.Println(enbüyük)

}
