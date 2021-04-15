package main

type IntSet struct {
	bit []uint64
}
func (is *IntSet) String() string {
	return string(is.bit[0])
}

func main()  {
	is := IntSet{bit: []uint64{1}}
	// 正确
	var _ = is.String()
	// 错误
	var _ = IntSet{[]uint64{1}}.String()
}

