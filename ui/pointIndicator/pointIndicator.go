package pointIndicator

import (
	"fmt"

	"github.com/rivo/tview"
)

type PointIndicator struct {
	pointTextViews     [4]*tview.TextView
	gameStatusTextView *tview.TextView
	ui                 *tview.Grid
}

func Init() *PointIndicator {
	pi := new(PointIndicator)
	for i := 0; i < 4; i++ {
		pi.pointTextViews[i] = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("25000")
	}
	pi.gameStatusTextView = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("東1局")
	pi.ui = tview.NewGrid().
		SetRows(0, 0, 0).
		SetColumns(0, 0, 0).
		AddItem(pi.pointTextViews[0], 2, 1, 1, 1, 0, 0, false).
		AddItem(pi.pointTextViews[1], 1, 2, 1, 1, 0, 0, false).
		AddItem(pi.pointTextViews[2], 0, 1, 1, 1, 0, 0, false).
		AddItem(pi.pointTextViews[3], 1, 0, 1, 1, 0, 0, false).
		AddItem(pi.gameStatusTextView, 1, 1, 1, 1, 0, 0, false)
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

func (i *PointIndicator) UI() tview.Primitive {
	return i.ui
}
