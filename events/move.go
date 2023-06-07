package events

import (
	"time"
)

type MoveItem struct {
	id    uint16
	title string
	// options []string
	MoveEvent
}

type MoveEvent struct {
	title Action // enum one of "dialogue" | "quest" | "move" | "craft"
	tick  time.Duration
}

func (d *MoveItem) Tick() time.Duration {
	return d.tick
}

func (d *MoveItem) EventTitle() string {
	return d.MoveEvent.title.String()
}

func (d *MoveItem) ItemID() uint16 {
	return d.id
}

func NewMoveItem(tick time.Duration, id uint16) *MoveItem {
	return &MoveItem{
		id:    id,
		title: "Move item",
		MoveEvent: MoveEvent{
			title: Move,
			tick:  tick,
		},
	}
}
