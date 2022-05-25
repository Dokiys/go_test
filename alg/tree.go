package alg

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrderTraversal(p *TreeNode, res *[]int) {
	if p != nil {
		*res = append(*res, p.Val)
		PreOrderTraversal(p.Left, res)
		PreOrderTraversal(p.Right, res)
	}
}

func InOrderTraversal(p *TreeNode, res *[]int) {
	if p != nil {
		InOrderTraversal(p.Left, res)
		*res = append(*res, p.Val)
		InOrderTraversal(p.Right, res)
	}
}

func PostOrderTraversal(p *TreeNode, res *[]int) {
	if p != nil {
		PostOrderTraversal(p.Left, res)
		PostOrderTraversal(p.Right, res)
		*res = append(*res, p.Val)
	}
}