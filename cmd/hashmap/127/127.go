package _127

type MapSet map[string]struct{}

func NewMapSet(list []string) *MapSet {
	set := make(MapSet, len(list))
	for _, item := range list {
		set.Set(item)
	}
	return &set
}

func (m *MapSet) Has(key string) bool {
	_, ok := (*m)[key]
	return ok
}
func (m *MapSet) Set(key string) {
	(*m)[key] = struct{}{}
}

type Queue []string

func NewQueue() *Queue {
	queue := make(Queue, 0)
	return &queue
}
func (q *Queue) Push(value string) {
	*q = append(*q, value)
}

func (q *Queue) Ppp() string {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	set := NewMapSet(wordList)
	if !set.Has(endWord) {
		return 0
	}

	queue := NewQueue()
	queue.Push(beginWord)
	queue.Push(beginWord)
	visited := &MapSet{}
	visited.Set(beginWord)
	changes := 1
	for !queue.Empty() {
		size := queue.Size()
		for i := 0; i < size; i++ {
			word := queue.Ppp()
			if word == endWord {
				return changes
			}
			for j := 0; j < len(word); j++ {
				for k := 'a'; k <= 'z'; k++ {
					bytes := []byte(word)
					bytes[j] = byte(k)
					newWord := string(bytes)
					if set.Has(newWord) && !visited.Has(newWord) {
						queue.Push(newWord)
						visited.Set(newWord)
					}
				}
			}
		}
		changes++
	}
	return 0
}
