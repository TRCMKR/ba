package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceDiff(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		sliceA       []int
		sliceB       []int
		expectedDiff []int
	}{
		{
			name:         "sliceA is bigger",
			sliceA:       []int{1, 2, 3},
			sliceB:       []int{1, 10},
			expectedDiff: []int{2, 3},
		},
		{
			name:         "sliceA is empty",
			sliceA:       []int{},
			sliceB:       []int{1, 2, 3},
			expectedDiff: []int{},
		},
		{
			name:         "sliceB is empty",
			sliceA:       []int{1, 2, 3},
			sliceB:       []int{},
			expectedDiff: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			diff := SliceDiff(test.sliceA, test.sliceB)

			assert.Equal(t, test.expectedDiff, diff)
		})
	}
}

func TestSliceFullDiff(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		sliceA       []int
		sliceB       []int
		expectedDiff []int
	}{
		{
			name:         "sliceA is bigger",
			sliceA:       []int{1, 2, 3},
			sliceB:       []int{1, 10},
			expectedDiff: []int{2, 3, 10},
		},
		{
			name:         "sliceA is empty",
			sliceA:       []int{},
			sliceB:       []int{1, 2, 3},
			expectedDiff: []int{1, 2, 3},
		},
		{
			name:         "sliceB is empty",
			sliceA:       []int{1, 2, 3},
			sliceB:       []int{},
			expectedDiff: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			diff := SliceFullDiff(test.sliceA, test.sliceB)

			assert.Equal(t, test.expectedDiff, diff)
		})
	}
}
