package objects

import (
	"context"
	"fmt"

	"time"

	client "github.com/Gkemhcs/kubedash/internal/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


// ListDaemonSets  list out the clusterrolebings in cluster and returns it 
// parameters:
// - namespace(string):  the namespace to which  we need to scope  our search
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of daemonsets
// - error : if any error occurs returns that otherwise returns nil 

func ListDaemonSets(namespace string, clientSet *client.K8sConfig) ([][]string, error) {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}

	daemonSets, err := clientSet.Client.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Print(err)
		return [][]string{}, err
	}
	var daemonSetList [][]string
	for _, daemonSet := range daemonSets.Items {

		daemonSetList = append(daemonSetList, []string{
			daemonSet.Name,
			fmt.Sprintf("%d", daemonSet.Status.DesiredNumberScheduled),
			fmt.Sprintf("%d", daemonSet.Status.CurrentNumberScheduled),
			fmt.Sprintf("%d", daemonSet.Status.NumberReady),
			fmt.Sprintf("%d", daemonSet.Status.UpdatedNumberScheduled),
			formatDuration(time.Since(daemonSet.CreationTimestamp.Time)),
		})

	}
	return daemonSetList, nil
}

func DeleteDaemonSet(daemonSetName string, namespace string, clientSet *client.K8sConfig) error {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}
	err := clientSet.Client.AppsV1().DaemonSets(namespace).Delete(context.TODO(), daemonSetName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
