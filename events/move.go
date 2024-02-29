package events

import (
	"time"
)

type MoveItem struct {
	id    uint16
	title string
	MoveEvent
}

type MoveEvent struct {
	title Action // enum one of "dialogue" | "quest" | "move" | "craft" | "rest"
	tick  time.Duration
}

var _ Ticker = (*MoveItem)(nil)

func (d *MoveItem) Tick() time.Duration {
	return time.Second
}

func (d *MoveItem) ElapsedTimes() time.Duration {
	return d.tick
}

func (d *MoveItem) EventTitle() string {
	return d.MoveEvent.title.String()
}

func (d *MoveItem) ItemID() uint16 {
	return d.id
}

func NewMoveItem(tick string, id uint16) *MoveItem {

	duration, _ := time.ParseDuration(tick)

	return &MoveItem{
		id:    id,
		title: "Move item",
		MoveEvent: MoveEvent{
			title: Move,
			tick:  duration,
		},
	}
}
