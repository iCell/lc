import "math"

// 101 也就是 5 的反码就是 010，所以 101 ^ 111 进行异或操作即可求出结果
// 通过移位操作求出 N 的二进制长度
// 2 的 length 次方 - 1 就能得出全是 11111 的结果
func bitwiseComplement(N int) int {
    if N == 0 {
        return 1
    }

    var length int

    x := N
    for x > 0 {
        x = x >> 1
        length++
    }

    t := int(math.Pow(2.0, float64(length))) - 1
    return N ^ t
}

