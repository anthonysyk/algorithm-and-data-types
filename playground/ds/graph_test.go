package ds

import "testing"

// directed acyclic unweighted graph
type graph struct {
	Length       int
	AdjacentList map[interface{}][]interface{}
}

func NewGraph() graph {
	return graph{
		Length:       0,
		AdjacentList: make(map[interface{}][]interface{}),
	}
}

func (g *graph) AddNode(name interface{}) {

}
func (g *graph) RemoveNode(name interface{}) {

}
func (g *graph) AddEdge(node1, node2 interface{}) {

}
func (g *graph) RemoveEdge(node1, node2 interface{}) {

}

func (g *graph) Show() map[interface{}][]interface{} {
	return g.AdjacentList
}

func TestGraph(t *testing.T) {

}
