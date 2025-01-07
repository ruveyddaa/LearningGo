package defer_stantment

import "fmt"

// fonkisyonun en sonunda çalışan ve çalışmasını garantie ettiğimiz şeydir
// mesela bir fonkiyonu defer ile tanımlarsak çalışmasını garantie deriz

func A() {
	fmt.Println("a çalıştı")
}
func C() {
	fmt.Println("C çalıştı")
}
func D() {
	fmt.Println("d çalıştı")
}

// defer en sondaki süslü parantezi görğdkten sonra en sondan defer ile yazılanları okur
func B() {
	defer C()
	fmt.Println("b çalıştı") // program en son defer ile tanımlanalarıo çalıştırırı yani defersüz olanlar normal bir şekilde çalışır
	defer D()
	defer A() // defer kodu a nın yeri farketmeksizin en son çaılışır   LAST İIN FIRST OUT ÇALIŞIR

}
