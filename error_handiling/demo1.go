package error_handiling

import "fmt"

func Demo1() {
	// go da error yakalamk için error verebilceğimizden şüphelendiğimiz değikene iki tane şey vericez bunlardan birincisi önmeli olmadığı için _ olanbilit ama ikincisi err olucak ve araya virgül koyucaz ve değiken tanımlıyomuşuz gibi := yapaıp değikenmizi koyucaz

	var number int
	_, err := fmt.Scan(&number) // fmt.ccam ile kullınıcdan number variablesini aldık kullanıcın int girimemesi durumıunda error yakalıyıcak

	if err != nil { // nil null gibi tyani error boş değilse
		fmt.Println("We got an error")
	} else {
		fmt.Println(number)
	}

}
