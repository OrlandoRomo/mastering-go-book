package main

import "fmt"

type Element struct {
	Name    string
	Surname string
	Id      string
}

type Store struct {
	store map[string]*Element
}

func NewStore() *Store {
	return &Store{
		store: make(map[string]*Element),
	}
}

func (s *Store) Add(key string, element *Element) bool {
	if key == "" {
		return false
	}
	if s.LookUp(key) == nil {
		s.store[key] = element
		return true
	}

	return false
}

func (s *Store) Delete(key string) bool {
	if s.LookUp(key) != nil {
		delete(s.store, key)
		return true
	}
	return false
}

func (s *Store) LookUp(key string) *Element {
	_, ok := s.store[key]
	if ok {
		n := s.store[key]
		return n
	}
	return nil
}

func (s *Store) Update(key string, element *Element) bool {
	s.store[key] = element
	return true
}

func (s *Store) Print() {
	for key, element := range s.store {
		fmt.Printf("key: %v value: %v\n", key, element)
	}
}
