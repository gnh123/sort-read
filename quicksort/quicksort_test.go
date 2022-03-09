package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	data []int
	need []int
}

func TestSelectSort(t *testing.T) {
	for _, tc := range []testCase{
		{
			data: []int{3, 7, 2, 1},
			need: []int{1, 2, 3, 7},
		},
		{
			data: []int{1, 2, 3, 4},
			need: []int{1, 2, 3, 4},
		},
		{
			data: []int{4, 3, 2, 1},
			need: []int{1, 2, 3, 4},
		},
		{
			data: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			need: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	} {

		quicksort(tc.data)
		assert.Equal(t, tc.data, tc.need)
	}
}
