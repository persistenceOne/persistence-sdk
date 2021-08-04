package test_types


import "sort"

type SortedList interface {
	sort.Interface

	Sort() SortedList

	Insert(interface{}) SortedList
	Delete(interface{}) SortedList
	Search(interface{}) int
}

