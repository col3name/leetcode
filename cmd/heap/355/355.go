package _355

import (
	"container/heap"
)

type Tweet struct {
	Id          int
	PublishedAt int
}

type MaxHeap []Tweet

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) First() interface{} {
	return h[0]
}

func (h MaxHeap) Less(i int, j int) bool {
	return h[i].PublishedAt > h[j].PublishedAt
}

func (h MaxHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Pop() interface{} {
	a := *h
	res := a[len(a)-1]
	*h = a[:len(a)-1]
	return res
}

func (h *MaxHeap) Push(i interface{}) {
	*h = append(*h, i.(Tweet))
}

type MapSet map[int]struct{}

func (m *MapSet) Has(key int) bool {
	_, ok := (*m)[key]
	return ok
}
func (m *MapSet) Set(key int) {
	(*m)[key] = struct{}{}
}

type Twitter struct {
	tweets  map[int]*MaxHeap
	follows map[int]MapSet
	id      int
}

func Constructor() Twitter {
	return Twitter{
		tweets:  make(map[int]*MaxHeap, 0),
		follows: make(map[int]MapSet),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.tweets[userId]; !ok {
		pq := &MaxHeap{}
		heap.Init(pq)
		this.tweets[userId] = pq
	}
	if _, ok := this.follows[userId]; !ok {
		this.follows[userId] = make(MapSet, 0)
	}
	this.follows[userId][userId] = struct{}{}
	this.id++
	heap.Push(this.tweets[userId], Tweet{
		Id:          tweetId,
		PublishedAt: this.id,
	})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	followSet, ok := this.follows[userId]
	if !ok {
		return []int{}
	}
	pq := &MaxHeap{}
	heap.Init(pq)

	seen := MapSet{}
	for user, _ := range followSet {
		userTweets, ok := this.tweets[user]
		if !ok {
			continue
		}
		var tmp []*Tweet
		length := userTweets.Len()
		for i := 0; i < 10 && i < length; i++ {
			tweet := heap.Pop(userTweets).(Tweet)
			if !seen.Has(tweet.Id) {
				heap.Push(pq, tweet)
				seen.Set(tweet.Id)
			}
			tmp = append(tmp, &tweet)
		}
		for _, tweet := range tmp {
			heap.Push(this.tweets[user], *tweet)
		}
	}
	var result []int
	length := pq.Len()
	for i := 0; i < 10 && i < length; i++ {
		result = append(result, heap.Pop(pq).(Tweet).Id)
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	_, ok := this.follows[followerId]
	if !ok {
		this.follows[followerId] = make(MapSet, 0)
	}

	set := this.follows[followerId]
	if !set.Has(followeeId) {
		this.follows[followerId][followeeId] = struct{}{}
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.follows[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
