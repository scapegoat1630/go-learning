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

