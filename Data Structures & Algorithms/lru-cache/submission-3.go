type LRUCache struct {
	kv       map[int]*Node
	head     *Node // dummy head
	tail     *Node // dummy tail
	capacity int
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}

	head.Next = tail
	tail.Prev = head

	return LRUCache{
		kv:       map[int]*Node{},
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

// Remove node from its current position.
func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// Insert node right before dummy tail.
// The inserted node becomes MRU.
func (this *LRUCache) insertMRU(node *Node) {
	last := this.tail.Prev

	last.Next = node
	node.Prev = last

	node.Next = this.tail
	this.tail.Prev = node
}

// Move an existing node to MRU.
func (this *LRUCache) moveToMRU(node *Node) {
	this.remove(node)
	this.insertMRU(node)
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.kv[key]
	if !ok {
		return -1
	}

	this.moveToMRU(node)

	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	// Existing key: update value and mark as MRU.
	if node, ok := this.kv[key]; ok {
		node.Val = value
		this.moveToMRU(node)
		return
	}

	// New key: insert as MRU.
	node := &Node{
		Key: key,
		Val: value,
	}

	this.kv[key] = node
	this.insertMRU(node)

	// Remove LRU if capacity is exceeded.
	if len(this.kv) > this.capacity {
		lru := this.head.Next

		this.remove(lru)
		delete(this.kv, lru.Key)
	}
}