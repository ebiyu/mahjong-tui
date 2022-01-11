package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

type PointIndicator struct {
	pointTextViews     [4]*tview.TextView
	gameStatusView     *tview.Grid
	gameStatusTextView *tview.TextView
	deckCountTextView  *tview.TextView
	ui                 *tview.Grid
}

func NewPointIndicator() *PointIndicator {
	pi := new(PointIndicator)

	for i := 0; i < 4; i++ {
		pi.pointTextViews[i] = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("25000")
	}

	pi.gameStatusTextView = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("東1局")
	pi.deckCountTextView = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("")

	pi.gameStatusView = tview.NewGrid().SetRows(0, 0).SetColumns(0).
		AddItem(pi.gameStatusTextView, 0, 0, 1, 1, 0, 0, false).
		AddItem(pi.deckCountTextView, 1, 0, 1, 1, 0, 0, false)

	pi.ui = tview.NewGrid().SetRows(0, 0, 0).SetColumns(0, 0, 0).
		AddItem(pi.pointTextViews[0], 2, 1, 1, 1, 0, 0, false).
		AddItem(pi.pointTextViews[1], 1, 2, 1, 1, 0, 0, false).
		AddItem(pi.pointTextViews[2], 0, 1, 1, 1, 0, 0, false).
		AddItem(pi.pointTextViews[3], 1, 0, 1, 1, 0, 0, false).
		AddItem(pi.gameStatusView, 1, 1, 1, 1, 0, 0, false)
	pi.ui.SetBorder(true)
	return pi
}

func (pi *PointIndicator) SetPoints(points [4]int) *PointIndicator {
	for i := 0; i < 4; i++ {
		pi.pointTextViews[i].SetText(fmt.Sprintf("%d", points[i]))
	}
	return pi
}

func getWindChar(num int) string {
	switch num {
	case 0:
		return "東"
	case 1:
		return "南"
	case 2:
		return "西"
	case 3:
		return "北"
	default:
		return ""
	}
}

func (pi *PointIndicator) SetRound(wind int, number int) *PointIndicator {
	pi.gameStatusTextView.SetText(fmt.Sprintf("%s %d 局", getWindChar(wind), number))
	return pi
}
func (pi *PointIndicator) SetDeckCount(count int) *PointIndicator {
	pi.deckCountTextView.SetText(fmt.Sprintf("余 %d", count))
	return pi
}

func (i *PointIndicator) UI() tview.Primitive {
	return i.ui
}
