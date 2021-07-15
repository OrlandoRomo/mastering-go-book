package main

import (
	"fmt"
)

type ListData struct {
	key   string
	value interface{}
}

type Hash struct {
	size int
	data []ListData
}

func initializeHash(size int) *Hash {
	return &Hash{
		data: make([]ListData, size),
		size: size,
	}
}

func (h *Hash) Set(key string, value interface{}) { // O(1)
	hash := h.hash(key)
	h.data[hash] = ListData{key, value}
	return

}

func (h *Hash) Get(key string) reflect. { // O(1)
	hash := h.hash(key)
	return h.data[hash].value
}

func (h *Hash) Print() {
	fmt.Println(h.data)
}

func (h *Hash) Keys() []string {
	keys := []string{}
	for _, value := range h.data {
		if value != (ListData{}) {
			keys = append(keys, value.key)
		}
	}
	return keys
}

func (h *Hash) Values() []interface{} {
	var values []interface{}
	for _, v := range h.data {
		if v != (ListData{}) {
			values = append(values, v.value)
		}
	}
	return values
}

// Generate hash code or index where the value will be stored or where we can access
func (h *Hash) hash(key string) int { // O(n)
	hash := 0
	for i, s := range key {
		hash += (int(s) * i) % h.size
	}
	return hash % h.size
}

func main() {
	h := initializeHash(60)

	h.Set("name", "Orlando")

	h.Set("age", 23)

	h.Set("height", 1.90)

	h.Set("weight", "70kg")

	h.Set("appels", 65)

	fmt.Println(h.Keys())

	fmt.Println(h.Values())

	height := h.Get("height")
	age := h.Get("age")
	fmt.Printf("My current height is %v\n", height)
	fmt.Printf("My current age is %v\n", age)
}
