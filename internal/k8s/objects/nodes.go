package objects

import (
	"context"
	"time"

	"fmt"
	client "github.com/Gkemhcs/kubedash/internal/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//	"gopkg.in/yaml.v3"
)

// ListNodes  list out the clusterrolebings in cluster and returns it
// parameters:
// - namespace(string):
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of Nodes
// - error : if any error occurs returns that otherwise returns nil
func ListNodes(namespace string, clientSet *client.K8sConfig) ([][]string, error) {

	nodes, err := clientSet.Client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var nodeList [][]string
	for _, node := range nodes.Items {
		nodeList = append(nodeList, []string{
			node.Name,
			string(node.Status.Phase),
			fmt.Sprintf("%d", len(node.Spec.Taints)),
			node.Status.NodeInfo.KubeletVersion,
			node.Status.Capacity.Cpu().String(),
			node.Status.Capacity.Memory().String(),
			node.Status.Allocatable.Cpu().String(),
			node.Status.Allocatable.Memory().String(),
			formatDuration(time.Since(node.CreationTimestamp.Time)),
		})
	}
	return nodeList, nil

}
