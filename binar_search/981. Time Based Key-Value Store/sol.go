package main

import "fmt"

type node struct {
	timestamp *int // will hold the timestamp number
	nodeValue string
	key       string // will hold the value in that timestamp
	left      *node  // points to the next node with less value
	right     *node  // points right with the higher value
}

func (n *node) insert(key, value string, timestamp int) {

	if n.timestamp == nil {
		n.timestamp = &timestamp
		n.key = key
		n.nodeValue = value
		return
	}

	// if the new value is less than the current value, put it on the left
	if timestamp < *n.timestamp {
		if n.left == nil {
			n.left = &node{
				timestamp: &timestamp,
				key:       key,
				nodeValue: value,
			}
			return
		}

		// if there is a value, then do a comparision with the next vlaue
		n.left.insert(key, value, timestamp)
		return
	} else {
		if n.right == nil {
			n.right = &node{
				timestamp: &timestamp,
				key:       key,
				nodeValue: value,
			}
			return
		}
		n.right.insert(key, value, timestamp)
		return
	}
	return

}

func (n *node) get(key string, timestamp int) (string, int) {

	// if the current node value is equals to timestamp return the value
	// if the value is not equals, check if next left node is empty
	// because we are trying to pass the left one

	val := ""
	if n.key == key && timestamp == *n.timestamp {
		val = *&n.nodeValue
		return val, *n.timestamp
	}

	if *&n.key == key && *n.timestamp <= timestamp {
		val = *&n.nodeValue
	}

	if timestamp > *n.timestamp {
		val = *&n.nodeValue
		if n.right != nil {
			local_val := ""
			local_val, t := n.right.get(key, timestamp)
			if local_val != "" && t > *n.timestamp {
				val = local_val
			}
		}
	} else if timestamp < *n.timestamp {
		if n.left != nil {
			local_val := ""
			local_val, t := n.left.get(key, timestamp)
			if local_val != "" && t > *n.timestamp {
				val = local_val
			}
		}

	}

	fmt.Println("current", val, *n.timestamp)

	return val, *n.timestamp

}

type TimeMap struct {
	n *node
}

func Constructor() TimeMap {
	return TimeMap{
		n: &node{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.n.insert(key, value, timestamp)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	s, _ := this.n.get(key, timestamp)
	return s
}

func main() {
	//obj := Constructor()
	/*obj.Set("foo", "bar", 1)
	fmt.Println("for 1", obj.Get("foo", 1))
	fmt.Println("for 2", obj.Get("foo", 2))
	obj.Set("foo", "bar2", 4)
	fmt.Println("for 4", obj.Get("foo", 4))
	fmt.Println("for 5s", obj.Get("foo", 5))*/

	/*obj.Set("love", "high", 10)
	obj.Set("love", "low", 20)
	fmt.Println("love 5----------", obj.Get("love", 5))

	fmt.Println("love 10---------", obj.Get("love", 10))
	fmt.Println("love 15---------", obj.Get("love", 15))
	fmt.Println("staring for 20")
	fmt.Println("love 20----------", obj.Get("love", 20))
	fmt.Println("love 25----------", obj.Get("love", 25))*/

	obj2 := Constructor()
	obj2.Set("ljfvbut", "tatthnvvid", 3)
	obj2.Get("ljfvbut", 4)
	obj2.Get("ljfvbut", 5)
	obj2.Get("ljfvbut", 6)
	obj2.Get("ljfvbut", 7)
	obj2.Set("eimdon", "pdjbnnvje", 8)
	fmt.Println("FOR 9")
	obj2.Get("eimdon", 9)
	obj2.Get("eimdon", 10)
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
