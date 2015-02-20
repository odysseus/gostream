package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

type Expectation struct {
	expected float64
	data     []float64
}

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
	cases := []Expectation{
		{4, []float64{1, 2, 3, 4}},
		{6, []float64{1, 1, 1, 2, 3, 4}},
	}
	for _, c := range cases {
		fi := NewFloatIndex()
		for _, f := range c.data {
			fi.Put(f)
		}
		assert.Equal(c.expected, fi.CountTotal())
	}
}

// CountUnique should only return the number of unique
// values, not the total # of values inserted
func TestPutCountUnique(t *testing.T) {
	assert := assert.New(t)
	cases := []Expectation{
		{4, []float64{1, 2, 3, 4}},
		{4, []float64{1, 1, 1, 2, 3, 4}},
	}
	for _, c := range cases {
		fi := NewFloatIndex()
		for _, f := range c.data {
			fi.Put(f)
		}
		assert.Equal(c.expected, fi.CountUnique())
	}
}

func TestIndex(t *testing.T) {
	assert := assert.New(t)
	cases := []Expectation{
		{2, []float64{1, 2, 2, 5, 5, 8, 12}},
		{3, []float64{1, 1, 3, 4, 3, 3, 5, 2, 1}},
	}

	for _, c := range cases {
		fi := NewFloatIndexFromSlice(c.data)
		assert.Equal(c.expected, fi.index[c.expected])
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

func TestMin(t *testing.T) {
	assert := assert.New(t)
	cases := []Expectation{
		{1, []float64{1, 3, 4, 1, 5, 8, 9}},
		{4, []float64{9, 5, 6, 4, 5, 10, 22}},
	}

	for _, c := range cases {
		fi := NewFloatIndexFromSlice(c.data)
		assert.Equal(c.expected, fi.Min())
	}
}

func TestMax(t *testing.T) {
	assert := assert.New(t)
	cases := []Expectation{
		{9, []float64{5, 6, 7, 8, 9}},
		{21, []float64{21, 13, 14, 5, 6}},
	}

	for _, c := range cases {
		fi := NewFloatIndexFromSlice(c.data)
		assert.Equal(c.expected, fi.Max())
	}
}

func TestMean(t *testing.T) {
	assert := assert.New(t)
	cases := []Expectation{
		{4.5, []float64{4, 5, 3, 4, 5, 6}},
		{3.5, []float64{1, 2, 3, 4, 5, 6}},
	}

	for _, c := range cases {
		fi := NewFloatIndexFromSlice(c.data)
		assert.Equal(c.expected, fi.Mean())
	}
}

func TestMedian(t *testing.T) {
	assert := assert.New(t)
	cases := []Expectation{
		{5, []float64{3, 4, 4, 5, 6, 6, 7}},
		{4, []float64{1, 2, 3, 4, 5, 6}},
	}

	for _, c := range cases {
		fi := NewFloatIndexFromSlice(c.data)
		assert.Equal(c.expected, fi.Median())
	}
}

func TestExactMedian(t *testing.T) {
	assert := assert.New(t)
	cases := []Expectation{
		{5, []float64{3, 4, 4, 5, 6, 6, 7}},
		{3.5, []float64{1, 2, 3, 4, 5, 6}},
	}

	for _, c := range cases {
		fi := NewFloatIndexFromSlice(c.data)
		assert.Equal(c.expected, fi.ExactMedian())
	}
}

func TestMode(t *testing.T) {
	assert := assert.New(t)
	cases := []Expectation{
		{3, []float64{1, 2, 3, 3, 4, 5, 6, 7}},
		{5, []float64{5, 3, 2, 1, 5, 6, 6, 5, 5, 5}},
	}

	for _, c := range cases {
		fi := NewFloatIndexFromSlice(c.data)
		assert.Equal(c.expected, fi.Mode())
	}
}
