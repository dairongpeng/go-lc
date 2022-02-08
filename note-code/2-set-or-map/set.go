package main

type Set interface {
	Add(elements ...interface{})
	Remove(elements ...interface{})
	Contains(elements ...interface{}) bool
}

var itemExists = struct{}{}

type HashSet struct {
	items map[interface{}]struct{}
}

func New(values ...interface{}) *HashSet {
	set := &HashSet{items: make(map[interface{}]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

func (set *HashSet) Add(items ...interface{}) {
	for _, item := range items {
		set.items[item] = itemExists
	}
}

func (set *HashSet) Remove(items ...interface{}) {
	for _, item := range items {
		delete(set.items, item)
	}
}

func (set *HashSet) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}
