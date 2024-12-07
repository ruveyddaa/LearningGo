package maps

import (
	"fmt"
)

func Demo1() {
	// bir dözlük yapısı denebilir
	//key -value mimarisi
	sozluk := make(map[string]string) // birnic stirng keyin tipi ikinci sözlük valuenin tipi
	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"

	fmt.Println(sozluk["book"])

	// sözlüğün i.inde arafoğımız key var mı diye bakma
	deger, varMi := sozluk["table"]
	fmt.Println(deger)
	fmt.Println("listede olma durumu: ", varMi)

	sozluk2 := map[string]string{"galass": "bardak", "microphone": "mikrofon"}

	fmt.Println(sozluk2)
}
