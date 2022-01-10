package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/ebiyuu1121/tui-test/ui/kawaUI"
)

func haipai() []string {
	return []string{"1m", "1m", "1m", "2m", "3m", "4m", "5m", "6m", "7m", "8m", "9m", "9m", "9m", "1p"}
}

const (
	JICHA    = 0
	SHIMOCHA = 1
	TOIMEN   = 2
	KAMICHA  = 3
)

func main() {
	app := tview.NewApplication()

	pies := haipai()

	// ui
	buttons := []tview.Primitive{}
	for i := 0; i < 14; i++ {
		buttons = append(buttons, tview.NewButton(pies[i]))
	}

	// kawa
	kawaUIList := [4]kawaUI.KawaUI{
		JICHA:    kawaUI.Init(JICHA),
		SHIMOCHA: kawaUI.Init(SHIMOCHA),
		TOIMEN:   kawaUI.Init(TOIMEN),
		KAMICHA:  kawaUI.Init(KAMICHA),
	}

	// kawa test data
	kawaUIList[JICHA].Update([]string{"1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p"})
	kawaUIList[SHIMOCHA].Update([]string{"1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p"})
	kawaUIList[TOIMEN].Update([]string{"1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p"})
	kawaUIList[KAMICHA].Update([]string{"1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p"})

	activePie := 0
	buttonFlex := tview.NewFlex()
	for i := 0; i < 13; i++ {
		buttonFlex.AddItem(buttons[i], 5, 1, false)
	}

	pointIndicator := tview.NewGrid().
		SetRows(0, 0, 0).
		SetColumns(0, 0, 0).
		AddItem(
			tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("25000"),
			0, 1, 1, 1, 0, 0, false).
		AddItem(
			tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("25000"),
			1, 0, 1, 1, 0, 0, false).
		AddItem(
			tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("25000"),
			1, 2, 1, 1, 0, 0, false).
		AddItem(
			tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("25000"),
			2, 1, 1, 1, 0, 0, false).
		AddItem(
			tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("東2局"),
			1, 1, 1, 1, 0, 0, false)
	pointIndicator.SetBorder(true)

	bottom := tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(buttonFlex, 0, 1, false).
		AddItem(buttons[13], 5, 1, false).
		AddItem(tview.NewBox(), 0, 1, false)

	kawaRow := []tview.Primitive{
		tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(kawaUIList[TOIMEN].Grid, 0, 1, false).
			AddItem(tview.NewBox(), 0, 1, false),
		tview.NewFlex().
			AddItem(kawaUIList[KAMICHA].Grid, 0, 1, false).
			AddItem(pointIndicator, 0, 1, false).
			AddItem(kawaUIList[SHIMOCHA].Grid, 0, 1, false),
		tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(kawaUIList[JICHA].Grid, 0, 1, false).
			AddItem(tview.NewBox(), 0, 1, false),
	}
	kawa := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(kawaRow[0], 0, 1, false).
		AddItem(kawaRow[1], 0, 1, false).
		AddItem(kawaRow[2], 0, 1, false)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 1, 1, false).
		AddItem(kawa, 0, 9, false).
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
