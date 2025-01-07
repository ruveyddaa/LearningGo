package struncts

import "fmt"

type constumer struct {
	firstName string
	lastName  string
	age       int
}

func (c constumer) save() { // struct için bir methıd oluşturduk   zorunlu değil bunsuzda bir structı main içinde çalığıarbilşrsin am abu stracta bir fonksiyon atar ve strctla oluşturdupun obj ile bu fonkişsyonu kolayca çağırabilirsin
	fmt.Println("eklendi")
}

func Demo2() {
	c := constumer{firstName: "engin", lastName: "polat", age: 23}
	c.save()

}
