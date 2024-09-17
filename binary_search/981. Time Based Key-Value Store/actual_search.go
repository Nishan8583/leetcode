/*
lettcode link: https://leetcode.com/problems/time-based-key-value-store/submissions/1239973401/
solution: https://youtu.be/fu2cD_6E8Hw?si=hzBLIPp0VL1UG-Ix

Map of key to struct of vlaues and timestamp
question already states that the timestamp are going to be in ascending order
Do a binary search on the list of values
*/

package main

import "fmt"

type listVals struct {
	value     string
	timestamp int
}
type TimeMap struct {
	values map[string][]listVals
}

func Constructor() TimeMap {
	return TimeMap{
		values: make(map[string][]listVals),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if v, ok := this.values[key]; ok {
		v = append(v, listVals{value: value, timestamp: timestamp})
		this.values[key] = v
	} else {
		this.values[key] = []listVals{
			listVals{value: value, timestamp: timestamp},
		}
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	res := ""
	values, ok := this.values[key]
	if !ok || len(values) == 0 {
		return res
	}
	l := 0
	r := len(values) - 1

	// this is binary search
	for l <= r {
		//fmt.Println(l, r)
		m := int(r + l/2) // get the mid value
		//fmt.Printf(" left = %d mid =%d r=%d values=%v vlaue timestamp= %d timestamp=%d \n", l, m, r, values, values[m].timestamp, timestamp)
		if timestamp == values[m].timestamp {
			return values[m].value
		} else if values[m].timestamp < timestamp {
			res = values[m].value
			l = m + 1
			//fmt.Println("low value new l", res, m)
		} else {
			r = m - 1
		}
	}
	return res
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func main() {
	obj := Constructor()
	obj.Set("foo", "bar", 1)
	fmt.Println("for 1", obj.Get("foo", 1))
	fmt.Println("for 2", obj.Get("foo", 2))
	obj.Set("foo", "bar2", 4)
	fmt.Println("Set for 4")
	fmt.Println("Getting for 4")
	fmt.Println("for 4", obj.Get("foo", 4))
	fmt.Println("for 5s", obj.Get("foo", 5))

}
