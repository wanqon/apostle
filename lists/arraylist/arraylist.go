package arraylist

import (
	"fmt"
	"strings"
)

type List struct {
	elements	[]interface{}
	size		int
}

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	list.groupBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withRange(index) {
		return nil, false
	}
	return list.elements[index], true
}

func (list *List) Remove(index int) {
	if !list.withRange(index) {
		return
	}

	list.elements[index] = nil
	copy(list.elements[index:],list.elements[index+1:list.size])
	list.size--

	list.shrink()
}

func (list *List) Contains(values ...interface{}) bool {
	for _, searchValue := range values {
		found := false
		for _, element := range list.elements {
			if element == searchValue {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list *List) Values() []interface{} {
	newElements := make([]interface{}, list.size, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

func (list *List) Indexof(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.elements {
		if element == value {
			return index
		}
	}
	return -1
}

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

func (list *List) Sort() {

}

func (list *List) Swap(i, j int) {
	if list.withRange(i) && list.withRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

func (list *List) Insert(index int, values ...interface{}) {
	if !list.withRange(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	l := len(values)
	list.groupBy(l)
	list.size += l
	copy(list.elements[index+l:],list.elements[index:list.size-l])
	copy(list.elements[index:], values)
}

func (list *List) Set(index int, value interface{}) {
	if !list.withRange(index) {
		if index == list.size {
			list.Add(value)
		}
		return
	}
	list.elements[index] = value
}

func (list *List) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

func (list *List) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *List) groupBy(n int) {
	currentCapacity := cap(list.elements)
	if list.size + n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity + n))
		list.resize(newCapacity)
	}
}

func (list *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}

func (list *List) withRange(index int) bool {
	return index >= 0 && index < list.size
}