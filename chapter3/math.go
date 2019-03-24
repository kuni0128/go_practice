package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Printf("int8 min value = %v\n", math.MinInt8)
    fmt.Printf("int8 max value = %v\n", math.MaxInt8)
    fmt.Printf("int16 min value = %v\n", math.MinInt16)
    fmt.Printf("int16 max value = %v\n", math.MaxInt16)
    fmt.Printf("int32 min value = %v\n", math.MinInt32)
    fmt.Printf("int32 max value = %v\n", math.MaxInt32)

    fmt.Printf("uint8 max value = %v\n", math.MaxUint8)
    fmt.Printf("uint16 max value = %v\n", math.MaxUint16)
    fmt.Printf("uint32 max value = %v\n", math.MaxUint32)
    fmt.Printf("uint64 max value = %v\n", uint64(math.MaxUint64))
}
