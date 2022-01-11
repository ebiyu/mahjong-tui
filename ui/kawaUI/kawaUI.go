package kawaUI

import (
	"github.com/rivo/tview"

	"github.com/ebiyuu1121/mahjong-tui/ui/tile"
)

const (
	JICHA    = 0
	SHIMOCHA = 1
	TOIMEN   = 2
	KAMICHA  = 3
)

type KawaUI struct {
	Grid      *tview.Grid
	direction int // 0: jicha, 1: shimocha, 2: toimen, 3: shimocha
}

func Init(direction int) KawaUI {
	kawa := KawaUI{direction: direction}
	switch kawa.direction {
	case JICHA, TOIMEN:
		kawa.Grid = tview.NewGrid().SetRows(0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0, 0)
	case KAMICHA, SHIMOCHA:
		kawa.Grid = tview.NewGrid().SetRows(0, 0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0)
	}
	return kawa
}

func (kawa KawaUI) SetTiles(tiles []string) {
	kawa.Grid.Clear()
	for i, v := range tiles {
		switch kawa.direction {
		case JICHA:
			kawa.Grid.AddItem(tile.Init(v).UI(), i/6, i%6, 1, 1, 0, 0, false)
		case SHIMOCHA:
			kawa.Grid.AddItem(tile.Init(v).UI(), 5-i%6, i/6, 1, 1, 0, 0, false)
		case TOIMEN:
			kawa.Grid.AddItem(tile.Init(v).UI(), 3-i/6, 5-i%6, 1, 1, 0, 0, false)
		case KAMICHA:
			kawa.Grid.AddItem(tile.Init(v).UI(), i%6, 3-i/6, 1, 1, 0, 0, false)
		}
	}
}
