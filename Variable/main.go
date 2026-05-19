package main

import "fmt"

func main(){

	//Shorthand version
	name := "Robin"
	age := 18

	fmt.Println(name)
	fmt.Println(age)

	//Explicit Version
	var commission float64  = 18.50
	var isAdult bool = true

	fmt.Printf("%f \n", commission);
	fmt.Printf("%t \n", isAdult);	

	//Rune ane Bytes
	var b byte = 65
	fmt.Println(b)
	fmt.Printf("%c\n", b)

	
	var r rune = '🔥'
	fmt.Println(r)
	fmt.Printf("%c\n", r)


	//Iterating By Bytes
	s := "😎"

	for i := 0; i < len(s); i++ {
    		fmt.Println(s[i])
	}


	//Iterating by Runes
	for _, r := range s {
	    fmt.Printf("%c\n", r)
	}

}
