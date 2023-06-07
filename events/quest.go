package events

import (
	"time"
)

type QuestItem struct {
	id    uint16
	title string
	// options []string
	QuestEvent
}

type QuestEvent struct {
	title Action // enum one of "dialogue" | "quest" | "move" | "craft"
	tick  time.Duration
}

func (d *QuestItem) Tick() time.Duration {
	return d.tick
}

func (d *QuestItem) EventTitle() string {
	return d.QuestEvent.title.String()
}

func (d *QuestItem) ItemID() uint16 {
	return d.id
}

func NewQuestItem(tick time.Duration, id uint16) *QuestItem {
	return &QuestItem{
		id:    id,
		title: "Quest item",
		QuestEvent: QuestEvent{
			title: Quest,
			tick:  tick,
		},
	}
}
