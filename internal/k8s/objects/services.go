package objects

import (
	"context"
	"fmt"
	"time"

	client "github.com/Gkemhcs/kubedash/internal/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


// ListServices  list out the clusterrolebings in cluster and returns it 
// parameters:
// - namespace(string):  the namespace to which  we need to scope  our search
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of services
// - error : if any error occurs returns that otherwise returns nil 

func ListServices(namespace string, clientSet *client.K8sConfig) ([][]string, error) {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}

	services, err := clientSet.Client.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Print(err)
		return [][]string{}, err
	}
	var serviceList [][]string
	for _, service := range services.Items {

		var portString string = ""
		for _, port := range service.Spec.Ports {
			portString = portString + fmt.Sprintf("%s:%d 📎%d ", port.Name, port.Port, port.NodePort)
		}
		serviceList = append(serviceList, []string{

			service.Name,
			string(service.Spec.Type),
			service.Spec.ClusterIP,
			service.Spec.LoadBalancerIP,
			portString,
			formatDuration(time.Since(service.CreationTimestamp.Time)),
		})
	}
	return serviceList, nil
}

func DeleteService(serviceName string, namespace string, clientSet *client.K8sConfig) error {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}
	err := clientSet.Client.CoreV1().Services(namespace).Delete(context.TODO(), serviceName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
