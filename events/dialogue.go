package events

import (
	"time"
)

type DialogueItem struct {
	id    uint16
	title string
	DialogueEvent
}

type DialogueEvent struct {
	title Action // enum one of "dialogue" | "quest" | "move" | "craft" | "rest"
	tick  time.Duration
}

func (d *DialogueItem) Tick() time.Duration {
	return time.Second
}

var _ Ticker = (*DialogueItem)(nil)

func (d *DialogueItem) ElapsedTimes() time.Duration {
	return d.tick
}

func (d *DialogueItem) EventTitle() string {
	return d.DialogueEvent.title.String()
}

func (d *DialogueItem) ItemID() uint16 {
	return d.id
}

func NewDialogueItem(tick string, id uint16) *DialogueItem {

	duration, _ := time.ParseDuration(tick)

	return &DialogueItem{
		id:    id,
		title: "Dialogue item",
		DialogueEvent: DialogueEvent{
			title: Dialogue,
			tick:  duration,
		},
	}
}
