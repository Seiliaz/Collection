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

func (collection *Collection) SearchByValue(data string) (bool, string) {
	for key, value := range collection.data {
		if value == data {
			return true, key
		}
	}
	return false, ""
}

func (collection *Collection) SearchByKey(key string) (bool, string) {
	_, ok := collection.data[key]
	if ok {
		return true, collection.data[key]
	}
	return false, ""
}

func main() {
	myCollection := constructor()
	myCollection.Append("hello")
	myCollection.Append("name", "seiliaz")
	status, key := myCollection.SearchByValue("seiliaz")
	fmt.Println(status, key)
	anotherStatus, value := myCollection.SearchByKey("seiliaz")
	fmt.Println(anotherStatus, value)
}
