package struncts

import "fmt"

func Demo1() {

	fmt.Println(product{"biligisayar", 500, "XYZ"})
	fmt.Println(product{"biligisayar", 500, "XYZ"})
	fmt.Println(product{name: "bilgisayar"})

}

type product struct {
	name      string
	unitprice float64
	brand     string
}
