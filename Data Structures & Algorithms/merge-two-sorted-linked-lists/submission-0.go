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
	var curr *ListNode
	if i == nil{
		curr = j
		return curr
	}else if j == nil{
		curr = i
		return curr
	}

	var currHead *ListNode
	for i != nil || j != nil {

		if i == nil{
			curr.Next = j
			return currHead
		}else if j == nil{
			curr.Next = i
			return currHead
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

		if currHead == nil{
			currHead = next
			curr = next
			continue
		}

		curr.Next = next
		curr = curr.Next

	}

	return nil
}
