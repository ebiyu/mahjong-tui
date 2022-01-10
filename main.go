package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	pies := []string{"1m", "1m", "1m", "2m", "3m", "4m", "5m", "6m", "7m", "8m", "9m", "9m", "9m", "1p"}

	buttons := []tview.Primitive{}
	for i := 0; i < 14; i++ {
		buttons = append(buttons, tview.NewButton(pies[i]))
	}

	activePie := 0
	buttonFlex := tview.NewFlex()
	for i := 0; i < 13; i++ {
		buttonFlex.AddItem(buttons[i], 5, 1, false)
	}

	bottom := tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(buttonFlex, 0, 1, false).
		AddItem(buttons[13], 5, 1, false).
		AddItem(tview.NewBox(), 0, 1, false)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 1, 1, false).
		AddItem(tview.NewBox().SetBorder(true), 0, 9, false).
		AddItem(bottom, 1, 1, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'q':
				app.Stop()
				return nil
			case 'h':
				if activePie > 0 {
					activePie -= 1
				}
				app.SetFocus(buttons[activePie])
			case 'l':
				if activePie < 13 {
					activePie += 1
				}
				app.SetFocus(buttons[activePie])
			}
		}
		return event
	})

	if err := app.SetRoot(flex, true).SetFocus(buttons[0]).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
