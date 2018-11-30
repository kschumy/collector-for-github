package types

type Accessible int

const (
	PublicOrPrivate Accessible = iota
	Public
	Private
)

func (accessible Accessible) IsValid() bool {
	return accessible == PublicOrPrivate || accessible == Public || accessible == Private
}

