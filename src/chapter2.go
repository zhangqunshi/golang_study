package main

import "fmt"

func main() {
    var a int = 0;
    var b int8 = 127;
    var c int16 = 32767;
    var d int32 = 2147483647;
    var e int64 = 9223372036854775807;
    fmt.Println(a, b, c, d, e);

    var f uint8 = 255;
    var g uint16 = 65535;
    var h uint32 = 4294967295;
    var i uint64 = 18446744073709551615;
    fmt.Println(f, g, h, i);


    var k float32 = 3.1415;
    var l float64 = 3.1415;
    var m complex64 = 3.1415;
    var n complex128 = 3.1415;
    fmt.Println(k, l, m, n);

    w := 1;
    var y bool = true
    z := false
    fmt.Println(w, y, z);

    var aa, bb int = 1, 3
    fmt.Println(aa, bb);
}
