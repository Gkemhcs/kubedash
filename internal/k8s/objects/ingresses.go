package objects

import (
	"context"
	"fmt"
	client "github.com/Gkemhcs/kubedash/internal/k8s"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ingress_info struct {
	Name              string
	Ready             string
	UpdatedReplicas   int32
	AvailableReplicas int32
}

func ListIngresses(namespace string, clientSet *client.K8sConfig) ([][]string, error) {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}

	ingresses, err := clientSet.Client.NetworkingV1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Print(err)
		return [][]string{}, err
	}
	var ingressesList [][]string
	for _, ingress := range ingresses.Items {

		hosts := ""
		for _, rules := range ingress.Spec.Rules {
			hosts += fmt.Sprintf("%s ", rules.Host)
		}
		ports := []string{"90"}
		if len(ingress.Spec.TLS) != 0 {
			ports = append(ports, "443")
		}
		ingressesList = append(ingressesList, []string{
			ingress.Name,
			*ingress.Spec.IngressClassName,
			hosts,
			ingress.Status.LoadBalancer.Ingress[0].IP,
			strings.Join(ports, ", "),
			formatDuration(time.Since(ingress.CreationTimestamp.Time)),
		})

	}
	return ingressesList, nil
}

func DeleteIngress(ingressName string, namespace string, clientSet *client.K8sConfig) error {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}
	err := clientSet.Client.NetworkingV1().Ingresses(namespace).Delete(context.TODO(), ingressName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
