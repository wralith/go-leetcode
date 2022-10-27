package leetcode

import "sort"

type Data struct {
	value string
	time  int
}

type TimeMap map[string][]Data

func Constructor() TimeMap {
	return make(map[string][]Data)
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	_, ok := (*this)[key] // Does key already set
	if !ok {
		(*this)[key] = make([]Data, 1)
	}
	(*this)[key] = append((*this)[key], Data{
		value: value,
		time:  timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	n := (*this)[key]

	res := sort.Search(len(n), func(i int) bool {
		return n[i].time > timestamp
	})

	if res != 0 {
		return n[res-1].value
	}
	return ""
}

// -- Binary Search -- Get --
// func (this *TimeMap) Get(key string, timestamp int) string {
// 	n := (*this)[key]
// 	l, r := 0, len(n)

// 	for l < r {
// 		mid := l + (r-l)/2
// 		if n[mid].time < timestamp {
// 			l = mid + 1
// 		} else {
// 			r = mid
// 		}
// 	}

// 	if l >= 0 && l < len(n) && n[l].time == timestamp {
// 		return n[l].value
// 	}

// 	if l > 0 {
// 		return n[l-1].value
// 	}

// 	return ""
// }

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
