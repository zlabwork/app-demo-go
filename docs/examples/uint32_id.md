## 创建 uint32 id

```go
func main() {

    for i := 0; i < 10; i++ {
        fmt.Printf("%d\n", NewDistributedId(i))
    }

}

// n is 32 bits, uint32
func NewDistributedId(n int) int {
    s0 := n & 255 << 24
    s1 := n & (255 << 8) << 8
    s2 := n & (255 << 16) >> 8
    s3 := n & (255 << 24) >> 24
    return s0 | s1 | s2 | s3
}
```

`e.g. output`
```text
0
16777216
33554432
50331648
67108864
83886080
100663296
117440512
134217728
150994944
```
