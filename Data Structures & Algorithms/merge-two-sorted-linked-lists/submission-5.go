/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    i := list1
    j := list2
	// if i == nil{
	// 	return j
	// }else if j == nil{
	// 	return i
	// }

	dummy := &ListNode{}
	curr := dummy
	for  i != nil && j != nil {
		//compare i & j
		//if i/j smaller, move i/j and insert to new node
		if i.Val <= j.Val{
			curr.Next = i
			i = i.Next
		}else{
			curr.Next = j
			j = j.Next
		}

		curr = curr.Next

	}

	if i == nil{
		curr.Next = j
	}else if j == nil{
		curr.Next = i
	}

	return dummy.Next
}
