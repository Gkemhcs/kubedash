package objects

import (
	"context"
	"fmt"
	"time"

	client "github.com/Gkemhcs/kubedash/internal/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//	"gopkg.in/yaml.v3"
)

// ListPersistentVolumes  list out the clusterrolebings in cluster and returns it
// parameters:
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of peristentvolumes
// - error : if any error occurs returns that otherwise returns nil
func ListPersistentVolumes( clientSet *client.K8sConfig) ([][]string, error) {

	persistentVolumes, err := clientSet.Client.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var persistentVolumesList [][]string
	for _, persistentVolume := range persistentVolumes.Items {
		persistentVolumesList = append(persistentVolumesList, []string{
			persistentVolume.Name,
			fmt.Sprintf("%v", *persistentVolume.Spec.Capacity.Storage()),
			fmt.Sprintf("%v", persistentVolume.Spec.AccessModes),
			string(persistentVolume.Spec.PersistentVolumeReclaimPolicy),
			string(persistentVolume.Status.Phase),
			fmt.Sprintf("%s/%s", persistentVolume.Spec.ClaimRef.Namespace, persistentVolume.Spec.ClaimRef.Name),
			persistentVolume.Spec.StorageClassName,

			fmt.Sprintf(persistentVolume.Status.Reason),

			formatDuration(time.Since(persistentVolume.CreationTimestamp.Time)),
		})
	}
	return persistentVolumesList, nil

}

// DeletePersistentVolume  delete the PersistentVolume and returns the status of deletion
// Parameters:
// - persistentVolumeName : the name of clusterRoleBinding we need to delete
// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - if deletion succeeds returns nil, otherwise returns the error occured
func DeletePersistentVolume(volumeName string,clientSet *client.K8sConfig) error {

	err := clientSet.Client.CoreV1().PersistentVolumes().Delete(context.TODO(), volumeName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
