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
	title Action // enum one of "dialogue" | "quest" | "move" | "craft"
	tick  time.Duration
}

func (d *CraftItem) Tick() time.Duration {
	return d.tick
}

func (d *CraftItem) EventTitle() string {
	return d.CraftEvent.title.String()
}

func (d *CraftItem) ItemID() uint16 {
	return d.id
}

func NewCraftItem(tick time.Duration, id uint16) *CraftItem {
	return &CraftItem{
		id:    id,
		title: "Craft item",
		CraftEvent: CraftEvent{
			title: Craft,
			tick:  tick,
		},
	}
}
