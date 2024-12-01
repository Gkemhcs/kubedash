package objects

import (
	"context"
	"time"

	client "github.com/Gkemhcs/kubedash/internal/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"bytes"
	"text/template"
)

// ListClusterRoles list out the clusterroles in cluster and returns it
// parameters:
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of clusterroles
// - error : if any error occurs returns that otherwise returns nil
func ListClusterRoles(clientSet *client.K8sConfig) ([][]string, error) {

	clusterRole, err := clientSet.Client.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var clusterRoleList [][]string
	for _, clusterRole := range clusterRole.Items {
		clusterRoleList = append(clusterRoleList, []string{
			clusterRole.Name,

			formatDuration(time.Since(clusterRole.CreationTimestamp.Time)),
		})
	}
	return clusterRoleList, nil

}

// DescribeClusterRole  returns the description of clusterrolebindings resource
// Parameters:
// - clusterRoleName : the name of clusterRole we need to describe
// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - description of clusterrole as a buffer of bytes
// - err will be returned if anything occurs ,otherwise returned nil
func DescribeClusterRole(clusterRoleName string, clientSet *client.K8sConfig) (bytes.Buffer, error) {

	clusterRole, err := clientSet.Client.RbacV1().ClusterRoles().Get(context.TODO(), clusterRoleName, metav1.GetOptions{})
	if err != nil {
		return bytes.Buffer{}, err
	}
	const clusterRoleTemplate = `
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
	
	PolicyRule:
	{{- if .Rules }}
    {{ printf "%-40s %-25s %-25s %-30s" "Resources" "Non-Resource URLs" "Resource Names" "Verbs" }}
    {{ printf "%-40s %-25s %-25s %-30s" "---------" "-----------------" "--------------" "-----" }}
    {{- range .Rules }}
        {{ printf "%-40s %-25s %-25s %-30s" .Resources .NonResourceURLs .ResourceNames .Verbs }}
    {{- end }}
	{{- else }}
		<none>
	{{- end }}
	`
	tmpl, err := template.New("describe").Parse(clusterRoleTemplate)
	if err != nil {
		return bytes.Buffer{}, err
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, clusterRole)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return output, nil
}

// DeleteClusterRole  delete the ClusterRole and returns the status of deletion
// Parameters:
// - clusterRoleName : the name of clusterRole we need to delete
// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - if deletion succeeds returns nil, otherwise returns the error occured
func DeleteClusterRole(clusterRoleName string, clientSet *client.K8sConfig) error {

	err := clientSet.Client.RbacV1().ClusterRoles().Delete(context.TODO(), clusterRoleName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
