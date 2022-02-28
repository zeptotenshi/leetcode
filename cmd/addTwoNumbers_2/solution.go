/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// current node
	temp := &ListNode{}
	// set result list head
	result := temp
	// carry value
	carry := 0

	// WARNING: verify valid loop exit-condition 
    for {
		// add carry-in to sum
        sum := carry
		// check if list one node exists
        if l1 != nil {
			// add list one node value to sum
            sum += l1.Val
        }
		// repeast for list two
        if l2 != nil {
            sum += l2.Val
        }

		// leetcode solution-check expects single digit for node value; set the current node value to remainder of sum / 10
        temp.Val = sum % 10
		// set the carry out value for next iteration
        carry = (sum - temp.Val) / 10

		// set list one node to next node
        if l1 != nil {
            l1 = l1.Next
        }
		// set list two node to next node
        if l2 != nil {
            l2 = l2.Next
        }
		// exit loop if both nodes for respective lists are nil
        if l1 == nil && l2 == nil {
			// check for non-zero carry-out
            if carry != 0 {
				// add node to result list
                temp.Next = &ListNode{
                    Val: carry,
                }
            }
            break
        }

		// create next node
        temp.Next = &ListNode{}
		// set next node to currently active node
        temp = temp.Next
    }
    
    return result
}

/* Solution Accepted
	Runtime: 4 ms
		faster than 97.99% of Go online submissions
	Memory Usage: 4.6 MB
		less than 70.24% of Go online submissions
*/