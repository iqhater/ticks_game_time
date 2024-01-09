package events

import "time"

type Ticker interface {
	Tick() time.Duration
	ElapsedTimes() time.Duration
	EventTitle() string
	ItemID() uint16
}
