package events

import (
	"time"
)

type DialogueItem struct {
	id    uint16
	title string
	// options []string
	DialogueEvent
}

type DialogueEvent struct {
	title Action // enum one of "dialogue" | "quest" | "move" | "craft"
	tick  time.Duration
}

func (d *DialogueItem) Tick() time.Duration {
	return d.tick
}

func (d *DialogueItem) EventTitle() string {
	return d.DialogueEvent.title.String()
}

func (d *DialogueItem) ItemID() uint16 {
	return d.id
}

func NewDialogueItem(tick time.Duration, id uint16) *DialogueItem {
	return &DialogueItem{
		id:    id,
		title: "Dialogue item",
		DialogueEvent: DialogueEvent{
			title: Dialogue,
			tick:  tick,
		},
	}
}
