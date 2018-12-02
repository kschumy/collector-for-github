package types

type Relative int

const (
	AnyDateTime Relative = iota
	AfterDateTime
	BeforeDateTime
)

func (r Relative) IsValid() bool {
	return r == AnyDateTime || r == AfterDateTime || r == BeforeDateTime
}
