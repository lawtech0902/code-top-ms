package designlinkedlist

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Val:  -1,
		Next: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	node, count := this, 0
	for node != nil {
		if count == index {
			return node.Val
		}

		node = node.Next
		count++
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.Val == -1 {
		this.Val = val
		return
	}

	node := &MyLinkedList{
		Val:  this.Val,
		Next: this.Next,
	}

	this.Val, this.Next = val, node
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := this
	for node.Next != nil {
		node = node.Next
	}

	node.Next = &MyLinkedList{
		Val:  this.Val,
		Next: nil,
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		node := &MyLinkedList{
			Val:  this.Val,
			Next: this.Next,
		}

		this.Val, this.Next = val, node
		return
	}

	count, node, prev := 0, this, this
	for node != nil {
		if count == index {
			prev.Next = &MyLinkedList{
				Val:  val,
				Next: node,
			}

			return
		}

		count++
		prev, node = node, node.Next
	}

	if count == index {
		prev.Next = &MyLinkedList{
			Val:  val,
			Next: nil,
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if this.Next != nil {
			this.Val, this.Next = this.Next.Val, this.Next.Next
		} else {
			this.Val, this.Next = -1, nil
		}

		return
	}

	count, node, prev := 0, this, this
	for node != nil {
		if count == index {
			prev.Next = node.Next
			return
		}

		count++
		prev, node = node, node.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
