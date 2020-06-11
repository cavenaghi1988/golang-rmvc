package util

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBabbleSort(t *testing.T) {
	el := []int{9, 8, 7, 6, 5}

	BabbleSort(el)

	assert.NotNil(t, el)
	assert.EqualValues(t, 5, len(el))
	assert.EqualValues(t, 5, el[0])
	assert.EqualValues(t, 6, el[1])
	assert.EqualValues(t, 7, el[2])
	assert.EqualValues(t, 8, el[3])
	assert.EqualValues(t, 9, el[4])
}

func getNum(n int) []int {
	bug := make([]int, n)
	for i := 0; i < n; i++ {
		bug = append(bug, i)
	}

	return bug
}

func BenchmarkBabbleSort(b *testing.B) {
	res := getNum(10)
	for i := 0; i < b.N; i++ {
		SortElemnts(res)
	}
}

func BenchmarkNativeBabbleSort(b *testing.B) {
	res := getNum(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(res)
	}
}
