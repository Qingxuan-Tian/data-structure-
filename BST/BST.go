package BST
var size int
var root treenode
type treenode struct {
	value int
	left, right *treenode
}
func BST() {
	size=0
}
func getSize(size int)int  {
	return size
}
func isEmpty() bool  {
	return size==0
}
