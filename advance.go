package xlist

import (
	mrand "math/rand"
	"os"
	"time"
)

var rnd = mrand.New(mrand.NewSource(time.Now().UnixNano() + int64(os.Getpid())<<10))

func Map(n int, get func(i int) (interface{}, bool), f func(cur interface{}, i int) interface{}) List {
	var ret = New(0, n)
	for i := 0; i < n; i++ {
		v, _ := get(i)
		ret = append(ret, f(v, i))
	}
	return ret
}

func Reduce(n int, get func(i int) (interface{}, bool), f func(prev, cur interface{}, i int) interface{}, initialVal interface{}) interface{} {
	var prev = initialVal
	var cur interface{}
	for i := 0; i < n; i++ {
		cur, _ = get(i)
		prev = f(prev, cur, i)
	}
	return prev
}

func Filter(n int, get func(i int) (interface{}, bool), f func(cur interface{}, i int) bool) List {
	var ret = New(0, n)
	for i := 0; i < n; i++ {
		v, ok := get(i)
		if ok && f(v, i) {
			ret = append(ret, v)
		}
	}
	return ret
}

func ForEach(n int, get func(i int) (interface{}, bool), f func(cur interface{}, i int) bool) {
	for i := 0; i < n; i++ {
		v, ok := get(i)
		if ok && !f(v, i) {
			break
		}
	}
	return
}

func Some(n int, get func(i int) (interface{}, bool), f func(cur interface{}, i int) bool) bool {
	for i := 0; i < n; i++ {
		v, ok := get(i)
		if ok && f(v, i) {
			return true
		}
	}
	return false
}

func Every(n int, get func(i int) (interface{}, bool), f func(cur interface{}, i int) bool) bool {
	for i := 0; i < n; i++ {
		v, ok := get(i)
		if ok && !f(v, i) {
			return false
		}
	}
	return true
}

func Find(n int, get func(i int) (interface{}, bool), f func(cur interface{}, i int) bool) (interface{}, int, bool) {
	for i := 0; i < n; i++ {
		v, ok := get(i)
		if ok && f(v, i) {
			return v, i, true
		}
	}
	return nil, -1, false
}

func Sort(n int, less func(i, j int) bool, swap func(i, j int)) {
	insertSort(n, less, swap)
}

func Random(n int, swap func(i, j int)) {
	for i := 0; i < n; i++ {
		j := i + rnd.Intn(n-i)
		swap(i, j)
	}
}

func insertSort(n int, less func(i, j int) bool, swap func(i, j int)) {
	for i := 1; i < n; i++ {
		for j := i; j >= 1 && less(j, j-1); j-- {
			swap(j, j-1)
		}
	}
}
