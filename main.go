package main

import (
	"bufio"
	"fmt"
	"iqhater/tick_based_game_time/events"
	"iqhater/tick_based_game_time/ui"
	"log"
	"os"
	"strconv"
	"time"
)

/* func generateActions() []events.Ticker {

	result := []events.Ticker{}

	dialogueAmount := 2
	questAmount := 5
	moveAmount := 8
	craftAmount := 12

	for i := 0; i < dialogueAmount; i++ {
		dialogue := events.NewDialogueItem(600*time.Millisecond, uint16(i))
		result = append(result, dialogue)
	}

	for i := dialogueAmount; i < questAmount; i++ {
		quest := events.NewQuestItem(1000*time.Millisecond, uint16(i))
		result = append(result, quest)
	}

	for i := questAmount; i < moveAmount; i++ {
		move := events.NewMoveItem(100*time.Millisecond, uint16(i))
		result = append(result, move)
	}

	for i := moveAmount; i < craftAmount; i++ {
		craft := events.NewCraftItem(250*time.Millisecond, uint16(i))
		result = append(result, craft)
	}
	return result
} */

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	ui.ShowUI()

	// ticks store
	var totalTicks time.Duration

	for scanner.Scan() {

		parsedNumber, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("%v\n", err)
			continue
		}

		if parsedNumber < 1 || parsedNumber > 6 {
			log.Printf("%v\n", "Unknown action number!")
			continue
		}

		var action events.Ticker
		action = events.NewEvent(parsedNumber)
		ticker := time.NewTicker(action.Tick())

		done := make(chan struct{})

		go func() {
			// defer close(done)

			for {

				// guarantee to be executed
				select {
				case <-done:
					return
				default:
				}

				select {
				case <-done:
					return
				case t := <-ticker.C:
					// totalTicks += time.Duration(t.Nanosecond())
					totalTicks += action.ElapsedTimes()
					fmt.Printf("EVENT TYPE: %s ITEM ID: %d TICK: %v \n", action.EventTitle(), action.ItemID(), t.Round(time.Millisecond))
				}
			}
		}()

		time.Sleep(3 * time.Second)

		ticker.Stop()
		done <- struct{}{}
		close(done)

		fmt.Println("Total time:", totalTicks.Round(time.Millisecond))
		ui.ShowUI()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
