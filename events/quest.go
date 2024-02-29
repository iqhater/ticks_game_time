package events

import (
	"time"
)

type QuestItem struct {
	id    uint16
	title string
	QuestEvent
}

type QuestEvent struct {
	title Action // enum one of "dialogue" | "quest" | "move" | "craft" | "rest"
	tick  time.Duration
}

var _ Ticker = (*QuestItem)(nil)

func (d *QuestItem) Tick() time.Duration {
	return time.Second
}

func (d *QuestItem) ElapsedTimes() time.Duration {
	return d.tick
}

func (d *QuestItem) EventTitle() string {
	return d.QuestEvent.title.String()
}

func (d *QuestItem) ItemID() uint16 {
	return d.id
}

func NewQuestItem(tick string, id uint16) *QuestItem {

	duration, _ := time.ParseDuration(tick)

	return &QuestItem{
		id:    id,
		title: "Quest item",
		QuestEvent: QuestEvent{
			title: Quest,
			tick:  duration,
		},
	}
}
