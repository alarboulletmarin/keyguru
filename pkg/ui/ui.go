package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type UI struct {
	reader *bufio.Reader
}

func NewUI() *UI {
	return &UI{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (ui *UI) DisplayMessage(message string) {
	fmt.Println(message)
}

func (ui *UI) ReadInput() (string, error) {
	input, err := ui.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func (ui *UI) ReadIntInput() (int, error) {
	input, err := ui.ReadInput()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(input)
}

func (ui *UI) DisplayResults(wpm float64, accuracy float64, errors int, totalTime float64) {
	fmt.Printf("WPM: %.2f\n", wpm)
	fmt.Printf("Accuracy: %.2f%%\n", accuracy*100)
	fmt.Printf("Errors: %d\n", errors)
	fmt.Printf("Total Time: %.2f seconds\n", totalTime)
}
