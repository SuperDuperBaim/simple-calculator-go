package main

import "fmt"

func main(){
	fmt.Println("Calculator")

	var number1 float64
    fmt.Println("Enter the number")
    fmt.Scanln(&number1)

    //operator
    var operator string
    fmt.Println("Enter the operator. (+, *, -, /): ")
    fmt.Scanln(&operator)

    var number2 float64
    fmt.Println("Enter the number")
    fmt.Scanln(&number2)

    var result float64
    fmt.Println(&result)
    
    //Process
    if operator == "+"{
        result  = number1 + number2
    } else if operator == "*"{
        result  = number1 * number2
    } else if operator == "-"{
        result  = number1 - number2
    } else if operator == "/"{
        if number2 != 0 {
            result = number1 / number2
        } else {
            fmt.Println("Error ERROR: Cannot divide by 0.")
            return
        }
    } else {
        fmt.Println("ERROR: Invalid operator!")
        return
    }


    //OUTPUT
     fmt.Printf("Hasil: %.2f\n", result)
}