package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func InOrder(node *Node) []int {
	var ans []int
	inorderTraversal(node, ans)
	return ans
}

func inorderTraversal(node *Node, ans []int) []int {
	if node == nil {
		return ans
	}
	in(node.Left, ans)
	ans = append(ans, node.Val)
	in(node.Right, ans)
	return ans
}

func NumTrees(n int) int {
	if n == 0 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}

func LevelOrder(n *Node) (ans []int) {
	if n == nil {
		return []int{}
	}
	q := []*Node{n}
	for len(q) > 0 {
		next := []*Node{}
		for _, node := range q {
			ans = append(ans, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		q = next
	}
	return ans
}

// binary tree 二叉树
// binary search tree 二叉搜索树

func ArrayToBinarySearchTree(nums []int) *Node {
	return arrayToBinarySearchTree(nums, 0, len(nums)-1)
}

func arrayToBinarySearchTree(nums []int, left, right int) *Node {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	node := &Node{Val: nums[mid]}
	node.Left = arrayToBinarySearchTree(nums, left, mid-1)
	node.Right = arrayToBinarySearchTree(nums, mid+1, right)
	return node
}

func minDepth(node *Node) int {
	if node == nil {
		return 0
	}
	minLeft := minDepth(node.Left)
	minRight := minDepth(node.Right)
	if minLeft < minRight {
		return minLeft + 1
	} else {
		return minRight + 1
	}
}

func GetCommonRoot(root, p, q *Node) *Node {
	if root == nil {
		return nil
	}
	curr := root
	for curr != nil {
		if p.Val < curr.Val && q.Val < curr.Val {
			curr = curr.Left
		} else if p.Val > curr.Val && q.Val > curr.Val {
			curr = curr.Right
		} else {
			return curr
		}
	}
	return curr
}
