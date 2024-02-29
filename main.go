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

	"github.com/fatih/color"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	ui.ShowUI()

	// ticks store
	var totalTicks time.Duration
	var minEventID, maxEventID = 1, 6

	for scanner.Scan() {

		parsedNumber, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("%v\n", err)
			continue
		}

		if parsedNumber < minEventID || parsedNumber > maxEventID {
			log.Printf("%v\n", "Unknown action number!")
			continue
		}

		var action events.Ticker
		action = events.NewEvent(parsedNumber)
		ticker := time.NewTicker(action.Tick())

		done := make(chan struct{})

		go func() {

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
					totalTicks += action.ElapsedTimes()
					fmt.Printf("EVENT TYPE: %s ITEM ID: %d TICK: %v \n", action.EventTitle(), action.ItemID(), t.Round(time.Millisecond))
				}
			}
		}()

		time.Sleep(3 * time.Second)

		ticker.Stop()
		done <- struct{}{}
		close(done)

		totalTime := color.New(color.FgYellow).PrintfFunc()
		totalTime("Total time: %s\n", totalTicks.Round(time.Millisecond))
		ui.ShowUI()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
