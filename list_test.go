package xlist

import "testing"

func Test_IntMapReduce(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	f := func(prev, cur interface{}, _ int) interface{} {
		return prev.(int) + cur.(int)
	}
	ret := IntSlice(s).ToList().Reduce(f, 1)

	t.Log(ret)
}

func Test_IntFilter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	f := func(cur interface{}, _ int) bool {
		return cur.(int) > 3
	}
	ret := IntSlice(s).ToList().Filter(f)

	t.Log(ret)
}

func Test_IntSort(t *testing.T) {
	s := New(10).Map(WithMapperToIndex())

	less := func(i, j int) bool {
		return s[i].(int) > s[j].(int)
	}
	t.Log(s)
	s.Sort(less)
	t.Log(s)
}

func Test_PushPop(t *testing.T) {
	s := New(0, 10)
	s.Push(1)
	t.Log(s)
	s.Pop()
	t.Log(s)
}

func Test_Find(t *testing.T) {
	s := New(10).Map(func(e interface{}, i int) interface{} {
		return i * 2
	})
	t.Log(s)
	v, i, ok := s.Find(WithEqual(6))
	t.Log(v, i, ok)
}

func Test_Some(t *testing.T) {
	s := New(10).Map(func(e interface{}, i int) interface{} {
		return i / 2
	})
	t.Log(s)
	ok := s.Some(WithEqual(2))
	t.Log(ok)
}

func Test_Every(t *testing.T) {
	s := New(10).Map(func(e interface{}, i int) interface{} {
		return i * 2
	})
	t.Log(s)
	ok := s.Every(func(cur interface{}, i int) bool {
		return cur.(int)%2 == 0
	})
	t.Log(ok)
}

func Test_Random(t *testing.T) {
	s := New(10).Map(WithMapperToIndex())
	t.Log(s)
	v := s.Random()
	t.Log(v)
}

func Test_CreateAndSort(t *testing.T) {
	s := New(10).Map(func(e interface{}, i int) interface{} {
		return rnd.Intn(100)
	})
	t.Log(s)
	less := func(i, j int) bool {
		return s[i].(int) < s[j].(int)
	}
	t.Log(s.Sort(less))
}
