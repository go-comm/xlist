package xlist

type Getter interface {
	Get(i int) (interface{}, bool)
}

type Setter interface {
	Set(i int, x interface{}) interface{}
}

func WithMapper() func(e interface{}, i int) interface{} {
	return func(e interface{}, i int) interface{} {
		return e
	}
}

func WithMapperToIndex() func(e interface{}, i int) interface{} {
	return func(e interface{}, i int) interface{} {
		return i
	}
}

func WithEqual(x interface{}) func(cur interface{}, i int) bool {
	return func(cur interface{}, i int) bool {
		return cur == x
	}
}

// New
func New(length int, capacity ...int) List {
	var cap = length
	if len(capacity) > 0 {
		cap = capacity[0]
	}
	return make(List, length, cap)
}

func From(args ...interface{}) List {
	return List(args)
}

type List []interface{}

func (l *List) Push(x ...interface{}) List {
	*l = append(*l, x...)
	return *l
}

func (l *List) Pop() (interface{}, bool) {
	n := len(*l)
	if n <= 0 {
		return nil, false
	}
	v := (*l)[n-1]
	(*l)[n-1] = nil
	*l = (*l)[:n-1]
	return v, true
}

func (l List) Len() int {
	return len(l)
}

func (l List) Get(i int) (interface{}, bool) {
	if len(l) == 0 {
		return nil, false
	}
	return l[i], true
}

func (l List) Swap(i, j int) {
	v := l[i]
	l[i] = l[j]
	l[j] = v
}

func (l List) Map(f func(cur interface{}, i int) interface{}) List {
	return Map(len(l), l.Get, f)
}

func (l List) Reduce(f func(prev, cur interface{}, i int) interface{}, initialVal interface{}) interface{} {
	return Reduce(len(l), l.Get, f, initialVal)
}

func (l List) Filter(f func(cur interface{}, i int) bool) List {
	return Filter(len(l), l.Get, f)
}

func (l List) Sort(less func(i, j int) bool) List {
	Sort(len(l), less, l.Swap)
	return l
}

func (l List) ForEach(f func(cur interface{}, i int) bool) {
	ForEach(len(l), l.Get, f)
}

func (l List) Some(f func(cur interface{}, i int) bool) bool {
	return Some(len(l), l.Get, f)
}

func (l List) Every(f func(cur interface{}, i int) bool) bool {
	return Every(len(l), l.Get, f)
}

func (l List) Find(f func(cur interface{}, i int) bool) (interface{}, int, bool) {
	return Find(len(l), l.Get, f)
}

func (l List) Random() List {
	Random(len(l), l.Swap)
	return l
}
