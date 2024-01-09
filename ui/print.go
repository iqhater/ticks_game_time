package ui

import (
	"fmt"

	"github.com/fatih/color"
)

func ShowUI() {

	fmt.Println("Select action type:")

	dialogue := color.New(color.FgWhite).PrintfFunc()
	dialogue("[1] - Dialogue event:\n")

	quest := color.New(color.FgCyan).PrintfFunc()
	quest("[2] - Quest event:\n")

	move := color.New(color.FgBlue).PrintfFunc()
	move("[3] - Move event:\n")

	craft := color.New(color.FgGreen).PrintfFunc()
	craft("[4] - Craft event:\n")

	rest := color.New(color.FgMagenta).PrintfFunc()
	rest("[5] - Rest event:\n")

	fight := color.New(color.FgRed).PrintfFunc()
	fight("[6] - Fight event:\n")
}
