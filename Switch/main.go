package main

import "fmt"

func main() {

	//Basic Syntax

	option := 2

	switch option {

	case 1:
		fmt.Println("English")

	case 2:
		fmt.Println("Spanish")

	default:
		fmt.Println("Invalid")
	}

    fmt.Printf("\n")

    //Multiple Value in one case

    grade := "B"

    switch grade {

    case "A", "B":
         fmt.Println("Good")

    case "C":
         fmt.Println("Average")

    default:
        fmt.Println("Needs Improvement")
}

    fmt.Printf("\n")


    //Switch Without Expression

    age := 14

    switch {

    case age >= 18:
         fmt.Println("Adult")

    case age >= 13:
         fmt.Println("Teenager")

    default:
        fmt.Println("Child")
}

    fmt.Printf("\n")

    //Fallthrough - Continuing the statement even after the case is met

    role := "editor"

    fmt.Printf("%s can: \n", role)

    switch role {

    case "admin":
         fmt.Println("Can delete Users")
         fallthrough

    case "editor":
         fmt.Println("Can Edit Content")
         fallthrough

    case "user":
         fmt.Println("Can view Content")

    default:
        fmt.Println("Invalid Role")

}


}
