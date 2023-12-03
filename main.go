package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

func getAnswer(n1, n2 string, sign string) int {
	num1, _ := strconv.Atoi(n1)
	num2, _ := strconv.Atoi(n2)
	switch sign {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "/":
		return num1 / num2
	case "*":
		return num1 * num2
	}
	return 0
}

func calculate(entry *widget.Entry) {
	operators := []string{"+", "-", "*", "/"}
	for _, operator := range operators {
		if strings.Contains(entry.Text, operator) {
			split := strings.Split(entry.Text, operator)
			result := getAnswer(split[0], split[1], operator)
			entry.Text = strconv.Itoa(result)
			entry.Refresh()
			return
		}
	}
}

func main() {
	calculator := app.New()
	window := calculator.NewWindow("Calculator")

	entry := widget.NewEntry()
	entry.MultiLine = true
	entry.PlaceHolder = "0"
	entry.Text = "0"

	buttons := container.NewGridWithColumns(4,
		widget.NewButton("7", func() { addNumber("7", entry) }),
		widget.NewButton("8", func() { addNumber("8", entry) }),
		widget.NewButton("9", func() { addNumber("9", entry) }),
		widget.NewButton("/", func() { addNumber("/", entry) }),

		widget.NewButton("4", func() { addNumber("4", entry) }),
		widget.NewButton("5", func() { addNumber("5", entry) }),
		widget.NewButton("6", func() { addNumber("6", entry) }),
		widget.NewButton("*", func() { addNumber("*", entry) }),

		widget.NewButton("1", func() { addNumber("1", entry) }),
		widget.NewButton("2", func() { addNumber("2", entry) }),
		widget.NewButton("3", func() { addNumber("3", entry) }),
		widget.NewButton("-", func() { addNumber("-", entry) }),

		widget.NewButton("0", func() { addNumber("0", entry) }),
		widget.NewButton(".", func() { addNumber(".", entry) }),
		widget.NewButton("=", func() { calculate(entry) }),
		widget.NewButton("+", func() { addNumber("+", entry) }),
		widget.NewButton("Clear", func() { clearEntry(entry) }),
	)

	window.SetContent(container.NewVBox(
		entry,
		container.NewGridWithColumns(1, buttons),
	))

	window.ShowAndRun()
}

func addNumber(s string, entry *widget.Entry) {
	text := entry.Text
	if text == "0" {
		text = ""
	}
	entry.Text = text + s
	entry.Refresh()
}

func clearEntry(entry *widget.Entry) {
	entry.Text = "0"
	entry.Refresh()
}
