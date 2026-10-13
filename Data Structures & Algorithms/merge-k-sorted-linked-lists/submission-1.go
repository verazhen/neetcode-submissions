/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {

	length := len(lists)

	if length == 0{
		return nil
	}
	for length > 1{
		l := 0
		r := length - 1
		for l < r {
			//merge list to the left listNode
			dummy := &ListNode{}
			curr := dummy
			i := lists[l]
			j := lists[r]
			for i != nil || j != nil{
				if i == nil{
					curr.Next = j
					break
				}else if j == nil{
					curr.Next = i
					break
				}


				if i.Val <= j.Val{
					curr.Next = i
					i = i.Next
				}else{
					curr.Next = j
					j = j.Next
				}
				curr = curr.Next
			}
			lists[l] = dummy.Next 

			l++
			r--
		}

		length = (length+1)/2
	}
	
	return lists[0]
}
