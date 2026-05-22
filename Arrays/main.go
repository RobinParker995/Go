package main

import "fmt"

func main(){

    //Array declaration and Access
    nums := [3]int {10, 20, 30}

    fmt.Println(nums[0])
    fmt.Println(nums)
    fmt.Println(len(nums))

    fmt.Println()

    //Iterating the array

    for i := 0; i < len(nums); i++ {

        fmt.Printf("%d ", nums[i])
    }

        fmt.Println()
        fmt.Println()

    //Using range for iterating the array

    for i, e := range nums {

        fmt.Printf("%d %d \n", i, e)
    }
}
