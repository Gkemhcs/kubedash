package objects

import (
	"context"
	"fmt"
	"time"

	client "github.com/Gkemhcs/kubedash/internal/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//	"gopkg.in/yaml.v3"
)

func ListPersistentVolumes(namespace string, clientSet *client.K8sConfig) ([][]string, error) {

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

func DeletePersistentVolume(volumeName string, namespace string, clientSet *client.K8sConfig) error {

	err := clientSet.Client.CoreV1().PersistentVolumes().Delete(context.TODO(), volumeName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
