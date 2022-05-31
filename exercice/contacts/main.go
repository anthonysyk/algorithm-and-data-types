package contacts

import (
	"contacts/store"
	"fmt"
	"time"

	cg "contacts/contactgraph"
)

func main() {
	graph, user1 := store.LoadGraph(100000, 50)

	// Test Lookup
	users, duration := WithDurationLookup(graph.Lookup, user1.PhoneNumber)
	fmt.Printf("Lookup returned %v users \n", len(users))
	fmt.Printf("Lookup took : %s \n", duration)

	// Test RLookup
	users, duration = WithDurationLookup(graph.RLookup, user1.PhoneNumber)
	fmt.Printf("RLookup returned %v users \n", len(users))
	fmt.Printf("RLookup took : %s \n", duration)

	// Test Suggest
	suggestions, duration := WithDurationSuggestion(graph.Suggest, user1)
	fmt.Printf("Suggestions:\n %v\n", suggestions)
	fmt.Printf("Suggest took : %s\n", duration)
}

func WithDurationLookup(function func(string) []cg.User, phone string) ([]cg.User, time.Duration) {
	startAt := time.Now()
	users := function(phone)
	return users, time.Since(startAt)
}

func WithDurationSuggestion(function func(cg.User) []cg.Suggestion, user cg.User) ([]cg.Suggestion, time.Duration) {
	startAt := time.Now()
	suggestions := function(user)
	return suggestions, time.Since(startAt)
}
