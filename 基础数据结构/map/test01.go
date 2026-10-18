package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	fmt.Println(len(nums))
	m1 := make(map[string]int, 8)
	fmt.Println(m1)
	m1["key1"] = 1
	m1["key2"] = 2
	m1["key3"] = 3
	m1["key4"] = 4
	m1["key5"] = 5
	if v, ok := m1["key"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("key does not exist")
	}
	fmt.Println("--------------")
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	a := m1["key1"]
	b, ok := m1["key2"]
	fmt.Println(a, b, ok)
	c, ok := m1["key3"]
	fmt.Println(c, ok)
	delete(m1, "key1")
	fmt.Println(m1)
}

/*
4
map[]
key does not exist
--------------
key1 1
key2 2
key3 3
key4 4
key5 5
1 2 true
3 true
map[key2:2 key3:3 key4:4 key5:5]
*/
