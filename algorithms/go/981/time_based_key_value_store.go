package main

type timeValue struct {
	timestamp int
	value     string
}

type timeValueList []*timeValue

func (tv timeValueList) search(timestamp int) string {
	if tv == nil || len(tv) == 0 {
		return ""
	}
	l, r := 0, len(tv)
	for l < r {
		m := (l + r) / 2
		if tv[m].timestamp <= timestamp {
			l = m + 1
		} else {
			r = m
		}
	}
	if l > 0 {
		return tv[l-1].value
	}
	return ""
}

type TimeMap struct {
	m map[string]timeValueList
}

func Constructor() TimeMap {
	return TimeMap{
		m: map[string]timeValueList{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], &timeValue{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	return this.m[key].search(timestamp)
}
