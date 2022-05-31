package contactgraph

type GraphInterface interface {
	Lookup(phoneNumber string) []User
	RLookup(phoneNumber string) []User
	Suggest(phoneNumber string) []User
}
