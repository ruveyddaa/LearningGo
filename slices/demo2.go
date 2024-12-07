package slices

import "fmt"

func Demo2() {

	sehirler := []string{"ankara", "istanbul", "izmir"}
	fmt.Println(sehirler)
	sehirlerkopya := make([]string, len(sehirler))
	fmt.Println(sehirlerkopya)
	copy(sehirlerkopya, sehirler) // burda sehrilerkopyaya sehielri kopyaladık
	fmt.Println(sehirlerkopya)
	sehirler = append(sehirler, "adana")
	fmt.Println(sehirler)
	fmt.Println(sehirlerkopya) //append ile sehilere eklediğim eleman eklemeden önce kopyalanmış sehirler kpyay slicesinde ekelenmez

	fmt.Println(sehirler[1:3]) // birinci indisten 3 e kadar yaz ama 3 dahil olmasın demek
	fmt.Println(sehirler[1:])  // birden sonrtaki bütün insisleri al
	fmt.Println(sehirler[:2])  // 2 ye kafdar bütün indisleri al
}
