/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil{
		return nil
	}


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
	newHead := head.Next 
	for curr != nil{
		copied := curr.Next 
		next := curr.Next.Next
		if curr.Next.Next != nil{
			copied.Next = curr.Next.Next.Next 
		}

		//split two list
		curr.Next = next
		curr = curr.Next

	}
	return newHead

}