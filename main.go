package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/ebiyuu1121/mahjong-tui/game"
	"github.com/ebiyuu1121/mahjong-tui/ui/kawaUI"
)

const (
	JICHA    = 0
	SHIMOCHA = 1
	TOIMEN   = 2
	KAMICHA  = 3
)

func main() {
	app := tview.NewApplication()
	game := game.Init()

	// ui
	buttons := []*tview.Button{}
	for i := 0; i < 14; i++ {
		buttons = append(buttons, tview.NewButton(game.Tehai()[JICHA][i]))
	}

	// kawa
	kawaUIList := [4]kawaUI.KawaUI{
		JICHA:    kawaUI.Init(JICHA),
		SHIMOCHA: kawaUI.Init(SHIMOCHA),
		TOIMEN:   kawaUI.Init(TOIMEN),
		KAMICHA:  kawaUI.Init(KAMICHA),
	}

	// update kawa
	update := func() {
		for i := 0; i < 4; i++ {
			kawaUIList[i].SetTiles(game.Kawa()[i])
		}
		for i := 0; i < 14; i++ {
			pai := game.Tehai()[JICHA][i]
			switch pai {
			case "1z":
				buttons[i].SetLabel("東")
			case "2z":
				buttons[i].SetLabel("南")
			case "3z":
				buttons[i].SetLabel("西")
			case "4z":
				buttons[i].SetLabel("北")
			case "5z":
				buttons[i].SetLabel("白")
			case "6z":
				buttons[i].SetLabel("發")
			case "7z":
				buttons[i].SetLabel("中")
			default:
				buttons[i].SetLabel(pai)
			}
			switch pai[1] {
			case 'm':
				buttons[i].SetLabelColor(tcell.ColorPink)
			case 'p':
				buttons[i].SetLabelColor(tcell.ColorLightCyan)
			case 's':
				buttons[i].SetLabelColor(tcell.ColorLightGreen)
			default:
				buttons[i].SetLabelColor(tcell.ColorWhite)
			}
		}
	}
	update()

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
		AddItem(bottom, 3, 1, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			game.Kill(JICHA, 13).Ripai(JICHA).Tsumo(JICHA)
			update()
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
			case ' ':
				game.Kill(JICHA, activePie).Ripai(JICHA).Tsumo(JICHA)
				update()
			}
		}
		return event
	})

	if err := app.SetRoot(flex, true).SetFocus(buttons[0]).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
