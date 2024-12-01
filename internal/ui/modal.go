package ui 
import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/Gkemhcs/kubedash/internal/k8s/objects"
	
)

// showModal displays the modal for confirmation of delete operation
func showModal(appUI *AppUI,resourceName string){
	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}
	var err error
	modalBox := tview.NewModal().
				SetText(fmt.Sprintf("Do you want to delete the  %s", resourceName)).
				AddButtons([]string{"Confirm", "Cancel"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "Confirm" {
						switch appUI.CurrentKind {
						case "pod":
							err = objects.DeletePod(resourceName, appUI.getCurrentNamespace(), appUI.K8sConfig)
						case "deployment":
							err = objects.DeleteDeployment(resourceName, appUI.getCurrentNamespace(), appUI.K8sConfig)
						case "clusterrole":
							err = objects.DeleteClusterRole(resourceName, appUI.K8sConfig)
						case "configmap":
							err = objects.DeleteConfigMap(resourceName, appUI.getCurrentNamespace(), appUI.K8sConfig)
						case "cronjob":
							err = objects.DeleteCronJob(resourceName, appUI.getCurrentNamespace(), appUI.K8sConfig)
						case "daemonset":
							err = objects.DeleteDaemonSet(resourceName, appUI.getCurrentNamespace(), appUI.K8sConfig)
						case "ingress":
							err = objects.DeleteIngress(resourceName, appUI.getCurrentNamespace(), appUI.K8sConfig)
						case "job":
							err = objects.DeleteJob(resourceName, appUI.getCurrentNamespace(), appUI.K8sConfig)
						case "persistentvolume":
							err = objects.DeletePersistentVolume(resourceName, appUI.K8sConfig)
						case "replicaset":
							err = objects.DeleteReplicaSet(resourceName, appUI.getCurrentNamespace(), appUI.K8sConfig)
						case "rolebinding":
							err = objects.DeleteRoleBinding(resourceName, appUI.getCurrentNamespace(), appUI.K8sConfig)
						case "role":
							err = objects.DeleteRole(resourceName, appUI.getCurrentNamespace(), appUI.K8sConfig)
						case "secret":
							err = objects.DeleteSecret(resourceName, appUI.getCurrentNamespace(), appUI.K8sConfig)
						case "serviceaccount":
							err = objects.DeleteServiceAccount(resourceName, appUI.getCurrentNamespace(), appUI.K8sConfig)
						case "service":
							err = objects.DeleteService(resourceName, appUI.getCurrentNamespace(), appUI.K8sConfig)
						case "storageclass":
							err = objects.DeleteStorageClass(resourceName, appUI.K8sConfig)
						case "namespace":
							err = objects.DeleteNamespace(resourceName, appUI.K8sConfig)
						}
						if err != nil {
							appUI.LoggerConfig.Logger.Warn(err)

						}
						appUI.AppConfig.Pages.RemovePage("modal")
						
					} else {
						appUI.AppConfig.Pages.RemovePage("modal")
						

					}
				})

				appUI.AppConfig.Pages.AddPage("modal", modal(modalBox, 40, 10), true, true)


}