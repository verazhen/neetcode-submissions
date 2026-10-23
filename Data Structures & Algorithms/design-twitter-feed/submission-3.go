type Twitter struct {
    followeeMap map[int]map[int]struct{}
    postMap map[int][]post
	index int
}
type post struct{
	id int
	ts int
}
type node struct {
    post       post
    queueIndex int
    pointer  int
}


func Constructor() Twitter {
    return Twitter{
		followeeMap:map[int]map[int]struct{}{},
		postMap:map[int][]post{},
		index: 0,
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
	newPost := post{
		id: tweetId,
		ts: this.index,
	}
	this.index++
	this.postMap[userId] = append(this.postMap[userId],newPost)
	if len(this.postMap[userId]) > 10{
		this.postMap[userId] = this.postMap[userId][1:]
	}
}


func (this *Twitter) GetNewsFeed(userId int) []int {
	//get users
	users := []int{userId}
	for followee,_:=range this.followeeMap[userId]{
		users = append(users,followee)
	}

	postList := [][]post{}
	maxHeap:=[]node{}
	//get users posts
	queueIndex:=0
	for _,user:=range users{
		posts := this.postMap[user]
		if len(posts) == 0{
			continue
		}
		postList = append(postList, posts)
		maxHeap = append(maxHeap, node{
				post: posts[len(posts)-1],
				queueIndex:queueIndex,
				pointer:len(posts)-1,
			})
		queueIndex++
	}
	//heapify(sift down)
	maxHeap= heapify(maxHeap)

	//10 rounds to found the latest posts
	result:=[]int{}
	//kth user
	//pointer := pointers[k]
	//postToCompare := postList[k][pointer]
	for len(result) < 10 && len(maxHeap) > 0{
		//get the root post and k
		result = append(result,maxHeap[0].post.id)
		queueIndex := maxHeap[0].queueIndex
		pointer := maxHeap[0].pointer
		pointer--
		maxHeap[0].pointer = pointer
		//1. move the kth user's pointer and add pointer's post to the heap
		//2. if pointer = 0, move the last element
		//3. if nothings to move, break
		if pointer >= 0{
			post := postList[queueIndex][pointer]
			maxHeap[0] = node{
				post: post,
				queueIndex: queueIndex,
				pointer:    pointer, 
			}
		}else if len(maxHeap) > 0{
			maxHeap[0] = maxHeap[len(maxHeap)-1]
			maxHeap = maxHeap[:len(maxHeap)-1]
		}else{
			break
		}
		//siftdown
		maxHeap = siftDown(maxHeap,0)
	}

	return result
}

func heapify(heap []node)[]node{
	for i:=len(heap)/2-1; i >= 0;i--{
		heap = siftDown(heap,i)
	}
	return heap
}

func siftDown(heap []node, i int)[]node{
	for i < len(heap){
		root := i
		if 2*i+1 >= len(heap){
			break
		}
		child := 2*i+1
		if 2*i+2 < len(heap) && heap[2*i+2].post.ts > heap[2*i+1].post.ts{
			child = 2*i+2
		}
		if heap[root].post.ts >= heap[child].post.ts{
			break
		}
		heap[root], heap[child] = heap[child], heap[root]
		i = child
	}
	return heap
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    if followerId == followeeId{
		return
	}
	if _,ok:=this.followeeMap[followerId];!ok{
		this.followeeMap[followerId] = map[int]struct{}{
			followeeId: struct{}{},
		}
		return
	}
	this.followeeMap[followerId][followeeId] = struct{}{}
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
     if _,ok:=this.followeeMap[followerId];!ok{
		return
	}
	delete(this.followeeMap[followerId],followeeId)
}
