package main

import "golesson/restful"

// go da PaccalCase yazıluır yani ilk harf byük
func main() {

	// variables.Demo1()    // variables paceksinden demo 1 içağırdık

	// conditionals.Demo2() // conditional packagesinden demo 2 yi çağırdık

	// loops.Demo1()

	// 				    arrays.Demo1()

	// 					maps.Demo1()

	// 					for_range.Demo1()

	// 					for_range.Demo2()

	// 					for_range.Demo3()

	// 					sayi := 20
	// 			        pointers.Demo1(sayi) // buraya sayının kendisi gönderilmnez sayının gösterdiği 20 değeri gönderilir bu yüzden demo fonksiyonndaki ssayı değerindeki değikliklik burdaki sayı değerini etkilemez
	// 			        fmt.Println("maindeki sayi", sayi)

	// sayi2 := 30
	// fmt.Println("maindeki sayi", sayi2)

	// sayilar := []int{1, 2, 3}
	// pointers.Demo3(sayilar)
	// fmt.Println("maindaki sayı", sayilar[0])  // araylerde bu değişiklik her yerde değişir çünkü ereyer zaten bir şeyin aderes değerine referans verir

	// struncts.Demo1()
	// struncts.Demo2()

	// goroutimes.Ciftsayilar() // burda görüldüğü gibi snytax sıarsına göre fonksiyonlar çalıştı
	// goroutimes.Teksayilar()

	// go goroutimes.Ciftsayilar()  // go yu koyduktan sonra hangisinin ilk çalışcağınu bellke karar verri hangisi daha hızlı ise o çalışır
	// go goroutimes.Teksayilar()
	// time.Sleep(5 * time.Second) // burda main fonkisyonun nun 5 sn sonra başlamsının söyeldik
	// fmt.Println("main bitti")

	// interfacess.Demo1()
	// interfacess.Demo2()

	// defer_stantment.B()
	// defer_stantment.Test()
	// defer_stantment.Demo3()

	// error_handiling.Demo1()

	// string_functions.Demo1()

	restful.Demo1()

}
