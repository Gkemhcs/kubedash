package objects

import (
	"context"
	"time"

	client "github.com/Gkemhcs/kubedash/internal/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//	"gopkg.in/yaml.v3"
)

// Listnamespaces  list out the namespaces in cluster and returns it
// parameters:
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of namespaces
// - error : if any error occurs returns that otherwise returns nil
func ListNamespaces(clientSet *client.K8sConfig) ([][]string, error) {
	

	namespaces, err := clientSet.Client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var namespaceList [][]string
	for _, namespace := range namespaces.Items {
		namespaceList = append(namespaceList, []string{
			namespace.Name,
			string(namespace.Status.Phase),
			formatDuration(time.Since(namespace.CreationTimestamp.Time)),
		})
	}
	return namespaceList, nil

}


func GetAllNamespacesNames(clientSet *client.K8sConfig)([]string,error){
var namespaceList []string 

namespaces, err := clientSet.Client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
if err != nil {
	return nil, err
}

for _, namespace := range namespaces.Items {
	namespaceList = append(namespaceList,namespace.Name) 
}
return namespaceList,nil 

}

// DeleteNamespace  delete the  Namespace and returns the status of deletion
// Parameters:
// - namespace : the name of namespace we need to delete

// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - if deletion succeeds returns nil, otherwise returns the error occured
func DeleteNamespace( namespace string, clientSet *client.K8sConfig) error {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}
	err := clientSet.Client.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
