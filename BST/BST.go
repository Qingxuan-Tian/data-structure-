package main

import "fmt"

var size int
var root treenode
type treenode struct {
	value int
	left, right *treenode
}
func BST() {
	root.value=-1
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
//向以node为节点的二分搜索树添加元素，递归写法
func addelem(node *treenode, e int)  {
	if node.value==e{
		return
	}
	if node.value >e && node.left==nil{
		node.left = creatNode(e)
		size++
		return
	} else if node.value< e && node.right==nil{
		node.right=creatNode(e)
		size++
		return
	}
	if node.value < e{
		addelem(node.right,e)
	}else {
		addelem(node.left, e)
	}

}
//向二分搜索树中插入元素
func add(e int ){
	if root.value==-1{
		root=treenode{value: e}
		size++
	} else {
		addelem(&root,e)
	}


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