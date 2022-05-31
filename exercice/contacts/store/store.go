package store

import (
	cg "contacts/contactgraph"
	"fmt"
)

func generateGraphName(nbOfNodes, nbOfContactsPerNodes int) string {
	return fmt.Sprintf("graph-%v-%v.db", nbOfNodes, nbOfContactsPerNodes)
}

func LoadGraph(nbOfNodes, nbOfContactsPerNodes int) (cg.ContactGraph, cg.User) {
	graph := cg.NewContactGraph()

	var user1 cg.User

	// fetch data from local storage boltdb
	filename := generateGraphName(nbOfNodes, nbOfContactsPerNodes)
	st, err := Open(fmt.Sprintf("./store/%s", filename))
	if err != nil {
		panic(err)
	}
	err = st.Get("graph", &graph)
	if err != nil {
		panic(err)
	}
	err = st.Get("user1", &user1)
	if err != nil {
		panic(err)
	}

	return graph, user1
}
