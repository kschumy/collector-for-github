package types

type SearchIn int

const (
	AnyLocation SearchIn = iota
	Body
	Comments
	Title
)

func (searchIn SearchIn) IsValid() bool {
	return searchIn == AnyLocation || searchIn == Body || searchIn == Comments || searchIn == Title
}
