package struncts

import "fmt"

type constumer struct {
	firstName string
	lastName  string
	age       int
}

func (c constumer) save() {
	fmt.Println("eklendi")
}

func Demo2() {
	c := constumer{firstName: "engin", lastName: "polat", age: 23}
	c.save()

}
