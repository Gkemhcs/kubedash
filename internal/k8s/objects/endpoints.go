package objects

import (
	"context"
	"time"

	"fmt"
	client "github.com/Gkemhcs/kubedash/internal/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//	"gopkg.in/yaml.v3"
)

// ListEndpoints  list out the clusterrolebings in cluster and returns it
// parameters:
// - namespace(string):  the namespace to which  we need to scope  our search
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of endpoints
// - error : if any error occurs returns that otherwise returns nil
func ListEndpoints(namespace string, clientSet *client.K8sConfig) ([][]string, error) {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}

	endpoints, err := clientSet.Client.CoreV1().Endpoints(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var endpointList [][]string
	for _, endpoint := range endpoints.Items {
		addresses := []string{}
		for _, address := range endpoint.Subsets[0].Addresses {
			addresses = append(addresses, address.IP)
		}
		endpointList = append(endpointList, []string{
			endpoint.Name,
			fmt.Sprintf("%v", addresses),

			formatDuration(time.Since(endpoint.CreationTimestamp.Time)),
		})
	}
	return endpointList, nil

}
