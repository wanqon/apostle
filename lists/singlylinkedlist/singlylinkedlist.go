package singlylinkedlist

import (
	"github.com/wanqon/apostle/lists"
	"github.com/wanqon/apostle/utils"
)

type List struct {
	first	*element
	last	*element
	size	int
}

func (list *List) Sort(comparator utils.Comparator) {
	panic("implement me")
}

func (list *List) Swap(index1, index2 int) {
	panic("implement me")
}

func (list *List) Insert(index int, values ...interface{}) {
	panic("implement me")
}

func (list *List) Set(index int, value interface{}) {
	panic("implement me")
}

type element struct {
	value	interface{}
	next	*element
}

func assertListImplementation() {
	var _ lists.List = (*List)(nil)
}

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	for _, value := range values{
		newElement := &element{value: value}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

func (list *List) Append(values ...interface{}) {
	list.Add(values)
}

func (list *List) Prepend(values ...interface{}) {
	for v := len(values); v > 0; v-- {
		newElement := &element{value:values[v], next:list.first}
		list.first = newElement
		if list.size == 0 {
			list.last = newElement
		}
		list.size++
	}
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}

	element := list.first
	for i:=0; i != index; i, element = i+1, element.next {
	}
	return element.value, true
}

func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}
	
	var beforeElement *element
	element := list.first
	for i:=0; i != index; i, element = i+1, element.next {
		beforeElement = element
	}
	if element == list.first {
		list.first = element.next
	}
	if element == list.last {
		list.last = beforeElement
	}
	if beforeElement != nil {
		beforeElement.next = element.next
	}
	element = nil
	list.size--
}

func (list *List) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}

	if list.size == 0 {
		return false
	}

	for _, value := range values {
		found := false
		for element := list.first; element != nil; element = element.next {
			if element.value == value {
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

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Values() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for i,element := 0, list.first; element != nil; i, element = i+1, element.next {
		values[i] = element
	}
	return values
}

func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

func (list *List) withinRange(index int) bool {
	return 0 <= index && list.size > index
}
