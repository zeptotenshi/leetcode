package addTwoNumbers

import (
	"testing"
)

type addTwoNumbersTestCase struct {
	nl1 []int
	nl2 []int
	output []int
}

func (tc addTwoNumbersTestCase) checkResult(_r *ListNode) bool {
	rl := _r
	for _, ov := range tc.output {
		if rl == nil || rl.Val != ov {
			return false
		}
		rl = rl.Next
	}

	return true
}

func TestAddTwoNumbers(t *testing.T) {
	for _, tc := range []addTwoNumbersTestCase{
		{
			nl1:    []int{ 2, 4, 3 },
			nl2:    []int{ 5, 6, 4 },
			output: []int{ 7, 0, 8 },
		},
		{
			nl1:    []int{ 0 },
			nl2:    []int{ 0 },
			output: []int{ 0 },
		},
		{
			nl1:    []int{ 9, 9, 9, 9, 9, 9, 9 },
			nl2:    []int{ 9, 9, 9, 9 },
			output: []int{ 8, 9, 9, 9, 0, 0, 0, 1 },
		},
	}{
		nl1 := NewListNode(tc.nl1)
		nl2 := NewListNode(tc.nl2)

		result := addTwoNumbers(nl1, nl2)

		if !tc.checkResult(result) {
			t.Errorf("unexpected result, got: %s, want: %v", result, tc.output)
		}
	}
}
