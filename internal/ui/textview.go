package ui

import (
	//	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func createTextView(appUI *AppUI) *tview.Flex {
	textView := appUI.AppConfig.TextView.SetScrollable(true)
	textView.SetDynamicColors(true).
		SetRegions(true).SetBorder(true)

	textViewFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).AddItem(initClusterContextBox(appUI.AppConfig.App), 0, 1, false).AddItem(tview.NewBox().SetBorder(true), 0, 1, false), 0, 5, false).
		AddItem(textView, 0, 15, true)

	// Return the application and flex layout

	return textViewFlex

}
