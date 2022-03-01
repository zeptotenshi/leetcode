package summaryRanges

import (
    "fmt"
)

/* Solution Accepted
	Runtime: 2 ms
		faster than 40.63% of Go online submissions
	Memory Usage: 2 MB
		less than 73.96% of Go online submissions
*/

func summaryRanges(nums []int) []string {
    // check for zero-length input
    if len(nums) == 0 {
        return []string{}
    }
    
    // set current range minimum to first input element
    rangeMin := nums[0]
    // create range map - key: start of range, value: end of range
    ranges := map[int]int{
        rangeMin: rangeMin,
    }
    
    // iterate through sorted int slice input
    for i, n := range nums {
        // ignore first element; already set in map creation
        if i == 0 {
            continue
        }
        // get current range max for number(n)
        rangeMax := ranges[rangeMin]
        // check if current number(n) breaks range
        if n > rangeMax + 1 {
            // reset range minimum to create new range
            rangeMin = n
        }

        // set range maximum
        ranges[rangeMin] = n
    }
    
    // result value
    rv := []string{}
    // temporary string
    ts := ""
    // iterate through input again; leetcode solution-check expects sorted output
    // golang map orderings are random by design
    for _, rs := range nums {
        // get the range max
        re, ok := ranges[rs]
        // verify range start exists in range map
        if ok {
            // i.e. input: [1, 2, 3, 5]     out: ["1->3", "5"]
            // single-digit range case; "5"
            if rs == re {
                ts = fmt.Sprintf("%d", rs)
            
            // multi-digit range case; "1->3"
            } else {
                ts = fmt.Sprintf("%d->%d", rs, re)
            }

            rv = append(rv, ts)
        }
    }
    
    return rv
}
