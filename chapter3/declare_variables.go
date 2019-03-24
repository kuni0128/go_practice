package main

import (
    "fmt"
)

func one() int {
    return 1
}

func main() {
    /*
     *  明示的な変数宣言
     */
    var n int
    var x, y, z int
    var (
        a int
        b string
    )
    //異なる型を代入しようとすると、 コンパイルエラー
    n = "string"

    // 異なる変数の数を代入しようとすると、コンパイルエラー
    x, y, z = 1, 2

    // 同じ変数を再度定義しようとすると、コンパイルエラー
    var n int

    /*
     *  暗黙的な変数宣言
     */
    max := 1
    prefecture := "Hokkaido"
    pi := 3.14
    is_lie := true

    // 関数の戻り値の型でそのまま変数宣言
    num := one()

    // 同じ変数を再度定義しようとすると、コンパイルエラー
    max := 2

    // varでまとめて暗黙的な変数宣言もできる
    var (
        name = "Kuniaki"
        birthday = 19850128
    )
}
