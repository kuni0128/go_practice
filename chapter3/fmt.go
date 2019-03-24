package main

import (
    "fmt"
)

func main() {
    // 標準出力に表示する
    fmt.Println("Printlnは改行付きで表示する")
    fmt.Println("Printlnは", "複数", "文字列を", "半角スペースで", "連結表示する")

    fmt.Printf("10進数=%d 2進数=%b 8進数=%o 16進数=%x\n", 17, 17, 17, 17)
    fmt.Printf("%d年%d月%d日\n", 2019, 3)
    fmt.Printf("%d年%d月%d日", 2019, 3, 20, 10)
    fmt.Printf("\n")

    fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 77, "Hokkaido", [...]int{1, 2, 3})
    fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 77, "Hokkaido", [...]int{1, 2, 3})
    fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 77, "Hokkaido", [...]int{1, 2, 3})

    // 標準エラー出力に表示する
    print("printは改行無しで表示する")
    println("printは改行有りで表示する")
    println("printlnは", "複数", "文字列を", "半角スペースで", "連結表示する")
}
