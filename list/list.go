package list

type Node struct {
	Val  int
	Next *Node
}

func Reverse(head *Node) *Node {
	// 定义一个前驱节点
	var pre *Node
	// 当头节点不为空时循环
	for head != nil {
		// 临时保存头节点的下一个节点
		next := head.Next
		// 将头节点的下一个节点指向前驱节点
		head.Next = pre
		// 将前驱节点更新为头节点
		pre = head
		// 将头节点更新为临时保存的下一个节点
		head = next
	}
	// 返回反转后的头节点
	return pre
}

// 递归形式解决该问题
func Reverse2(head *Node) *Node {
	// 如果链表为空或只有一个节点，则直接返回原链表
	if head == nil || head.Next == nil {
		return head
	}
	// 递归调用Reverse2函数，传入下一个节点作为参数，得到反转后的新链表头节点
	newHead := Reverse2(head.Next)

	// 将当前节点的下一个节点的next指向当前节点，实现反转
	head.Next.Next = head
	// 将当前节点的next置为nil，避免链表循环
	head.Next = nil

	// 返回反转后的新链表头节点
	return newHead
}

func Reverse3(head *Node, left, right int) *Node {
	dummyNode := &Node{
		Val:  -1,
		Next: head,
	}
	var leftPre *Node = dummyNode
	var leftNode *Node
	var rightNext *Node
	var rightNode *Node
	for i := 0; i < left-1; i++ {
		leftPre = leftPre.Next
	}

	leftNode = leftPre.Next
	rightNode = leftNode

	for i := 0; i < right-left; i++ {
		rightNode = rightNode.Next
	}
	rightNext = rightNode.Next
	rightNode.Next = nil
	leftPre.Next = nil
	Reverse(leftNode)

	leftPre.Next = rightNode
	leftNode.Next = rightNext

	return dummyNode.Next

}

// 使用头插法翻转部分链表
func reverseBetween(head *Node, left int, right int) *Node {
	dummyNode := &Node{
		Val:  -1,
		Next: head,
	}
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	curr := pre.Next
	for i := 0; i < right-left; i++ {
		next := curr.Next
		curr.Next = next.Next

		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}

func mergeList(head1 *Node, head2 *Node) *Node {
	dummyNode := &Node{
		Val:  -1,
		Next: nil,
	}
	curr, temp1, temp2 := dummyNode, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val < temp2.Val {
			curr.Next = temp1
			temp1 = temp1.Next
		} else {
			curr.Next = temp2
			temp2 = temp2.Next
		}
	}
	if temp1 != nil {
		curr.Next = temp1
	}
	if temp2 != nil {
		curr.Next = temp2
	}
	return dummyNode.Next
}
func sortList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	temp := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(temp)
	return mergeList(left, right)
}
func mergeKList(lists []*Node) *Node {
	if len(lists) == 0 {
		return nil
	}
	var ans *Node
	for i := 1; i < len(lists); i++ {
		ans = mergeList(ans, lists[i])
	}
	return ans
}

func rotateRight(head *Node, k int) *Node {
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}
	length := 1
	curr := head
	for curr.Next != nil {
		length++
		curr = curr.Next
	}
	k = k % length
	curr.Next = head
	for i := 0; i < length-k; i++ {
		curr = curr.Next
	}
	ret := curr.Next
	curr.Next = nil
	return ret
}

func partition(head *Node, x int) *Node {
	dummy1 := &Node{}
	dummy2 := &Node{}
	curr1 := dummy1
	curr2 := dummy2
	curr := head
	for curr != nil {
		if curr.Val <= x {
			curr1.Next = curr
			curr1 = curr
		} else {
			curr2.Next = curr
			curr2 = curr
		}
		curr = curr.Next
	}
	curr2.Next = nil
	curr1.Next = dummy2.Next
	return dummy1.Next
}

func swapPairs(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}

func swapPairs2(head *Node) *Node {
	dummy := &Node{}
	dummy.Next = head
	temp := dummy
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next

		node1.Next = node2.Next
		temp.Next = node2
		node2.Next = node1

		temp = node1
	}
	return dummy.Next
}

func hasCycle(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}

		if fast.Val == slow.Val {
			return true
		}
	}
	return false
}

func MeetPoint(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			newNode := head
			for slow != newNode {
				newNode = newNode.Next
				slow = slow.Next
			}
			return newNode
		}
	}
	return nil
}

func MeetPoint2(headA, headB *Node) *Node {
	currA := headA
	currB := headB
	for currA != currB {
		if currA == nil {
			currA = headB
		} else {
			currA = currA.Next
		}
		if currB == nil {
			currB = headA
		} else {
			currB = currB.Next
		}
	}
	return currA

}

// IsPalindrome 判断一个链表是否为回文链表
func IsPalindrome(head *Node) bool {
	if head == nil {
		return true
	}

	// 快慢指针找到中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 反转后半部分链表
	var prev *Node
	for fast != nil {
		next := fast.Next
		fast.Next = prev
		prev = fast
		fast = next
	}

	// 比较前半和反转后半是否相同
	for prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}

	return true
}

func reverseList(head *Node) *Node {
	var prev, cur *Node = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func endOfFirstHalf(head *Node) *Node {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func isPalindrome(head *Node) bool {
	if head == nil {
		return true
	}

	// 找到前半部分链表的尾节点并反转后半部分链表
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}
