package defer_stantment

import "fmt"

func Demo2(sayi int32) string {
	defer Deferfunc()
	if sayi%2 == 0 {
		return " çift sayıdır" // burda return ifdadesi var yani program biterken çaışıcak demek deferden bile sonra çalışıcak
	}
	if sayi%2 != 0 {
		return "tek saıdır"
	}
	return "belli değil"

}

func Test() {
	sonuc := Demo2(9)
	fmt.Println(sonuc)
}

func Deferfunc() { // eğer bu fonkisyonun başına defer koyamzsak çalışmaz çünkü yukardaki fonksiyonlarda program retrunü gördü ve durdu
	fmt.Println("bitti")
}
