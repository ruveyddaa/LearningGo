package defer_stantment

import "fmt"

type product struct {
	productname string
	unitprice   int
}

func (p product) save() {
	fmt.Println("ürün kaydedildi: ", p.productname)
}

func Demo3() {
	p := product{productname: "laptop", unitprice: 5000}
	defer p.save()
	// p = product{productname: "Mouse", unitprice: 5000}  // burda p nin name değerini değiştirdik ama outputta bir farklılık olmaz çünkü program deferü gördüğü yere kadar değişiklikleri alır eğer deferün üstüde değiştirseyidk değişikliği okurud
	fmt.Println("işlem başarılı")
}
