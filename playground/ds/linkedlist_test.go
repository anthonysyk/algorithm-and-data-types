package ds

type node struct {
	value interface{}
	Next  *node
}

type LinkedList struct {
	root   *node
	length int
}

func (ll *LinkedList) Push() {

}
