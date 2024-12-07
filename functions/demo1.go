package functions

import "fmt"

func Topla(sayi1 int, sayi2 int) int { // sondaki return tipidir

	var toplam = sayi1 + sayi2
	return toplam

}

func selamver(kullanıcıadı string) {
	fmt.Println("merhaba", kullanıcıadı)
}
