package interfacess

import (
	"fmt"
	"math"
)

// temel amaç aynı methodu birden farklı kez birden farklı şekilde kullanambilmel
type shape interface { //isimi sahpe olan bir intaetface taanımladık
	area() float64 // boş bir method tanımladık
}

type rectengle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectengle) area() float64 { // burda rectnge structı için bir methıd tanımladık(mu method interfacenin methıdu) rectengle strutında bir r objesi olupturduk ve en sonada tipini yazdık float64 olarak
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape) { //paraemtre olarak interface verdik    burda i.inde area olan herşey bu fonksiyona gelebilir

	fmt.Println("şeklin alanı", s.area()) // burdakş araea methıodu s hangi struct ise ona gçre çalılşır
}

func Demo1() {

	r := rectengle{width: 10, height: 6}
	school(r) // burda school methodu r yi pssrmetre olarak almış ve e bir rectangale structının objesi bu yüzden area methody rectagle için çalışır
}
