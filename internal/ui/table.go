package ui

import (
	"bytes"
	"fmt"
	"time"

	client "github.com/Gkemhcs/kubedash/internal/k8s"
	"github.com/Gkemhcs/kubedash/internal/k8s/objects"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// var table *tview.Table = tview.NewTable().SetSelectable(true, false)
type TableConfig struct {
	table *tview.Table
}

func (tableConfig *TableConfig) startAutoRefresh(appUI *AppUI) {
	ticker := time.NewTicker(1 * time.Second) // Tick every second

	go func() {
		defer ticker.Stop()
		for range ticker.C {
			// Fetch the latest data

			// Update the table
			appUI.AppConfig.App.QueueUpdateDraw(func() {
				tableConfig.initCustom(appUI.getCurrentKind(), appUI.getCurrentNamespace(), appUI.K8sConfig)
			})
		}
	}()

}
func (tableConfig *TableConfig) initDefaultTable(appUI *AppUI) {

	tableConfig.table = tview.NewTable().SetSelectable(true, false)
	table := tableConfig.table
	headers := client.GetPodFields()
	pods, err := objects.ListPods(appUI.K8sConfig.DefaultNamespace, appUI.K8sConfig) // Assuming this function can handle namespace filtering
	if err != nil {
		fmt.Print(err)
		return
	}
	for col, header := range headers {
		table.SetCell(0, col, tview.NewTableCell(fmt.Sprintf("%-20s", header)).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignCenter)).Select(1, 1)

		// Get filtered pods based on search text

		// Populate table with filtered data
		for row, pod := range pods {
			for column, field := range pod {
				table.SetCell(row+1, column, tview.NewTableCell(field).SetAlign(tview.AlignLeft))

			}

		}
	}
	table.SetBorder(true).SetTitle(fmt.Sprintf("%s(%s)[%d]", "Pods", appUI.CurrentNamespace, len(pods))).SetTitleColor(tcell.ColorBlueViolet)

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'd' {

			row, _ := table.GetSelection()
			cell := table.GetCell(row, 0)

			if cell != nil {

				appUI.AppConfig.TextView.Clear()
				var bytedata bytes.Buffer
				var err error

				switch appUI.CurrentKind {
				case "pod":
					bytedata, err = objects.DescribePod(cell.Text, appUI.CurrentNamespace, appUI.K8sConfig)
				case "deployment":

					bytedata, err = objects.DescribeDeployment(cell.Text, appUI.CurrentNamespace, appUI.K8sConfig)
				case "configmap":

					bytedata, err = objects.DescribeConfigMap(cell.Text, appUI.CurrentNamespace, appUI.K8sConfig)
				case "clusterrolebinding":

					bytedata, err = objects.DescribeClusterRoleBinding(cell.Text, appUI.CurrentNamespace, appUI.K8sConfig)
				case "clusterrole":

					bytedata, err = objects.DescribeClusterRole(cell.Text, appUI.CurrentNamespace, appUI.K8sConfig)
				default:
					appUI.LoggerConfig.Logger.Info("default was chosen", appUI.CurrentKind)
				}
				if err != nil {
					appUI.LoggerConfig.Logger.Warn(err)
					fmt.Print(err)
				}

				appUI.LoggerConfig.Logger.Info(bytedata.String())

				appUI.AppConfig.TextView.SetText(bytedata.String()).SetScrollable(true).SetWrap(true).SetTitle(fmt.Sprintf("Describe(%s/%s)", appUI.getCurrentKind(), cell.Text))

				appUI.AppConfig.Pages.SwitchToPage("details")
			} else {
				fmt.Println("No cell selected")
			}
		}

		if event.Rune() == rune(tcell.KeyCtrlD) {

			row, _ := table.GetSelection()
			cell := table.GetCell(row, 0)
			var err error
			switch appUI.CurrentKind {
			case "pod":
				err = objects.DeletePod(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "deployment":
				err = objects.DeleteDeployment(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "clusterrole":
				err = objects.DeleteClusterRole(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "configmap":
				err = objects.DeleteConfigMap(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "cronjob":
				err = objects.DeleteCronJob(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "daemonset":
				err = objects.DeleteDaemonSet(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "ingress":
				err = objects.DeleteIngress(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "job":
				err = objects.DeleteJob(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "persistentvolume":
				err = objects.DeletePersistentVolume(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "replicaset":
				err = objects.DeleteReplicaSet(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "rolebinding":
				err = objects.DeleteRoleBinding(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "role":
				err = objects.DeleteRole(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "secret":
				err = objects.DeleteSecret(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "serviceaccount":
				err = objects.DeleteServiceAccount(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "service":
				err = objects.DeleteService(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			case "storageclass":
				err = objects.DeleteStorageClass(cell.Text, appUI.getCurrentNamespace(), appUI.K8sConfig)
			}
			if err != nil {
				appUI.LoggerConfig.Logger.Warn(err)

			}
			appUI.AppConfig.Pages.SwitchToPage("root")

		}
		return event
	})
	tableConfig.startAutoRefresh(appUI)
}

func (tableConfig *TableConfig) initCustom(kind string, namespace string, k8sClient *client.K8sConfig) {
	table := tableConfig.table
	table.Clear()

	var headers []string
	var dataList [][]string
	var err error
	switch kind {
	case "deployment":
		headers = client.GetDeploymentFields()
		dataList, err = objects.ListDeployments(namespace, k8sClient)
	case "configmap":
		headers = client.GetConfigMapFields()
		dataList, err = objects.ListConfigMaps(namespace, k8sClient)
	case "secret":
		headers = client.GetSecretFields()
		dataList, err = objects.ListSecrets(namespace, k8sClient)
	case "daemonset":
		headers = client.GetDaemonSetFields()
		dataList, err = objects.ListDaemonSets(namespace, k8sClient)
	case "replicaset":
		headers = client.GetReplicaSetFields()
		dataList, err = objects.ListReplicaSets(namespace, k8sClient)
	case "service":
		headers = client.GetServiceFields()
		dataList, err = objects.ListServices(namespace, k8sClient)
	case "storageclass":
		headers = client.GetStorageClassFields()
		dataList, err = objects.ListStorageClasses(namespace, k8sClient)
	case "job":
		headers = client.GetJobFields()
		dataList, err = objects.ListJobs(namespace, k8sClient)
	case "cronjob":
		headers = client.GetCronJobFields()
		dataList, err = objects.ListCronJobs(namespace, k8sClient)
	case "ingress":
		headers = client.GetIngressFields()
		dataList, err = objects.ListIngresses(namespace, k8sClient)
	case "persistentvolume":
		headers = client.GetPersistentVolumeFields()
		dataList, err = objects.ListPersistentVolumes(namespace, k8sClient)
	case "serviceaccount":
		headers = client.GetServiceAccountFields()
		dataList, err = objects.ListServiceAccounts(namespace, k8sClient)
	case "role":
		headers = client.GetRoleFields()
		dataList, err = objects.ListRoles(namespace, k8sClient)
	case "rolebinding":
		headers = client.GetRoleBindingFields()
		dataList, err = objects.ListRoleBindings(namespace, k8sClient)
	case "clusterrole":
		headers = client.GetClusterRoleFields()
		dataList, err = objects.ListClusterRoles(namespace, k8sClient)
	case "clusterrolebinding":
		headers = client.GetClusterRoleBindingFields()
		dataList, err = objects.ListClusterRoleBindings(namespace, k8sClient)
	case "endpoint":
		headers = client.GetEndpointFields()
		dataList, err = objects.ListEndpoints(namespace, k8sClient)
	case "node":
		headers = client.GetNodeFields()
		dataList, err = objects.ListNodes(namespace, k8sClient)

	default:
		headers = client.GetPodFields()
		dataList, err = objects.ListPods(namespace, k8sClient)

	}
	// Add table headers
	if err != nil {
		fmt.Print(err)
	}
	for col, header := range headers {
		table.SetCell(0, col, tview.NewTableCell(fmt.Sprintf("%-20s", header)).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignCenter))
	}

	// Populate table with filtered data
	for row, object := range dataList {

		for column, field := range object {

			table.SetCell(row+1, column, tview.NewTableCell(field).SetAlign(tview.AlignLeft))

		}

	}
	table.SetBorder(true).SetTitle(fmt.Sprintf("%s(%s)[%d]", kind, namespace, len(dataList))).SetTitleColor(tcell.ColorBlueViolet)

}
