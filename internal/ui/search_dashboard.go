package ui

import (
	"fmt"
	"strings"

	"github.com/Gkemhcs/kubedash/internal/k8s"
	"github.com/Gkemhcs/kubedash/internal/k8s/objects"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// filterSuggestions filters suggestions based on input
func filterSuggestions(input string,resourceType string ,options []string) []string {
	var suggestions []string
	if resourceType=="" {
		for _, option := range options {
			if strings.HasPrefix(option, input) {
				suggestions = append(suggestions, option)
			}
		}
	}else{
		for _, option := range options {
			if strings.HasPrefix(option, input) {
				suggestions = append(suggestions,fmt.Sprintf("%s %s",resourceType,option))
			}
		}
	}
	
	return suggestions
}

func fetchNamespaces(clientset *k8s.K8sConfig) []string {
	namespaces, err := objects.GetAllNamespacesNames(clientset)
	if err != nil {
		return []string{}
	}
	return namespaces
}
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
	searchBar.SetAutocompleteFunc(func(currentText string) []string {
		parts := strings.Split(currentText, " ")

		// First word (resource type) autocomplete
		if len(parts) == 1 {
			return filterSuggestions(parts[0],"", k8s.GetAllResources())
		}

		// Second word (namespace) autocomplete
		if len(parts) == 2 {
			if k8s.IsNamespaced(parts[0]) {
				namespaces := fetchNamespaces(appUi.K8sConfig)
				return filterSuggestions(parts[1], parts[0],namespaces)
			}
			return nil 
		
		}

		return nil
	})
	

	return searchBar
}

func createSearchDashboard(appUi *AppUI) (searchFlex *tview.Flex) {
	// Initialize the application and flex container

	searchFlex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).AddItem(initClusterContextBox(appUi.AppConfig.App), 0, 1, false).AddItem(initShortcutBox(), 0, 1, false), 0, 5, false).
		AddItem(initInputField(appUi), 0, 1, true).
		AddItem(appUi.AppConfig.Table.table, 0, 15, false)

		// Return the application and flex layout
	return searchFlex
}
