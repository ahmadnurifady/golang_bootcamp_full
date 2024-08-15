package materi

import "fmt"

type conversation interface {
	sayHello()
	sayGoodBye()
}

type Person struct {
	name string
	age  int
	born string
}

// func (p Person) sayHello() {
// 	fmt.Printf("Hello my name is %s,age %d, i boron from %s", p.name, p.age, p.born)
// }

// func (p Person) sayGoodBye() {
// 	fmt.Println("Goodbye Beach")
// }

func InterMuka() {

	var orang conversation

	person1 := Person{
		name: "Kocak",
		age:  69,
		born: "Alabama",
	}

	orang.sayHello()
	orang.sayGoodBye()
	fmt.Println(person1)

	// person1 := Person{
	// 	name: "Kocak",
	// 	age:  69,
	// 	born: "Alabama",
	// }

	// person1.sayHello()
	// person1.sayGoodBye()

}
