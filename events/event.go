package events

import "log"

func NewEvent(parsedNumber int) Ticker {

	switch parsedNumber {
	case 1:
		return NewDialogueItem("600ms", uint16(1))
	case 2:
		return NewQuestItem("5m", uint16(2))
	case 3:
		return NewMoveItem("100ms", uint16(3))
	case 4:
		return NewCraftItem("2s", uint16(4))
	case 5:
		return NewRestItem("7h", uint16(5))
	case 6:
		return NewFightItem("10m", uint16(6))
	default:
		log.Printf("%v\n", "Unknown action number!")
	}
	return nil
}
