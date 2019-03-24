package main

import (
    "fmt"
)

// 全てのmainパッケージで参照できる変数
var n = 100

func main() {
    n = n + 1
    fmt.Println(n)
}
