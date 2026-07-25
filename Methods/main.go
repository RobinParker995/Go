package main

import "fmt"

type Person struct {
	Name string
	Age int
}

//Basic Syntax with Value Recivers
//The p Person is called the Value Reciever, It only copies the value of struct.

func (p Person) greet() {
	fmt.Println("Hello,", p.Name) 
}


//Mathod with a parameter
func (p Person) intro(city string) {
	fmt.Printf("I am %s, and I am from %s \n", p.Name, city)
}

//Method with return value
func (p Person) year(currentYear int) int {
	return currentYear - p.Age;
}

//Pointer Recievers
//It copies the address of struct, so it can enable modifications.

func (p *Person) changeName(name string) {
	p.Name = name
}

func main(){

	person := Person{Name: "Robin", Age: 18}

	person.greet()

	fmt.Println()

	person.intro("Warsaw")

	fmt.Println()

	birthYear := person.year(2026)

	fmt.Println("I was born at", birthYear)

	fmt.Println()

	person.changeName("Dean")

	person.intro("Sydney")

}
