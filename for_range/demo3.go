package for_range

import "fmt"

func Demo3() {

	sozluk := map[string]string{"book ": " kitap", "table ": " masa"}

	for k, v := range sozluk { // for range maps yazdırıken ilk önce key için bir değiken sobra value için bşr değişken verilir

		fmt.Print(k)
		fmt.Println(v)
	}
}
