package main

import "fmt"

type Collection struct {
	data        map[string]string
	latestIndex uint
}

func constructor() Collection {
	return Collection{latestIndex: 0, data: make(map[string]string)}
}

func (collection *Collection) Append(data ...string) {
	var key string
	var value string
	if len(data) == 1 {
		key = fmt.Sprint(collection.latestIndex)
		value = data[0]
		collection.latestIndex++
	} else if len(data) >= 2 {
		key = data[0]
		value = data[1]
	}
	val, ok := collection.data[key]
	if ok {
		fmt.Println("the key is already taken and the value is", val)
	} else {
		collection.data[key] = value
	}
}

func (collection *Collection) SearchByValue(val string) (bool, string) {
	for key, value := range collection.data {
		if value == val {
			return true, key
		}
	}
	return false, ""
}

func (collection *Collection) SearchByKey(key string) (bool, string) {
	_, ok := collection.data[key]
	var data string
	if data = ""; ok {
		data = collection.data[key]
	}
	return ok, data
}

func (collection *Collection) _remove(key string) {
	delete(collection.data, key)
}

func (collection *Collection) RemoveByValue(val string) {
	for key, value := range collection.data {
		if value == val {
			collection._remove(key)
		}
	}
}

func (collection *Collection) RemoveByKey(key string) {
	for field := range collection.data {
		if field == key {
			collection._remove(key)
		}
	}
}

func (collection *Collection) _update(key, after string) {
	collection.data[key] = after
}

func (collection *Collection) UpdateEveryValue(before, after string) {
	for key, value := range collection.data {
		if value == before {
			collection._update(key, after)
		}
	}
}

func (collection *Collection) UpdateByKey(key string, after string) {
	for field := range collection.data {
		if field == key {
			collection._update(key, after)
		}
	}
}

func (collection *Collection) First() (bool, string) {
	if len(collection.data) > 0 {
		for _, value := range collection.data {
			return true, value
		}
	}
	return false, ""
}

func main() {
	myCollection := constructor()
	myCollection.Append("hello")
	myCollection.Append("name", "seiliaz")
	myCollection.Append("another_name", "seiliaz")
	fmt.Println(myCollection.data)
}
