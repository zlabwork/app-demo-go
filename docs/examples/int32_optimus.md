## Optimus Lib
```go
// NewOptimus opt := NewOptimus(1580030173, 59260789, 1163945558)
func NewOptimus(prime, inverse, random int) *Optimus {
	// https://github.com/jenssegers/optimus
	max := math.Pow(2, 31) - 1
	return &Optimus{
		Prime:   prime,
		Inverse: inverse,
		Xor:     random,
		Max:     int(max),
	}
}

type Optimus struct {
	Prime   int
	Inverse int
	Xor     int
	Max     int
}

func (op *Optimus) Encode(n int) int {
	return ((n * op.Prime) & op.Max) ^ op.Xor
}

func (op *Optimus) decode(n int) int {

	return ((n ^ op.Xor) * op.Inverse) & op.Max
}
```

## Usage
```go
func main() {
	opt := NewOptimus(1580030173, 59260789, 1163945558)
	for i := 0; i < 20; i++ {
		n := opt.Encode(i)
		s := opt.decode(n)
		fmt.Println(s, "=>", n)
	}
}
```


## Ref
[https://github.com/jenssegers/optimus](https://github.com/jenssegers/optimus)
