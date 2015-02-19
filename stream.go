package main

import (
	"fmt"
	"sort"
)

/// MAIN ///
func main() {
	fmt.Println("Hello, stats!")
}

/// STRUCT DEFINITON ///
type FloatIndex struct {
	index map[float64]int
	count int
	order sort.Float64Slice
}

// Constructor
func NewFloatIndex() *FloatIndex {
	fi := FloatIndex{
		index: make(map[float64]int),
		count: 0,
		order: make([]float64, 0),
	}
	return &fi
}

/// STRUCT IMPLEMENTATION

// Adds a value to the index
func (fi *FloatIndex) Put(f float64) {
	// If this is the first time seeing this number
	if fi.index[f] == 0 {
		// Increment the value in the map
		fi.index[f]++
		// Add it to the ordering slice and insertion sort into place
		fi.order = append(fi.order, f)
		for i := len(fi.order) - 1; i > 0 && fi.order[i-1] > fi.order[i]; i-- {
			fi.order.Swap(i-1, i)
		}
	} else {
		// Otherwise simply increment the map value
		fi.index[f]++
	}
	// Finally increment the overall count
	fi.count++
}

// String representation of the struct - An ordered list of
// keys followed by their value
func (fi *FloatIndex) String() string {
	s := "{ "
	for _, v := range fi.order {
		s += fmt.Sprintf("%v:%v ", v, fi.index[v])
	}
	s += "}"
	return s
}

// Returns the grand total of numbers inserted into the struct
func (fi *FloatIndex) CountTotal() int {
	return fi.count
}

// Returns the count of unique numbers placed into the struct
func (fi *FloatIndex) CountUnique() int {
	return len(fi.order)
}

// Prints the ordered set of values in the struct
func (fi *FloatIndex) StringOrder() string {
	return fmt.Sprint(fi.order)
}

// Minimum by value, not number of occurences
func (fi *FloatIndex) Min() float64 {
	return fi.order[0]
}

// Maximum by value, not number of occurences
func (fi *FloatIndex) Max() float64 {
	return fi.order[len(fi.order)-1]
}

// Average of the values
func (fi *FloatIndex) Mean() float64 {
	mean := float64(0)
	for k, v := range fi.index {
		mean += k * float64(v)
	}
	mean /= float64(fi.count)
	return mean
}

// Median of the values
func (fi *FloatIndex) Median() float64 {
	mid := fi.count / 2
	i := 0
	for _, v := range fi.order {
		i += fi.index[v]
		if i >= mid {
			return v
		}
	}
	return -1
}
