package main

import "fmt"

func main() {
    /*
    Go Variable Types
    -   int
    -   float32
    -   string
    -   bool
    */

    /*
    Declaring (Creating) Variables

    1. With the var keyword:
    var variablename type = value

    2. With the := sign
    variablename := value
    */

    //
    // Variable Declaration With Initial Value
    //

    // type is string
    var student1 string = "John"
    fmt.Println(student1)

    // type is inferred
    var student2 = "Jane"
    fmt.Println(student2)

    // type is inferred
    x := 2
    fmt.Println(x)

    //
    // Variable Declaration Without Initial Value
    //

    // a is ""
    var a string
    fmt.Println(a)

    // b is 0
    var b int
    fmt.Println(b)

    // c is false
    var c bool
    fmt.Println(c)

    //
    // Value Assignment After Declaration
    //
    var student1 string
    student1 = "John"
    fmt.Println(student1)
}
