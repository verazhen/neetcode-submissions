type LRUCache struct {
    kv map[int]*Node
	head *Node
	tail *Node
	capacity int
}

type Node struct{
	Key int
	Val int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	dummy := &Node{}
	dummyTail := &Node{}
	dummy.Next = dummyTail
	dummyTail.Prev = dummy
	
	return LRUCache{
		kv: map[int]*Node{},
		head: dummy,
		tail: dummyTail,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
    if curr,ok:= this.kv[key];ok{
		//if exist, update order

		prev := curr.Prev
		next := curr.Next // might be empty node
		prev.Next = next
		next.Prev = prev

		last := this.tail.Prev
		last.Next = curr
		curr.Next = this.tail
		this.tail.Prev = curr
		curr.Prev = last

		return curr.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
    curr,ok:= this.kv[key]
	if !ok{
		//append new node
		last := this.tail.Prev
		newNode := &Node{
			Key: key,
			Val: value,
			Prev: last,
			Next: this.tail,

		}
		this.tail.Prev = newNode
		last.Next = newNode
		this.kv[key] = newNode

		//if > cap, remove element
		if len(this.kv) > this.capacity{
			delete(this.kv, this.head.Next.Key)
			next := this.head.Next.Next
			next.Prev = this.head
			this.head.Next = next
		}
		return
	}

	//if exist, and not the latest, update order
	curr.Val = value

	prev := curr.Prev
	next := curr.Next // might be empty node
	prev.Next = next
	next.Prev = prev

	last := this.tail.Prev
	last.Next = curr
	curr.Next = this.tail
	this.tail.Prev = curr
	curr.Prev = last

	return 
	
}
