package types

type ObjectType int

const (
	Issues ObjectType = iota
	PRs
)

func (objectType ObjectType) IsValid() bool {
	return objectType == Issues || objectType == PRs
}