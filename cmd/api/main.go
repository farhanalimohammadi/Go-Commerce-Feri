package main 

import "fmt"



func main() {
    number := make(chan string , 2)

    go func() {
        number <- "farhan"
		number <- "erfan"
    }()

    value := <- number
	value2:= <- number

    fmt.Println(value , value2 )
}