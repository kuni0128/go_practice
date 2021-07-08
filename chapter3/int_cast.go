package main

import (
    "fmt"
)

func main() {
    // 暗黙的にint型の変数になる
    a := 100
    fmt.Printf("%T\n", a)

    // 型変換を使うとuint型の変数を宣言できる
    b := uint(100)
    fmt.Printf("%T\n", b)

    c := 100
    d := byte(c)
    e := uint32(c)
    fmt.Printf("%v %v\n", d, e)

    // 最大値255を超えた値を代入しようとすると、コンパイルエラー
    // x := byte(256)

    // 型変換した結果桁あふれする場合は、0になる（ラップアラウンド）
    y := 256
    fmt.Printf("%v\n", byte(y))

    // 負の整数-1をbyte型に変換した結果、255になる（ラップアラウンド）
    z := -1
    fmt.Printf("%v\n", byte(z))
}
