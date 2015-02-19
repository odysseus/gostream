package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

// When you put a number into the struct
// 1. The total count should increment by 1
// 2. The order array should always remain sorted
// 3. The value associated with the key in the index should increment
func TestPut(t *testing.T) {
	assert := assert.New(t)
	fi := NewFloatIndex()
	for i := 1; i <= 100; i++ {
		n := float64(rand.Intn(100))
		fi.Put(n)
		assert.Equal(i, fi.count)
		assert.True(sort.Float64sAreSorted(fi.order))
		assert.NotEqual(0, fi.index[n])
	}
}

// CountTotal should return the total number of values
// not the number of unique values
func TestPutCountTotal(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		expected float64
		input    []float64
	}{
		{4, []float64{1, 2, 3, 4}},
		{6, []float64{1, 1, 1, 2, 3, 4}},
	}
	for _, c := range cases {
		fi := NewFloatIndex()
		for _, f := range c.input {
			fi.Put(f)
		}
		assert.Equal(c.expected, fi.CountTotal())
	}
}

// CountUnique should only return the number of unique
// values, not the total # of values inserted
func TestPutCountUnique(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		expected float64
		input    []float64
	}{
		{4, []float64{1, 2, 3, 4}},
		{4, []float64{1, 1, 1, 2, 3, 4}},
	}
	for _, c := range cases {
		fi := NewFloatIndex()
		for _, f := range c.input {
			fi.Put(f)
		}
		assert.Equal(c.expected, fi.CountUnique())
	}
}

// The built in count should match the sum of the values
// stored in the index
func TestCountParity(t *testing.T) {
	assert := assert.New(t)
	fi := NewFloatIndex()
	tot := 0

	for i := 0; i < 100; i++ {
		n := float64(rand.Intn(1000))
		fi.Put(n)
	}

	for _, v := range fi.index {
		tot += v
	}
	assert.Equal(tot, fi.count)
}
