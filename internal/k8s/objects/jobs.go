package objects

import (
	"context"
	"fmt"

	client "github.com/Gkemhcs/kubedash/internal/k8s"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListJobs  list out the clusterrolebings in cluster and returns it
// parameters:
// - namespace(string):  the namespace to which  we need to scope  our search
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of configmaps
// - error : if any error occurs returns that otherwise returns nil
func ListJobs(namespace string, clientSet *client.K8sConfig) ([][]string, error) {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}

	jobs, err := clientSet.Client.BatchV1().Jobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Print(err)
		return [][]string{}, err
	}
	var jobsList [][]string
	for _, job := range jobs.Items {

		jobsList = append(jobsList, []string{
			job.Name,
			fmt.Sprintf("%d", job.Spec.Completions),
			formatDuration(time.Since(job.Status.CompletionTime.Time)),
			formatDuration(time.Since(job.CreationTimestamp.Time)),
		})

	}
	return jobsList, nil
}

// DeleteJob  delete the Job and returns the status of deletion
// Parameters:
// - jobName : the name of clusterRoleBinding we need to delete
// - namespace: tha namespace to which we need to scope our search
// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - if deletion succeeds returns nil, otherwise returns the error occured
func DeleteJob(jobName string, namespace string, clientSet *client.K8sConfig) error {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}
	err := clientSet.Client.BatchV1().Jobs(namespace).Delete(context.TODO(), jobName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
