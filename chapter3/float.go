package main

import (
    "fmt"
)

func main() {
    // float32とfloat64の2種類ある
    fmt.Printf("%T\n", float32(1.7))

    // 浮動小数点リテラルはfloat64型になる
    fmt.Printf("%T\n", 1.7)

    // 演算不能な場合は特殊な計算結果となる
    zero := 0.0
    fmt.Printf("%v\n", 1.0 / zero)    // +Inf
    fmt.Printf("%v\n", -1.0 / zero)   // -Inf
    fmt.Printf("%v\n", zero / zero)   // NaN

    // eの後ろの整数は基数が10の指数となる
    fmt.Printf("%v\n", 1.2e2)    // 120
    fmt.Printf("%v\n", -1.2e2)   // -120
}
