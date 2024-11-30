package ui

import (
	"fmt"

	"github.com/rivo/tview"

	
)
func initClusterContextBox(app  *tview.Application) *tview.TextView {

textView:=tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetChangedFunc(func(){
	app.Draw()
})


go func(){
	fmt.Fprintf(textView,"%s",clusterContext)
}()
return textView

}

func createDashboard(appUi *AppUI) (flex *tview.Flex) {
	// Initialize the application and flex container
	appUi.AppConfig.Table.initDefaultTable(appUi)
	
	flex = tview.NewFlex().SetDirection(tview.FlexRow).
	AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).AddItem(initClusterContextBox(appUi.AppConfig.App), 0, 1, false).AddItem(tview.NewBox().SetBorder(true), 0, 1, false),0,1,false).
	AddItem(appUi.AppConfig.Table.table, 0, 5, true)
		
	return flex

}