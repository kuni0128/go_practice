package main

import (
    "fmt"
)

func main() {
    // 以下はコンパイルエラー
    /*
    a := [3]string{
        "Hokkaido",
        "Chiba",
        "Tokyo"
    }
    */
    a := [3]string{
        "Hokkaido",
        "Chiba",
        "Tokyo",
    }
    fmt.Printf("%v", a)
    fmt.Println("")
    fmt.Printf("%T", a)
    fmt.Println("")
    fmt.Printf("%#v", a)
}
