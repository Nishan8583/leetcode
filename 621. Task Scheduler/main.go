package main

import (
	"fmt"
)

// FixedQueue represents a fixed-size queue.
type FixedQueue struct {
	data  []byte
	size  int
	front int
	rear  int
	count int
}

// NewFixedQueue initializes a new FixedQueue with the given size.
func NewFixedQueue(size int) *FixedQueue {
	return &FixedQueue{
		data: make([]byte, size),
		size: size,
	}
}

// Enqueue adds an element to the queue. Returns an error if the queue is full.
func (q *FixedQueue) Enqueue(value byte) error {
	if q.IsFull() {
		return fmt.Errorf("queue is full")
	}
	q.data[q.rear] = value
	q.rear = (q.rear + 1) % q.size
	q.count++
	return nil
}

// Dequeue removes and returns the front element of the queue. Returns an error if the queue is empty.
func (q *FixedQueue) Dequeue() error {
	if q.IsEmpty() {
		return fmt.Errorf("queue is empty")
	}
	value := q.data[q.front]
	_ = value
	q.data[q.front] = byte(0) // Clear the slot (optional)
	q.front = (q.front + 1) % q.size
	q.count--
	return nil
}

// IsFull checks if the queue is full.
func (q *FixedQueue) IsFull() bool {
	return q.count == q.size
}

// IsEmpty checks if the queue is empty.
func (q *FixedQueue) IsEmpty() bool {
	return q.count == 0
}

// Size returns the current number of elements in the queue.
func (q *FixedQueue) Size() int {
	return q.count
}

type queue struct {
	values []byte
}

func (q *FixedQueue) checkIfInQueue(value byte) bool {
	for _, v := range q.data {
		if v == value {
			return true
		}
	}

	return false
}

func (q *queue) enqueue(v byte) {
	q.values = append(q.values, v)
}

func (q *queue) dequeue() {
	if len(q.values) == 0 {
		return
	}
	if len(q.values) < 2 {
		q.values = []byte{}
	}
	q.values = q.values[1:]
}

func (q *queue) checkIfInQueue(value byte) bool {
	for _, v := range q.values {
		if v == value {
			return true
		}
	}

	return false
}

func leastInterval(tasks []byte, n int) int {
	charToCount := map[byte]int{} // holds character to its number of occurences in the tasks
	// charToInt := map[byte]int{}   // holds character to its left interval time
	q := NewFixedQueue(n)
	// we count the byte and the nubmer of occurences
	for _, v := range tasks {
		fmt.Println("Adding task ", string(v))
		charToCount[v] += 1
		// charToInt[v] = 0
	}
	fmt.Println("Done")
	sleepCount := 0 // will hold the number of interval it had to sleep
	i := 0          // when this is == len(tasks)-1, it means we completed all the tasks

	interval := 0
	for {

		if i == len(tasks)-1 { // task completed
			fmt.Println("break", i)
			break
		}

		task_completed := false
		// current_task_completed := false // by default
		for char, count := range charToCount {

			if count == 0 { // no more of this taks is left
				continue
			}
			fmt.Printf("Trying task %s\n", string(char))
			if !q.checkIfInQueue(char) {
				i += 1
				fmt.Printf("Task %s done\n", string(char))
				for j := 0; j < n; j++ {
					fmt.Println("enqueue", string(char))
					q.Enqueue(char)
				}
				charToCount[char] -= 1
				task_completed = true
				break
			}
			// if this value is addable
			/*	if charToInt[char] == 0 { // if this task is doable
					i += 1
					charToInt[char] = n    // add the internval
					charToCount[char] -= 1 // decrease the number of task
					current_task_completed = true
					// fmt.Println("Doing task ", string(char))
					break
				} else {
					charToInt[char] -= 1
					continue
				}*/

		}

		interval += 1

		if !task_completed {
			q.Dequeue()
		}
		fmt.Println("Current queue", q.data)
		fmt.Printf(
			"Interval finished %d Tasks finished %d Task completed this interval %t\n",
			interval,
			i,
			task_completed,
		)
		fmt.Println(q.data)
		// if none of the tasks were able to be completed, we need to sleep
	}

	fmt.Println(interval)
	return sleepCount
}

func main() {
	fmt.Println(leastInterval([]byte(`AAABBB`), 2))

	fmt.Println(leastInterval([]byte(`ACABDB`), 1))

	fmt.Println(leastInterval([]byte(`AAABBB`), 3))
}
