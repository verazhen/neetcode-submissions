type task struct{
	id byte
	freq int
	nextIndex int
}
type maxHeap []task

func leastInterval(tasks []byte, n int) int {
	
	// 1. initialize
	// check every tasks's frequency
	// initialize a max tasks's frequency heap for avaiavle task only
	freqMap := map[byte]int{}
	maxFreqAvailableHeap := maxHeap{}
	for _, id:=range tasks{
		freqMap[id]++
	}
	for id,freq:=range freqMap{
		maxFreqAvailableHeap = append(maxFreqAvailableHeap, task{id:id,freq:freq})
	}
	//sift down from the last non-leaf node to top
	for i := len(maxFreqAvailableHeap)/2-1;i >= 0;i--{
		maxFreqAvailableHeap.siftDown(i)
	}

	// 2. update task queue to avaiavle task heap
	// sift up the task heap and get the root
		//remove root and sift
		//deduct the remain frequemcy in map
		//if remain frequemcy >0, append back to the queue
	// if heap is empty, get latest task in queue and take idle as option
	// if heap is empty, and the queue is empty, break
	waitQueue := []task{}
	nowIndex := 0
	for len(maxFreqAvailableHeap)>0 ||len(waitQueue)>0{
		//update maxHeap
		for len(waitQueue) > 0{
			if waitQueue[0].nextIndex > nowIndex{
				break
			}
			pop := waitQueue[0]
			waitQueue = waitQueue[1:]
			maxFreqAvailableHeap = append(maxFreqAvailableHeap,pop)
			maxFreqAvailableHeap.siftUp(len(maxFreqAvailableHeap)-1)
		}
		if len(maxFreqAvailableHeap) > 0{
			do := maxFreqAvailableHeap[0]
			l := len(maxFreqAvailableHeap)
			maxFreqAvailableHeap[0] = maxFreqAvailableHeap[l-1]
			maxFreqAvailableHeap = maxFreqAvailableHeap[:l-1]
			if len(maxFreqAvailableHeap) > 0{
				maxFreqAvailableHeap.siftDown(0)
			}
			
			do.freq--
			if do.freq > 0{
				do.nextIndex = nowIndex + n + 1
				waitQueue = append(waitQueue,do)
			}
			nowIndex++
		}else if len(waitQueue)>0{
			nowIndex = waitQueue[0].nextIndex
		}
	}

	return nowIndex
}

func(h *maxHeap) siftDown(i int){
	parent := i
	for{
		child := parent*2+1
		if child >= len(*h){
			break
		}
		if child + 1 < len(*h) && (*h)[child].freq<(*h)[child+1].freq{
			child = child +1
		}
		if (*h)[parent].freq >= (*h)[child].freq{
			break
		}
		(*h)[parent],(*h)[child] = (*h)[child], (*h)[parent]
		parent = child
	}
}

func(h *maxHeap) siftUp(i int){
	child := i
	for child > 0{
		parent := (child-1)/2
		if (*h)[parent].freq >= (*h)[child].freq{
			break
		}
		(*h)[parent],(*h)[child] = (*h)[child], (*h)[parent]
		child = parent
	}
}
