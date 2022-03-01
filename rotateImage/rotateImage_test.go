package rotateImage

import (
	"fmt"
	"testing"
)

type rotateTestCase struct {
	input  [][]int
	output [][]int
}

func (tc rotateTestCase) check() bool {
	if len(tc.input) != len(tc.output) {
		return false
	}
	for i, x := range tc.input {
		ex := tc.output[i]
		for j, v := range x {
			if v != ex[j] {
				return false
			}
		}
	}
	return true
}

func TestRotateImage(t *testing.T) {
	for _, tc := range []rotateTestCase{
		{
			input:  [][]int{{ 5,1,9,11 }, { 2,4,8,10 }, { 13,3,6,7 }, { 15,14,12,16 }},
			output: [][]int{{ 15,13,2,5 }, { 14,3,4,1 }, { 12,6,8,9 }, { 16,7,10,11 }},
		},
		{
			input:  [][]int{{ 2,29,20,26,16,28 }, { 12,27,9,25,13,21 }, { 32,33,32,2,28,14 }, { 13,14,32,27,22,26 }, { 33,1,20,7,21,7 }, { 4,24,1,6,32,34 }},
			output: [][]int{{ 4,33,13,32,12,2 }, { 24,1,14,33,27,29 }, { 1,20,32,32,9,20 }, { 6,7,27,2,25,26 }, { 32,21,22,28,13,16 }, { 34,7,26,14,21,28 }},
		},
	}{
		fmt.Printf("\ninput: %v\n\n", tc.input)
	
		rotate(tc.input)
		
		fmt.Printf("\noutput: %v\n\n", tc.input)

		if !tc.check() {
			t.Errorf("rotation fail, got: %v, want: %v", tc.input, tc.output)
		}
	}
}
