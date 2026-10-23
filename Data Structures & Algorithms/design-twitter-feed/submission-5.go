
import stdheap "container/heap"

type post struct {
	id int
	ts int
}

type node struct {
	post       post
	queueIndex int
	pointer    int
}

// Max Heap
type MaxHeap []node

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].post.ts > h[j].post.ts
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(node))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type Twitter struct {
	followeeMap map[int]map[int]struct{}
	postMap     map[int][]post
	index       int
}

func Constructor() Twitter {
	return Twitter{
		followeeMap: make(map[int]map[int]struct{}),
		postMap:     make(map[int][]post),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	newPost := post{
		id: tweetId,
		ts: this.index,
	}
	this.index++

	this.postMap[userId] = append(this.postMap[userId], newPost)

	if len(this.postMap[userId]) > 10 {
		this.postMap[userId] = this.postMap[userId][1:]
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	users := []int{userId}

	for followee := range this.followeeMap[userId] {
		users = append(users, followee)
	}

	postList := [][]post{}
	maxHeap := &MaxHeap{}

	// Each user contributes their latest tweet.
	for _, user := range users {
		posts := this.postMap[user]
		if len(posts) == 0 {
			continue
		}

		queueIndex := len(postList)
		pointer := len(posts) - 1

		postList = append(postList, posts)

		*maxHeap = append(*maxHeap, node{
			post:       posts[pointer],
			queueIndex: queueIndex,
			pointer:    pointer,
		})
	}

	// O(F) heapify
	stdheap.Init(maxHeap)

	result := []int{}

	for len(result) < 10 && maxHeap.Len() > 0 {
		newest := (*maxHeap)[0]
		result = append(result, newest.post.id)

		nextPointer := newest.pointer - 1

		if nextPointer >= 0 {
			// Replace root with the next older tweet.
			(*maxHeap)[0] = node{
				post:       postList[newest.queueIndex][nextPointer],
				queueIndex: newest.queueIndex,
				pointer:    nextPointer,
			}

			// Restore heap order.
			stdheap.Fix(maxHeap, 0)
		} else {
			// No more tweets from this user.
			stdheap.Pop(maxHeap)
		}
	}

	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}

	if this.followeeMap[followerId] == nil {
		this.followeeMap[followerId] = make(map[int]struct{})
	}

	this.followeeMap[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.followeeMap[followerId], followeeId)
}
