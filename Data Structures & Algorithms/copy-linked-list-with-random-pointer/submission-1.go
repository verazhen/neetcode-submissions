/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    // A -> A' -> B -> B'->C ->C'
	curr := head
	for curr != nil{
		copied := &Node{
			Val: curr.Val,
			Next: curr.Next,
		}
		curr.Next = copied

		curr = copied.Next
	}

	curr = head
	for curr != nil{
		copied := curr.Next
		if curr.Random != nil{
			copied.Random = curr.Random.Next
		}

		curr = curr.Next.Next
	}


	curr = head 
	dummy := &Node{}
	copyPrev := dummy
	for curr != nil{
		copied := curr.Next 
		copyPrev.Next = copied

		//split two list
		curr.Next = curr.Next.Next
		curr = curr.Next 

		copyPrev = copied

	}
	return dummy.Next

}