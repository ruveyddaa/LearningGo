package variables

import "fmt"

func Demo1() {
	var metin string = "merhba dünya!"
	fmt.Println(metin) // print yazdırı println satır atlatıp yazdırı P büyk olamlı

	var kdv int = 10 // değiken tanımlamak için var zorunlu bir şey
	fmt.Println(kdv)
	fmt.Println(100 + 100*10/100)

	var kdv2 float32 = 0.1
	fmt.Println(kdv2)

	kdv3 := 25 // go ya özgü bir veri tip ayarklama yönetemi  değişken tipini kendisi ayaralar
	fmt.Println(kdv3)

	fmt.Printf("veri tipi: %T \n", kdv3) //%T yazdığımızada bu bise değişkenimizn tipini verir

	var durum bool = true

	var metin1 string = "engin"
	var metin2 string = "metin"

	durum = metin1 == metin2

	fmt.Println(durum)
	fmt.Println(!durum)
}
