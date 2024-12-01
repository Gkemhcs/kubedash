package objects

import (
	"context"
	"fmt"

	client "github.com/Gkemhcs/kubedash/internal/k8s"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListStorageClasses  list out the clusterrolebings in cluster and returns it
// parameters:
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of storageClasses
// - error : if any error occurs returns that otherwise returns nil
func ListStorageClasses(clientSet *client.K8sConfig) ([][]string, error) {

	storageClassments, err := clientSet.Client.StorageV1().StorageClasses().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Print(err)
		return [][]string{}, err
	}
	var storageClassList [][]string
	for _, storageClass := range storageClassments.Items {

		storageClassList = append(storageClassList, []string{
			storageClass.Name,
			storageClass.Provisioner,
			string(*storageClass.ReclaimPolicy),
			string(*storageClass.VolumeBindingMode),
			fmt.Sprintf("%v", *storageClass.AllowVolumeExpansion),
			formatDuration(time.Since(storageClass.CreationTimestamp.Time)),
		})

	}
	return storageClassList, nil
}

// DeleteStorageClass  delete the StorageClass and returns the status of deletion
// Parameters:
// - storageClassName : the name of clusterRoleBinding we need to delete
// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - if deletion succeeds returns nil, otherwise returns the error occured
func DeleteStorageClass(storageClassName string,clientSet *client.K8sConfig) error {

	err := clientSet.Client.StorageV1().StorageClasses().Delete(context.TODO(), storageClassName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
