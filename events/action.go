package events

import (
	"github.com/fatih/color"
)

type Action uint8

const (
	Dialogue Action = iota
	Quest
	Move
	Craft
	Rest
	Fight
)

func (a Action) String() string {

	switch a {
	case Dialogue:
		return "dialogue"
	case Quest:
		return color.New(color.FgCyan).SprintFunc()("quest")
	case Move:
		return color.New(color.FgBlue).SprintFunc()("move")
	case Craft:
		return color.New(color.FgGreen).SprintFunc()("craft")
	case Rest:
		return color.New(color.FgMagenta).SprintFunc()("rest")
	case Fight:
		return color.New(color.FgRed).SprintFunc()("fight")
	default:
		return "Unknown action type!"
	}
}
