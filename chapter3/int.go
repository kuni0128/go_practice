package main

func main() {
    /*
     * 符号付き整数型
     */
    var a int8
    var b int16
    var c int32
    var d int64
    println(a, b, c, d)

    /*
     * 符号無し整数型
     */
    var aa uint8
    var bb uint16
    var cc uint32
    var dd uint64
    println(aa, bb, cc, dd)

    /*
     * 実装依存の整数型
     */
    var x int
    var y uint
    var z uintptr
    println(x, y, z)

    // int64の最大値なので、32ビット環境ではコンパイルエラー
    n := 9223372036854775807
    println(n)

    // 異なる型なので代入しようとすると、コンパイルエラー
    // goは暗黙的な型変換を一切認めない
    var (
        n1 int
        n2 int64
    )
    n1 = 1
    n2 = n1
}
