package main

import "fmt"


//Declaring the Structure
type Vigilante struct {
	Name string
	Age int

}


//Structre Embedding
type Address struct {
	City string
}

type Person struct {
	Name string
	Address
}

//Struct tags (Usually Used in Json and API stuff, I am not going deep here)
type User struct {
    Name string `json:"name"` //The backticks contain metadata
    Age  int    `json:"age"`
}

func main() {

	//Intializing the Structure
	vig := Vigilante{
		Name: "Bruce Wayne",
		Age: 31,
	}

	fmt.Println(vig.Name)
	fmt.Println(vig.Age)

	fmt.Println()

	//Copying the struct
	a := Vigilante{"Dick", 21} //Another way of initailizing structure

	b := a

	b.Name = "Tim"

	fmt.Println(a.Name)
	fmt.Println(b.Name)

	//Pointer to the Structure
	//This allows us to store the address of structure which could enable modifications
	v2 := &Vigilante {
		Name: "Oliver Queen",
		Age: 34,
	}

	//Accessing elemenets here is same as before
	fmt.Println(v2.Name) //Internally go does (*v2).Name

	fmt.Println()

	p := Person{
		Name: "Rob",
		Address: Address{
			City: "London",
		},
	}

	fmt.Println(p.City)

	fmt.Println()


	//Anonymous Structs
	person := struct {
    Name string
    Age  int
	} {
    Name: "Rob",
    Age:  21,
	}
}