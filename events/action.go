package events

type Action uint8

const (
	Dialogue Action = iota
	Quest
	Move
	Craft
)

func (a Action) String() string {

	switch a {
	case Dialogue:
		return "dialogue"
	case Quest:
		return "quest"
	case Move:
		return "move"
	case Craft:
		return "craft"
	default:
		return "Unknown action type!"
	}
}
