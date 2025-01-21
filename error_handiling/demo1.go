package error_handiling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt") // os go da olabn bir pakettir    burda demo1.txt varsa f yoksa error dönüyo
	// eğer dosya bulunursa f e atandı ve err nil denilen bir değer aldı
	if err != nil {
		fmt.Println("dosya bulunamdı")
	} else {
		fmt.Println(f.Name()) // name de os paketinde olan bir method dur ve f i bir obj gibi düşün f ile name methodunu çağırdık
	}

	// şu anda demo1.txt diye bir dosya oluşturduk ama bunu dzigi main ile aynı yerde olmalı çünkü mainde kod çalışıyo

}
