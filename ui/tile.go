package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UITile struct {
	tileString string
	ui         *tview.TextView
}

func NewUITile(tileString string) *UITile {
	tile := new(UITile)
	tile.tileString = tileString
	tile.ui = tview.NewTextView().
		SetTextAlign(tview.AlignCenter)

	tile.SetTileString(tileString)
	return tile
}

func (tile *UITile) SetTileString(tileString string) *UITile {
	tile.tileString = tileString

	switch tileString {
	case "1z":
		tile.ui.SetText("東")
	case "2z":
		tile.ui.SetText("南")
	case "3z":
		tile.ui.SetText("西")
	case "4z":
		tile.ui.SetText("北")
	case "5z":
		tile.ui.SetText("白")
	case "6z":
		tile.ui.SetText("發")
	case "7z":
		tile.ui.SetText("中")
	default:
		tile.ui.SetText(tileString)
	}
	switch tileString[1] {
	case 'm':
		tile.ui.SetTextColor(tcell.ColorPink)
	case 'p':
		tile.ui.SetTextColor(tcell.ColorLightCyan)
	case 's':
		tile.ui.SetTextColor(tcell.ColorLightGreen)
	}
	return tile
}

func (tile *UITile) TileString() string {
	return tile.tileString
}

func (tile *UITile) UI() tview.Primitive {
	return tile.ui
}
