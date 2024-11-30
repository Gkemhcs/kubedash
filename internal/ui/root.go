package ui

import (
	client "github.com/Gkemhcs/kubedash/internal/k8s"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Gkemhcs/kubedash/internal/utils"
)

type AppConfig struct {
	App      *tview.Application
	Pages    *tview.Pages
	Table    TableConfig
	TextView *tview.TextView
}

type AppUI struct {
	AppConfig        *AppConfig
	CurrentKind      string
	CurrentNamespace string
	K8sConfig        *client.K8sConfig
	LoggerConfig     *utils.LoggerConfig
}

func (ui *AppUI) getCurrentKind() string {
	return ui.CurrentKind
}
func (ui *AppUI) getCurrentNamespace() string {
	return ui.CurrentNamespace
}
func (ui *AppUI) InitDashboard() error {
	k8sConfig := client.K8sConfig{}

	err := k8sConfig.InitClient()
	if err != nil {
		return err
	}
	ui.K8sConfig = k8sConfig.GetClient()

	loggerConfig := utils.LoggerConfig{}
	ui.LoggerConfig = loggerConfig.InitLogger()

	ui.AppConfig = &AppConfig{
		App:      tview.NewApplication(),
		Pages:    tview.NewPages(),
		TextView: tview.NewTextView(),
	}

	ui.CurrentNamespace = "monitoring"
	ui.CurrentKind = "pod"
	ui.AppConfig.Pages.AddPage("root", createDashboard(ui), true, true)
	ui.AppConfig.Pages.AddPage("search", createSearchDashboard(ui), true, false)
	ui.AppConfig.Pages.AddPage("details", createTextView(ui), true, false)

	ui.AppConfig.Pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlS {
			// Switch back to the table view
			ui.AppConfig.Pages.SwitchToPage("search")
		}
		if event.Key() == tcell.KeyEsc {
			// Switch back to the table view
			ui.AppConfig.Pages.SwitchToPage("root")

		}
		if event.Key() == tcell.KeyCtrlB {
			// Switch back to the table view
			ui.AppConfig.Pages.SwitchToPage("details")

		}

		return event
	})
	return nil
}
