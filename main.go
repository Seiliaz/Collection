package main

import "fmt"

type Collection struct {
	data        map[string]string
	latestIndex uint
}

func constructor() Collection {
	return Collection{latestIndex: 0}
}

func (collection Collection) Append(data ...string) {
	var key string
	var value string
	if len(data) == 1 {
		key = fmt.Sprint(collection.latestIndex)
		value = data[0]
	} else if len(data) >= 2 {
		key = data[0]
		value = data[1]
	}
	val, ok := collection.data[key]
	if ok {
		fmt.Println("the key is already taken and the value is", val)
	} else {
		fmt.Println("there we are!")
		// collection.data[key] = value
	}
}

func main() {
	myCollection := constructor()
	myCollection.Append("hello")
	fmt.Println(myCollection.data)
}
