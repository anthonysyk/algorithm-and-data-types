package store

import (
	"errors"
	"fmt"
	"os"
	"testing"

	cg "contacts/contactgraph"
)

func TestLoadData(t *testing.T) {
	nbOfNodes := 100000
	nbOfContactsPerNodes := 50
	filename := generateGraphName(nbOfNodes, nbOfContactsPerNodes)
	// open the store
	if _, err := os.Stat(filename); !errors.Is(err, os.ErrNotExist) {
		fmt.Println("graph already exists")
		return
	}
	store, err := Open(filename)
	checkErr(err)

	graph := cg.NewContactGraph()
	allPhones := cg.PopulateRandom(nbOfNodes, nbOfContactsPerNodes, graph)

	user1 := cg.User{PhoneNumber: allPhones[0]}

	// put: encodes value with gob and updates the boltdb
	err = store.Put("graph", graph)
	checkErr(err)

	err = store.Put("user1", user1)
	checkErr(err)

	// close the store
	store.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
