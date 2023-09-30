//今天刚学的go
//今天刚学的go
//今天刚学的go

package study

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func (node *Node) String() string {
	return fmt.Sprintf("{%v, %v}", node.Val, node.Next)
}

func (node *Node) tablePrint() {
	for i := node; i != nil; i = i.Next {
		fmt.Printf("%d\t", i.Val)
	}
}

func TestReverse() {
	var head = Node{1, &Node{2, &Node{3, &Node{4, nil}}}}
	fmt.Printf("%v\n", head)
	head1 := reverse(&head)
	head1.tablePrint()
	println()
	head2 := reverse2(head1)
	head2.tablePrint()
	println()
	var l1 = Node{1, &Node{4, &Node{8, &Node{11, nil}}}}
	var l2 = Node{3, &Node{4, &Node{9, &Node{19, nil}}}}
	for i := mergeByRecursion(&l1, &l2); i != nil; i = i.Next {
		fmt.Printf("%d\t", i.Val)
	}
	//for i := mergeByFor(&l1, &l2); i != nil; i = i.Next {
	//	fmt.Printf("%d\t", i.Val)
	//}
}

func mergeByRecursion(l1, l2 *Node) *Node {
	if l1 == nil {
		return l1
	}
	if l2 == nil {
		return l2
	}
	var tempList *Node
	if l1.Val < l2.Val {
		tempList = l1
		tempList.Next = mergeByRecursion(l1.Next, l2)
	} else {
		tempList = l2
		tempList.Next = mergeByRecursion(l1, l2.Next)
	}
	return tempList
}

func mergeByFor(l1, l2 *Node) *Node {
	dummyNode := &Node{}
	cur := dummyNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return dummyNode.Next
}

func reverse(head *Node) *Node {
	var pre, cur, pos *Node = nil, head, head.Next
	//var pre, cur, pos = &Node{}, head, head.Next
	for pos != nil {
		cur.Next = pre
		pre, cur, pos = cur, pos, pos.Next
	}
	cur.Next = pre
	//head.Next = nil
	return cur
}

func reverse2(head *Node) *Node {
	var pre *Node
	cur := head
	for cur != nil {
		nextTemp := cur.Next
		cur.Next = pre
		pre, cur = cur, nextTemp
	}
	return pre
}
