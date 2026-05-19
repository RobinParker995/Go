package main

import "fmt"

func main(){

	//Shorthand version
	name := "Robin"
	age := 18

    fmt.Println("Shorthand Versions: ")
	fmt.Println(name)
	fmt.Println(age)

    fmt.Printf("\n")

	//Explicit Version
	var commission float64  = 18.50
	var isAdult bool = true

    fmt.Println("Explict Versions: ")
	fmt.Printf("%f \n", commission);
	fmt.Printf("%t \n", isAdult);	

    fmt.Printf("\n")

	//Rune ane Bytes
	var b byte = 65

    fmt.Println("Bytes: ")
	fmt.Println(b)
	fmt.Printf("%c\n", b)

    fmt.Printf("\n")
	
	var r rune = '🔥'

    fmt.Println("Runes: ")
	fmt.Println(r)
	fmt.Printf("%c\n", r)

    fmt.Printf("\n")


    //Iteration
	s := "😎"


    fmt.Println("Iteration by Bytes: ")
	for i := 0; i < len(s); i++ {
    		fmt.Println(s[i])
	}

    fmt.Printf("\n")

	fmt.Println("Iterating by Runes")
	for _, r := range s {
	    fmt.Printf("%c\n", r)
	}

}
