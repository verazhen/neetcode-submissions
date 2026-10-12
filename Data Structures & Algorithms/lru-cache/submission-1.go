type LRUCache struct {
    kv map[int]*Node //save the previous node
	head *Node
	tail *Node
	capacity int
}

type Node struct{
	Key int
	Val int
	Next *Node
}

//map to get latest value: key to previous node
//link list: head is the oldest node
func Constructor(capacity int) LRUCache {
    dummy := &Node{}
	
	return LRUCache{
		kv: map[int]*Node{},
		head: dummy,
		tail: dummy,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {

    if prev,ok:= this.kv[key];ok{
		curr := prev.Next
		if curr == this.tail{
			return curr.Val
		}

		next := prev.Next.Next

		prev.Next = next
		if next!=nil{
			this.kv[next.Key] = prev
		}
		this.tail.Next = curr
		this.kv[key] = this.tail
		this.tail = this.tail.Next
		curr.Next = nil

		return curr.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	
	prev := this.kv[key]
	if prev == nil{
		newNode :=&Node{
			Key: key,
			Val: value,
		}
		this.tail.Next = newNode
		this.kv[key] = this.tail
		this.tail = this.tail.Next
	}else{
		curr := prev.Next
		curr.Val = value
		if curr == this.tail{
			return
		}

		next := prev.Next.Next

		prev.Next = next
		if next!=nil{
			this.kv[next.Key] = prev
		}
		this.tail.Next = curr
		this.kv[key] = this.tail
		this.tail = this.tail.Next
		curr.Next = nil
	}


	//check capacity
	if len(this.kv) > this.capacity{
		delete(this.kv, this.head.Next.Key)
		if this.head.Next.Next !=nil{
			this.kv[this.head.Next.Next.Key] = this.head
		}
		this.head.Next = this.head.Next.Next
	}
	return
}
