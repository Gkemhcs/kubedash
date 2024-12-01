package objects

import (
	"context"
	"time"

	"bytes"
	client "github.com/Gkemhcs/kubedash/internal/k8s"
	"text/template"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//	"gopkg.in/yaml.v3"
)

// ListClusterRoleBindings  list out the clusterrolebindings in cluster and returns it
// parameters:
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of clusterrolebindings
// - error : if any error occurs returns that otherwise returns nil
func ListClusterRoleBindings(clientSet *client.K8sConfig) ([][]string, error) {

	clusterRoleBindings, err := clientSet.Client.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var clusterRoleBindingList [][]string
	for _, clusterRoleBinding := range clusterRoleBindings.Items {
		clusterRoleBindingList = append(clusterRoleBindingList, []string{
			clusterRoleBinding.Name,
			clusterRoleBinding.RoleRef.Name,
			clusterRoleBinding.Subjects[0].Name,
			clusterRoleBinding.Subjects[0].Kind,

			formatDuration(time.Since(clusterRoleBinding.CreationTimestamp.Time)),
		})
	}
	return clusterRoleBindingList, nil

}

// DescribeClusterRoleBinding  returns the description of clusterrolebindings resource
// Parameters:
// - clusterRoleBindingName : the name of clusterRoleBinding we need to describe
// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - description of clusterrolebinding as a buffer of bytes
// - err will be returned if anything occurs ,otherwise returned nil
func DescribeClusterRoleBinding(clusterRoleBindingName string, clientSet *client.K8sConfig) (bytes.Buffer, error) {

	clusterRoleBinding, err := clientSet.Client.RbacV1().ClusterRoleBindings().Get(context.Background(), clusterRoleBindingName, metav1.GetOptions{})
	if err != nil {
		return bytes.Buffer{}, err
	}
	const clusterRoleBindingTemplate = `
	Name:         {{ .ObjectMeta.Name }}
	Namespace:    {{ .ObjectMeta.Namespace }}
	Labels:
	{{- if .ObjectMeta.Labels }}
	{{- range $key, $value := .ObjectMeta.Labels }}
	{{ $key }}: {{ $value }}
	{{- end }}
	{{- else }}
	<none>
	{{- end }}
	Annotations:
	{{- if .ObjectMeta.Annotations }}
	{{- range $key, $value := .ObjectMeta.Annotations }}
	{{ $key }}: {{ $value }}
	{{- end }}
	{{- else }}
	<none>
	{{- end }}

	Role:
	Kind:  {{ .RoleRef.Kind }}
	Name:  {{ .RoleRef.Name }}

	Subjects:
	{{- if .Subjects }}
	{{- printf "%-15s %-30s %-15s" "Kind" "Name" "Namespace" }}
	{{- printf "\n%-15s %-30s %-15s" "----" "----" "---------" }}
	{{- range .Subjects }}
	{{ printf "\n%-15s %-30s %-15s" .Kind .Name (default "<none>" .Namespace) }}
	{{- end }}
	{{- else }}
	<none>
	{{- end }}
	`
	tmpl, err := template.New("describe").Parse(clusterRoleBindingTemplate)
	if err != nil {

		return bytes.Buffer{}, err
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, clusterRoleBinding)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return output, nil

}

// DeleteClusterRoleBinding  delete the ClusterRoleBinding and returns the status of deletion
// Parameters:
// - clusterRoleBindingName : the name of clusterRoleBinding we need to delete
// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - if deletion succeeds returns nil, otherwise returns the error occured
func DeleteClusterRoleBinding(clusterRoleBindingName string, clientSet *client.K8sConfig) error {

	err := clientSet.Client.RbacV1().ClusterRoleBindings().Delete(context.TODO(), clusterRoleBindingName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
