package objects

import (
	"context"
	"fmt"

	client "github.com/Gkemhcs/kubedash/internal/k8s"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListCronJobs  list out the clusterrolebings in cluster and returns it
// parameters:
// - namespace(string):  the namespace to which we need to scope  our search
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of cronjobs
// - error : if any error occurs returns that otherwise returns nil
func ListCronJobs(namespace string, clientSet *client.K8sConfig) ([][]string, error) {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}

	cronJobs, err := clientSet.Client.BatchV1().CronJobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Print(err)
		return [][]string{}, err
	}
	var cronJobsList [][]string
	for _, cronJob := range cronJobs.Items {

		cronJobsList = append(cronJobsList, []string{
			cronJob.Name,
			cronJob.Spec.Schedule,
			fmt.Sprintf("%b", cronJob.Spec.Suspend),
			fmt.Sprintf("%d", len(cronJob.Status.Active)),
			formatDuration(time.Since(cronJob.Status.LastScheduleTime.Time)),
			formatDuration(time.Since(cronJob.CreationTimestamp.Time)),
		})

	}
	return cronJobsList, nil
}

// DescribeCronJob  returns the description of clusterrolebindings resource
// Parameters:
// - cronJobName : the name of cronJob we need to describe
// - namespace: tha namespace to which we need to scope our search
// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - description of cronJob as a buffer of bytes
// - err will be returned if anything occurs ,otherwise returned nil

// DeleteCronJob  delete the CronJob and returns the status of deletion
// Parameters:
// - cronJobName : the name of clusterRoleBinding we need to delete
// - namespace: tha namespace to which we need to scope our search
// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - if deletion succeeds returns nil, otherwise returns the error occured
func DeleteCronJob(cronJobName string, namespace string, clientSet *client.K8sConfig) error {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}
	err := clientSet.Client.BatchV1().CronJobs(namespace).Delete(context.TODO(), cronJobName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
