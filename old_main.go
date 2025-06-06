package main

import (
	"fmt"
	"go-learning/customModule"
	"reflect"
)

func isEven(num int) bool {
    if num%2 == 0 {
        return true
    }
    return false
}


func main() {
    num := 10
    isNumEven := isEven(num)
    fmt.Println("Is even", isNumEven)
    sum := customModule.Add(10, 20)
    fmt.Println("Sum of two:", sum)
	fmt.Println("Type of ", reflect.TypeOf(nil))
}
