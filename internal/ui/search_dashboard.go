package ui

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func initInputField(appUi *AppUI) *tview.InputField {

	searchBar := tview.NewInputField().
		SetLabel("search>").SetPlaceholder("🔍Please enter kubernetes resource and namespace").
		SetFieldWidth(0)
	searchBar.SetDoneFunc(func(key tcell.Key) {

		if key == tcell.KeyEnter {

			inputText := searchBar.GetText()
			searches := strings.Split(inputText, " ")

			//initCustom(searches[0], searches[1], k8sClient)

			if len(searches) == 0 {
				// Handle empty input case
				appUi.AppConfig.Pages.SwitchToPage("root")
			} else if len(searches) == 1 {
				appUi.CurrentKind = searches[0]

				appUi.AppConfig.Table.initCustom(appUi.CurrentKind, appUi.K8sConfig.DefaultNamespace, appUi.K8sConfig)

			} else if len(searches) >= 2 {
				appUi.CurrentKind = searches[0]
				appUi.CurrentNamespace = searches[1]

				appUi.AppConfig.Table.initCustom(appUi.CurrentKind, appUi.CurrentNamespace, appUi.K8sConfig)
				// Handle input with "in" correctly
				fmt.Print(appUi.getCurrentKind())

			}

			appUi.AppConfig.Pages.SwitchToPage("root")

		}

	})
	return searchBar
}

func createSearchDashboard(appUi *AppUI) (searchFlex *tview.Flex) {
	// Initialize the application and flex container

	searchFlex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).AddItem(initClusterContextBox(appUi.AppConfig.App), 0, 1, false).AddItem(tview.NewBox().SetBorder(true), 0, 1, false), 0, 5, false).
		AddItem(initInputField(appUi), 0, 1, true).
		AddItem(appUi.AppConfig.Table.table, 0, 15, false)

		// Return the application and flex layout
	return searchFlex
}
