package events

import "time"

type Ticker interface {
	Tick() time.Duration
	EventTitle() string
	ItemID() uint16
}
