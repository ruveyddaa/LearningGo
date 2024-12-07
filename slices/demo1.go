package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3) // bu bir slices tanımla yöntemidir virgülden önceki yere tipini sonraki yerede eleman sayısını yazarızn

	fmt.Println(isimler)
	isimler[0] = "engin"
	isimler[1] = "engin"
	isimler[2] = "engin"
	// YENİ BİR SLİCE EKLEME FONKSİYONU

	isimler = append(isimler, "büşra ") // append go da tnımlı bir fonksiyondur ve virgülden sonra yazdığımız şeyi arrayin sonuna yeni bir rlman olark ekler

	fmt.Println(isimler)
	fmt.Println(len(isimler))

}
