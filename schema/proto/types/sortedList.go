package types


import "sort"

type SortedList interface {

	ProtoInterface
	sort.Interface

	Sort() SortedList

	Insert(interface{}) SortedList
	Delete(interface{}) SortedList
	Search(interface{}) int
}

