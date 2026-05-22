package main

import "fmt"

func main(){

	nums := []int {1, 2, 3, 4} //Note that to declare a slice, you must keep [] empty; otherwise its an array

	fmt.Println(nums) //You can access Slices like Arrays

	fmt.Println()

	//Adding elements in the array
	nums = append(nums, 5) //This adds an element at last
	fmt.Println(nums)

	fmt.Println()

	//Slice Operator
	sub := nums[1:3] //Only include elements from index 1 to 3 from nums slice, note that it does not include the upper index

	fmt.Println(sub)

	fmt.Println()

	//Shorthands

	sub1 := nums[:3] //Starts from start but stops before 3
	sub2 := nums[2:] //Starts from 2 till the end

	fmt.Println(sub1)
	fmt.Println(sub2)

	fmt.Println()

	//Important Nuance
	x := []int {10, 20, 30, 40, 50}
	y := x[2:]

	y[0] = 100 //It also changes 2nd Index of orignal x Slice. This happend because slices shares underlaying arrays

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println()

	//Creating Slice with make
	arr := make([]int, 4, 5) //The syntax goes as, make(datatype, lenght, capacity)
	arr2 := []int {100, 200, 300, 400}

	copy(arr, arr2) //Copies the arr2 to arr, and creats a new slice completely, they don't share an underlying array, so change in one won't affect other

	fmt.Println(arr)
	fmt.Println(arr2)

}