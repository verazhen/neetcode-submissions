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
	if i == nil{
		return j
	}else if j == nil{
		return i
	}

	dummy := &ListNode{}
	curr := dummy
	for {
		if i == nil{
			curr.Next = j
			return dummy.Next
		}else if j == nil{
			curr.Next = i
			return dummy.Next
		}
		//compare i & j
		//if i/j smaller, move i/j and insert to new node
		var next *ListNode
		if i.Val <= j.Val{
			next = i
			i = i.Next
		}else{
			next = j
			j = j.Next
		}

		curr.Next = next
		curr = curr.Next

	}

	return nil
}
