package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	fmt.Println("Hello, stats!")
	fi := NewFloatIndex()

	for i := 0; i < 100; i++ {
		fi.Put(float64(rand.Intn(10)))
	}
	fi.PrintOrder()
}

type FloatIndex struct {
	index map[float64]int
	count int
	order sort.Float64Slice
}

func NewFloatIndex() *FloatIndex {
	fi := FloatIndex{
		index: make(map[float64]int),
		count: 0,
		order: make([]float64, 0),
	}
	return &fi
}

func (fi *FloatIndex) Put(f float64) {
	// If this is the first time seeing this number
	if fi.index[f] == 0 {
		// Increment the value in the map
		fi.index[f]++
		// Add it to the ordering slice and insertion sort into place
		fi.order = append(fi.order, f)
		//for i := fi.count; i > 0 && fi.order[i-1] > fi.order[i]; i-- {
		//fi.order.Swap(i, i-1)
		//}
		fi.order.Sort()
	} else {
		// Otherwise simply increment the map value
		fi.index[f]++
	}
	// Finally increment the overall count
	fi.count++
}

func (fi *FloatIndex) Get(f float64) int {
	return fi.index[f]
}

func (fi *FloatIndex) Count() int {
	return fi.count
}

func (fi *FloatIndex) PrintOrder() {
	fmt.Println(fi.order)
}
