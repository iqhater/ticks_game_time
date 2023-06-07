package main

import (
	"fmt"
	"iqhater/tick_based_game_time/events"
	"time"
)

func generateActions() []events.Ticker {

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
}

func main() {

	/* scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		fmt.Println("Select action type:")
		fmt.Println("[1] - Dialogue event:")
		fmt.Println("[2] - Quest event:")
		fmt.Println("[3] - Move event:")

		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	} */

	actions := generateActions()
	now := time.Now()
	var totalTime time.Duration

	for _, action := range actions {

		done := make(chan struct{})

		ticker := time.NewTicker(action.Tick())

		go func() {

			for {
				select {
				case <-done:
					return
				case t := <-ticker.C:
					totalTime = now.Sub(t)
					fmt.Printf("Tick: %v Item ID: %d Event type: %s\n", t, action.ItemID(), action.EventTitle())
				}
				// time.Sleep(150 * time.Millisecond)
			}
		}()

		time.Sleep(time.Second)

		ticker.Stop()
		done <- struct{}{}
	}
	fmt.Println("Total time:", totalTime)
}
