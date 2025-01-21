package string_functions

import (
	"fmt"
	s "strings" // burda alyans dediğimiz bir ifade kullandık yani strings i her kullanmak istediğimde s kullansak yeterli
)

func Demo1() {
	isim := "Engin"
	fmt.Println(isim)
	fmt.Println(s.Count(isim, "g")) //   strings adında go da tanımlanmış şey var bunun methodlarını çağırıp string değer için farklı işelemler yapabiliriz mesela count meyhodu idim için kaş kez g harfi olduğunu kontrıl eder

	// bunun gibi daha bir çok string methodu vrdır ihitycın olduğunda bakabnilisin
}
