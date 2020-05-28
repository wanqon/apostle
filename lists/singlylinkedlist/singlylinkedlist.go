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


func (list *List) Sort(comparator utils.Comparator) {
	if list.size < 2 {
		return
	}
	values := list.Values()
	utils.Sort(values, comparator)
	list.Clear()
	list.Add(values)
}

func (list *List) Swap(index1, index2 int) {
	if list.withinRange(index1) && list.withinRange(index2) && index1 != index2 {
		var e1,e2 *element
		for i,currentElem := 0,list.first; e1!=nil&&e2!=nil;i,currentElem=i+1,currentElem.next{
			switch i {
			case index1:
				e1 = currentElem
			case index2:
				e2 = currentElem
			}
		}
		e1.value, e2.value = e2.value, e1.value
	}
}

func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) {
		if list.size == index {
			list.Add(values...)
		}
		return
	}
	list.size += len(values)
	var beforeElem *element
	foundElem := list.first
	for i:=0;i!=index;i,foundElem=i+1,foundElem.next {
		beforeElem = foundElem
	}

	if foundElem == list.first {
		for i, value := range values {
			newElem := &element{value: value}
			if i == 0 {
				list.first = newElem
			} else {
				beforeElem.next = newElem
			}
			beforeElem = newElem
		}
		beforeElem.next = foundElem
	} else {
		for _,v := range values {
			newElem := &element{value: v}
			beforeElem.next = newElem
			beforeElem = newElem
		}
		beforeElem.next = foundElem
	}


}

func (list *List) Set(index int, value interface{}) {
	if !list.withinRange(index) {
		if list.size == index {
			list.Add(value)
		}
		return
	}
	foundElem := list.first
	for i:=0;i!=index; {
		i, foundElem = i+1, foundElem.next
	}
	foundElem.value = value
}

func (list *List) withinRange(index int) bool {
	return 0 <= index && list.size > index
}
