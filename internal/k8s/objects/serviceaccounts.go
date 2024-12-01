package objects

import (
	"context"
	"time"

	client "github.com/Gkemhcs/kubedash/internal/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//	"gopkg.in/yaml.v3"
)

// ListServiceAccounts list out the clusterrolebings in cluster and returns it
// parameters:
// - namespace(string):  the namespace to which  we need to scope  our search
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of serviceaccounts
// - error : if any error occurs returns that otherwise returns nil
func ListServiceAccounts(namespace string, clientSet *client.K8sConfig) ([][]string, error) {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}

	serviceAccounts, err := clientSet.Client.CoreV1().ServiceAccounts(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var serviceAccountList [][]string
	for _, serviceAccount := range serviceAccounts.Items {
		serviceAccountList = append(serviceAccountList, []string{
			serviceAccount.Name,

			formatDuration(time.Since(serviceAccount.CreationTimestamp.Time)),
		})
	}
	return serviceAccountList, nil

}

// DeleteServiceAccount  delete the ServiceAccount and returns the status of deletion
// Parameters:
// - serviceAccountName : the name of clusterRoleBinding we need to delete
// - namespace: tha namespace to which we need to scope our search
// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - if deletion succeeds returns nil, otherwise returns the error occured
func DeleteServiceAccount(serviceAccountName string, namespace string, clientSet *client.K8sConfig) error {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}
	err := clientSet.Client.CoreV1().ServiceAccounts(namespace).Delete(context.TODO(), serviceAccountName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
