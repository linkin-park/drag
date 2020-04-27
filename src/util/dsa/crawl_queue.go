package dsa

import "sync"

// CQueueString handles data string via FIFO manner
// and also data are unique
type CQueueString struct {
	dataLinks  []string
	internalmp map[string]bool
	mux        sync.RWMutex
}

// Has defines whether the queue has
// requested data or not
func (q *CQueueString) Has(data string) bool {
	if _, ok := q.internalmp[data]; ok {
		return true
	}
	return false
}

// Size returns number elements in the queue
func (q *CQueueString) Size() int {
	q.mux.RLock()
	defer q.mux.RUnlock()
	return len(q.dataLinks)
}

// Enqueue of data , if didnt exist early on
func (q *CQueueString) Enqueue(data ...string) {
	q.mux.Lock()
	defer q.mux.Unlock()

	for _, value := range data {
		if exist := q.Has(value); !exist {
			q.dataLinks = append(q.dataLinks, value)
			q.internalmp[value] = true

		}
	}
}

// Dequeue the data , if it exist early on
func (q *CQueueString) Dequeue() (string, bool) {
	q.mux.Lock()
	defer q.mux.Unlock()

	switch len(q.dataLinks) {
	case 0:
		return "", false
	case 1:
		firstValue := q.dataLinks[0]
		q.dataLinks = make([]string, 0, 50)
		// delete(q.internalmp, firstValue)

		return firstValue, true
	default:
		firstValue := q.dataLinks[0]
		q.dataLinks = q.dataLinks[1:]
		// delete(q.internalmp, firstValue)

		return firstValue, true
	}
}

// IntializeDefultValues to some value
func (q *CQueueString) IntializeDefultValues() {
	q.internalmp = make(map[string]bool)
}
