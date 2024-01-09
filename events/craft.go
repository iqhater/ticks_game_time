package events

import (
	"time"
)

type CraftItem struct {
	id    uint16
	title string
	// options []string
	CraftEvent
}

type CraftEvent struct {
	title Action // enum one of "dialogue" | "quest" | "move" | "craft" | "rest"
	tick  time.Duration
}

var _ Ticker = (*CraftItem)(nil)

func (d *CraftItem) Tick() time.Duration {
	return time.Second
}

func (d *CraftItem) ElapsedTimes() time.Duration {
	return d.tick
}

func (d *CraftItem) EventTitle() string {
	return d.CraftEvent.title.String()
}

func (d *CraftItem) ItemID() uint16 {
	return d.id
}

func NewCraftItem(tick string, id uint16) *CraftItem {

	duration, _ := time.ParseDuration(tick)

	return &CraftItem{
		id:    id,
		title: "Craft item",
		CraftEvent: CraftEvent{
			title: Craft,
			tick:  duration,
		},
	}
}
