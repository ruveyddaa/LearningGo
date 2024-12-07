package for_range

import "fmt"

func Demo1() {
	sehirler := []string{"ankara", "istanbul", "izmir"}
	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}
	for _, sehir := range sehirler { // range arrayin içinde kendisi gezear _ koyduğumuz yere i de koyabilridik bu sadece bize sıra belli ediyo for döngüsünde norma bi şekilde array gezmek için
		fmt.Println(sehir)
	}
}
