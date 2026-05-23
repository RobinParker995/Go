package main

import "fmt"


//Basic Function
func greet(name string){

	//This is a Void Function
	fmt.Printf("Hello There, %s \n", name)
}

//Function with paramters and a return type
func add(a, b int) int{
	
	return a + b
}

//Function with Multiple Return Values
func divide(a, b float64) (float64, string){

	if b == 0 {
		return 0, "Cant't Divide by Zero"
	}

	return a / b, "Success"
}


//Function as paramter
func operate(a, b int, op func(int, int) int) {
	fmt.Println(op(a, b))
}

//Closure
func counter() func() int {

	count := 0

	return func() int {
		count++
		return count
	}
}

//Variadic Functions
func sum(nums ...int) int {

	total := 0

	for _,n := range nums {
		total += n
	}

	return total
}


func main(){

	greet("Obi Wan Kenobi")
	result := add(4, 4)

	fmt.Println(result)

	res, msg := divide(10, 2)

	fmt.Printf("%.2f, %s \n", res, msg)

	//Anoynmous Functions
	sayHello := func(){
		fmt.Println("Hello")
	}

	sayHello()

	operate(25, 32, add)

	c := counter()

	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	fmt.Println(sum(1,2,3,4,5))

}