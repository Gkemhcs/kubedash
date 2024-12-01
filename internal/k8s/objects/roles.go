package objects

import (
	"context"
	"time"

	client "github.com/Gkemhcs/kubedash/internal/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//	"gopkg.in/yaml.v3"
)

// ListRoles  list out the clusterrolebings in cluster and returns it
// parameters:
// - namespace(string):  the namespace to which  we need to scope  our search
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of roles
// - error : if any error occurs returns that otherwise returns nil
func ListRoles(namespace string, clientSet *client.K8sConfig) ([][]string, error) {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}

	roles, err := clientSet.Client.RbacV1().Roles(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var roleList [][]string
	for _, role := range roles.Items {
		roleList = append(roleList, []string{
			role.Name,

			formatDuration(time.Since(role.CreationTimestamp.Time)),
		})
	}
	return roleList, nil

}

// DeleteRole  delete the  Role and returns the status of deletion
// Parameters:
// - roleName : the name of clusterRoleBinding we need to delete
// - namespace: tha namespace to which we need to scope our search
// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - if deletion succeeds returns nil, otherwise returns the error occured
func DeleteRole(roleName string, namespace string, clientSet *client.K8sConfig) error {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}
	err := clientSet.Client.RbacV1().Roles(namespace).Delete(context.TODO(), roleName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
