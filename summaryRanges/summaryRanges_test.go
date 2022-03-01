package summaryRanges

import (
	"testing"
)

type summaryRangeTestCase struct {
	input  []int
	output []string
}

func (tc summaryRangeTestCase) checkResult(_r []string) bool {
	if len(_r) != len(tc.output) {
		return false
	}

	for i, ov := range tc.output {
		if ov != tc.output[i] {
			return false
		}
	}

	return true
}

func TestSummaryRanges(t *testing.T) {
	for _, tc := range []summaryRangeTestCase{
		{
			input:  []int{ 0, 1, 2, 4, 5, 7 },
			output: []string{ "0->2", "4->5", "7" },
		},
		{
			input:  []int{ 0, 2, 3, 4, 6, 8, 9 },
			output: []string{ "0", "2->4", "6", "8->9" },
		},
	}{
		result := summaryRanges(tc.input)

		if !tc.checkResult(result) {
			t.Errorf("unexpected result, got: %v, want: %v", result, tc.output)
		}
	}
}
