package events

import (
	"time"
)

type FightItem struct {
	id    uint16
	title string
	FightEvent
}

type FightEvent struct {
	title Action // enum one of "dialogue" | "quest" | "move" | "craft" | "rest" | "fight"
	tick  time.Duration
}

var _ Ticker = (*FightItem)(nil)

func (d *FightItem) Tick() time.Duration {
	return time.Second
}

func (d *FightItem) ElapsedTimes() time.Duration {
	return d.tick
}

func (d *FightItem) EventTitle() string {
	return d.FightEvent.title.String()
}

func (d *FightItem) ItemID() uint16 {
	return d.id
}

func NewFightItem(tick string, id uint16) *FightItem {

	duration, _ := time.ParseDuration(tick)

	return &FightItem{
		id:    id,
		title: "Fight item",
		FightEvent: FightEvent{
			title: Fight,
			tick:  duration,
		},
	}
}
