package events

import (
	"time"
)

type RestItem struct {
	id    uint16
	title string
	// options []string
	RestEvent
}

type RestEvent struct {
	title Action // enum one of "dialogue" | "quest" | "move" | "craft" | "rest"
	tick  time.Duration
}

var _ Ticker = (*RestItem)(nil)

func (d *RestItem) Tick() time.Duration {
	return time.Second
}

func (d *RestItem) ElapsedTimes() time.Duration {
	return d.tick
}

func (d *RestItem) EventTitle() string {
	return d.RestEvent.title.String()
}

func (d *RestItem) ItemID() uint16 {
	return d.id
}

func NewRestItem(tick string, id uint16) *RestItem {

	duration, _ := time.ParseDuration(tick)

	return &RestItem{
		id:    id,
		title: "Rest item",
		RestEvent: RestEvent{
			title: Rest,
			tick:  duration,
		},
	}
}
