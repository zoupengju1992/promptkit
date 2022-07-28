package main

import (
	"fmt"
	"os"

	"github.com/erikgeiser/promptkit/selection"
)

func main() {
	prompt := []*selection.Choice{ // selection提示器信息列表
		&selection.Choice{Index: 1, String: "Horse", Value: 1},
		&selection.Choice{Index: 2, String: "Car", Value: 2},
		&selection.Choice{Index: 3, String: "Plane", Value: 3},
		&selection.Choice{Index: 4, String: "Bike", Value: 4},
	}
	sp := selection.New("What do you pick?", prompt)
	sp.PageSize = 3

	choice, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the final choice
	_ = choice
}
