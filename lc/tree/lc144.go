package tree

//树的结构体
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func PreOrderTraversal(treeNode *Node) []int {
	r := make([]int, 0)
	inner(treeNode, &r)
	return r
}

func inner(treeNode *Node, array *[]int) {
	if treeNode == nil {
		return
	}
	*array = append(*array, treeNode.Val)
	inner(treeNode.Left, array)
	inner(treeNode.Right, array)
}
