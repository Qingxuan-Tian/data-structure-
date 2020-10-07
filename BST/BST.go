package main

import "fmt"
//二分搜索树初始化（6.2）
var size int
var root *treenode= new(treenode)
type treenode struct {
	value int
	left, right *treenode
}
func BST() {
	size=0
	return
}
func creatNode(e int)*treenode  {
	return &treenode{value: e}
}
func getSize()int  {
	return size
}
func isEmpty() bool  {
	return size==0
}
//向以node为节点的二分搜索树添加元素，递归写法(6.3 6.4)
func addelem(node *treenode, e int) *treenode {
	//if node.value==e{
	//	return
	//}
	//if node.value >e && node.left==nil{
	//	node.left = creatNode(e)
	//	size++
	//	return
	//} else if node.value< e && node.right==nil{
	//	node.right=creatNode(e)
	//	size++
	//	return
	//}
	//if node.value < e{
	//	addelem(node.right,e)
	//}else if node.value>e{
	//	addelem(node.left, e)
	//}
	//注释部分改进写法
	if node==nil{
		size++
		return &treenode{value: e}
	}
	if node.value < e{
		node.right=addelem(node.right,e)
	}else if node.value>e{
		node.left=addelem(node.left, e)
	}
	return node

}
//向二分搜索树中插入元素（6.3 6.4）
func add(e int ){
	//if root==nil{
	//	root=&treenode{value: e}
	//	size++
	//} else {
	//	root = addelem(root,e)
	//}
	//改进写法
	root = addelem(root,e)
}

//查询二分搜索树中的元素（6.5）
func contain(root *treenode,e int)bool  {
	if root==nil{
		return false
	}
	if root.value==e{
		return true
	}else if root.value<e{
		return contain(root.right,e)
	}else {
		return contain(root.left, e)
	}
}
//前序遍历(6.6)
func preOrder(root *treenode) {
	if root==nil {
		return
	}
	fmt.Println(root.value)
	fmt.Println()
	preOrder(root.left)
	preOrder(root.right)

}
//中序遍历,输出结果为升序排序结果(6.7)
func inorder(root *treenode){
	if root==nil{
		return
	}
	inorder(root.left)
	fmt.Println(root.value)
	inorder(root.right)
}
//后序遍历,应用场景：二分搜索树释放内存空间(6.7)
func postOrder(root *treenode)  {
	if root==nil{
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Println(root.value)
}
//前序遍历的非递归写法(6.9)
func preOrderNR(root *treenode){
	return
}

func main()  {
	BST()
	add(2)
	add(5)
	add(10)
	fmt.Println(root)
	fmt.Println(root.right)
	fmt.Println(root.right.right)
	fmt.Println(getSize())
}