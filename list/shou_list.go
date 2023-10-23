package main

import "fmt"

// 定义一个元素
type ListNode struct {
	value int
	Prev  *ListNode
	Next  *ListNode
}

// 定义一个双向链表，只记录头和尾即可
type DoubleList struct {
	Head   *ListNode
	Tail   *ListNode
	length int
}

// 最后面插入元素
func (list *DoubleList) Append(x int) {
	node := &ListNode{ //将这个插入的元素进行元素化
		value: x,
		Prev:  nil,
		Next:  nil,
	}
	tail := list.Tail
	if tail != nil { //说明有末尾元素
		tail.Next = node
		node.Prev = tail
		list.Tail = node
	} else { //否则说明该链表为空
		list.Head = node
		list.Tail = node
	}
	list.length += 1
}

// 获取某个元素,首元素为第0个
func (list *DoubleList) Get(x int) *ListNode {
	if list.length <= x {
		return nil
	}
	curr := list.Head
	for i := 0; i < x; i++ {
		curr = curr.Next
	}
	return curr
}

// 插入某个元素的后面
func (list *DoubleList) InsertAfter(prevNode *ListNode, x int) {
	node := &ListNode{
		value: x,
		Prev:  nil,
		Next:  nil,
	}
	if prevNode.Next == nil {
		prevNode.Next = node
		node.Prev = prevNode
	} else {
		nextNode := prevNode.Next
		nextNode.Prev = node
		node.Next = nextNode
		prevNode.Next = node
		node.Prev = prevNode
	}
	list.length += 1
}

// 遍历
func (list *DoubleList) Traverse() {
	head := list.Head //取头元素
	for head != nil {
		fmt.Printf("%v ", head.value)
		head = head.Next
	}
	fmt.Println() //最后换行
}

func testDoubleList() {
	list := new(DoubleList)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Traverse() //1,2,3,4,5
	// node := list.Get(2)
	// fmt.Println(node.value)          //3
	fmt.Println(list.Get(2).value)   //3
	list.InsertAfter(list.Get(2), 9) //1,2,3,9,4,5
	list.Traverse()
}

func main() {
	testDoubleList()
}
