package xlist

type IntSlice []int

func (s IntSlice) Get(i int) (interface{}, bool) {
	return s[i], true
}

func (s IntSlice) Swap(i, j int) {
	v := s[i]
	s[i] = s[j]
	s[j] = v
}

func (s IntSlice) Sort(less func(i, j int) bool) {
	Sort(len(s), less, s.Swap)
}

func (s IntSlice) ToList() List {
	return Map(len(s), s.Get, WithMapper())
}

func (s IntSlice) Map(mapper func(cur interface{}, i int) interface{}) List {
	return Map(len(s), s.Get, mapper)
}

type Int64Slice []int64

func (s Int64Slice) Get(i int) (interface{}, bool) {
	return s[i], true
}

func (s Int64Slice) Swap(i, j int) {
	v := s[i]
	s[i] = s[j]
	s[j] = v
}

func (s Int64Slice) Sort(less func(i, j int) bool) {
	Sort(len(s), less, s.Swap)
}

func (s Int64Slice) ToList() List {
	return Map(len(s), s.Get, WithMapper())
}

func (s Int64Slice) Map(mapper func(cur interface{}, i int) interface{}) List {
	return Map(len(s), s.Get, mapper)
}

type StringSlice []string

func (s StringSlice) Get(i int) (interface{}, bool) {
	return s[i], true
}

func (s StringSlice) Swap(i, j int) {
	v := s[i]
	s[i] = s[j]
	s[j] = v
}

func (s StringSlice) Sort(less func(i, j int) bool) {
	Sort(len(s), less, s.Swap)
}

func (s StringSlice) ToList() List {
	return Map(len(s), s.Get, WithMapper())
}

func (s StringSlice) Map(f func(cur interface{}, i int) interface{}) List {
	return Map(len(s), s.Get, f)
}

type ByteSlice []byte

func (s ByteSlice) Get(i int) (interface{}, bool) {
	return s[i], true
}

func (s ByteSlice) Swap(i, j int) {
	v := s[i]
	s[i] = s[j]
	s[j] = v
}

func (s ByteSlice) Sort(less func(i, j int) bool) {
	Sort(len(s), less, s.Swap)
}

func (s ByteSlice) ToList() List {
	return Map(len(s), s.Get, WithMapper())
}

func (s ByteSlice) Map(f func(cur interface{}, i int) interface{}) List {
	return Map(len(s), s.Get, f)
}
