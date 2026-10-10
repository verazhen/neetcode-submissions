/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	nodeMap := map[*Node]*Node{} // old node to new node
	//traveres and copy node while record the relation of old and new node
	curr := head
	dummy := &Node{}
	newCurr := dummy
	for curr != nil {
		// 1. deep copy
		copied := &Node{
			Val: curr.Val,
		}
		// 2. record
		nodeMap[curr] = copied

		// 3. move and add Next value to copied
		newCurr.Next = copied
		newCurr = newCurr.Next
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		// 1. check newCurr and newRandom
		newCurr := nodeMap[curr]
		newRandom := nodeMap[curr.Random]
		// 2. record newRandom to newCurr
		newCurr.Random = newRandom
		
		curr = curr.Next
	}
	
    return dummy.Next
}
