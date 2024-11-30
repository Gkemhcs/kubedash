package objects 
import (
    "context"
	"time"
   
	"text/template"
	"bytes"
	client "github.com/Gkemhcs/kubedash/internal/k8s"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
   
//	"gopkg.in/yaml.v3"
)


func ListClusterRoleBindings(namespace string , clientSet *client.K8sConfig)([][]string,error){
	

	clusterRoleBindings,err:=clientSet.Client.RbacV1().ClusterRoleBindings().List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		return nil,err
	}
	var clusterRoleBindingList [][]string
	for _,clusterRoleBinding := range clusterRoleBindings.Items {
	  clusterRoleBindingList=append(clusterRoleBindingList,[]string{
		clusterRoleBinding.Name,
		clusterRoleBinding.RoleRef.Name,
		clusterRoleBinding.Subjects[0].Name,
		clusterRoleBinding.Subjects[0].Kind,
		
		formatDuration(time.Since(clusterRoleBinding.CreationTimestamp.Time)),	
	})
	}
	return clusterRoleBindingList,nil 
	
}

func DescribeClusterRoleBinding(clusterRoleBindingName string,namespace string , clientSet *client.K8sConfig)(bytes.Buffer,error){
	

	clusterRoleBinding,err:=clientSet.Client.RbacV1().ClusterRoleBindings().Get(context.Background(),clusterRoleBindingName,metav1.GetOptions{})
	if err != nil {
		return bytes.Buffer{},err
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
	err = tmpl.Execute(&output,clusterRoleBinding)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return output, nil

}