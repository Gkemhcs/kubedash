package objects

import (
	"bytes"
	"context"
	"fmt"
	"text/template"
	"time"

	client "github.com/Gkemhcs/kubedash/internal/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListDeployments  list out the clusterrolebings in cluster and returns it
// parameters:
// - namespace(string):  the namespace to which  we need to scope  our search
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of deployments
// - error : if any error occurs returns that otherwise returns nil
func ListDeployments(namespace string, clientSet *client.K8sConfig) ([][]string, error) {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}

	deployments, err := clientSet.Client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Print(err)
		return [][]string{}, err
	}
	var deploymentList [][]string
	for _, deploy := range deployments.Items {

		deploymentList = append(deploymentList, []string{
			deploy.Name,
			fmt.Sprintf("%d", deploy.Status.AvailableReplicas),
			fmt.Sprintf("%d", deploy.Status.UpdatedReplicas),
			fmt.Sprintf("%d/%d", deploy.Status.ReadyReplicas, deploy.Status.Replicas),
			formatDuration(time.Since(deploy.CreationTimestamp.Time)),
		})

	}
	return deploymentList, nil
}

// DescribeDeployment  returns the description of clusterrolebindings resource
// Parameters:
// - deploymentName : the name of deployment we need to describe
// - namespace: tha namespace to which we need to scope our search
// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - description of deployment as a buffer of bytes
// - err will be returned if anything occurs ,otherwise returned nil
func DescribeDeployment(deploymentName string, namespace string, clientSet *client.K8sConfig) (bytes.Buffer, error) {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}

	deployment, err := clientSet.Client.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {

		return bytes.Buffer{}, err
	}
	deploymentTemplate := `

Deployment Name: {{ .ObjectMeta.Name }}
	Namespace: {{ .ObjectMeta.Namespace }}
	Creation Timestamp: {{ .ObjectMeta.CreationTimestamp }}
	
	Labels:
	{{- if .ObjectMeta.Labels }}
	  {{- range $key, $value := .ObjectMeta.Labels }}
	  {{ $key }}: {{ $value }}
	  {{- end }}
	{{- else }}
	  None
	{{- end }}
	Selector:
		{{- if .Spec.Selector.MatchLabels }}
	  		{{- range $key, $value := .Spec.Selector.MatchLabels }}
	  			{{ $key }}: {{ $value }}
	  		{{- end }}
		{{- else }}
				None
		{{- end }}	
	Annotations:
	{{- if .ObjectMeta.Annotations }}
	  {{- range $key, $value := .ObjectMeta.Annotations }}
	  {{ $key }}: {{ $value }}
	  {{- end }}
	{{- else }}
	  None
	{{- end }}
	
		
	Replicas: 
	  Desired: {{ .Spec.Replicas }}
	  Updated: {{ .Status.UpdatedReplicas }}
	  Total: {{ .Status.Replicas }}
	  Available: {{ .Status.AvailableReplicas }}
	  Unavailable: {{ .Status.UnavailableReplicas }}
	
	Strategy Type: {{ .Spec.Strategy.Type }}
	Min Ready Seconds: {{ .Spec.MinReadySeconds }}
	
	Pod Template:
	  Labels:
	  {{- range $key, $value := .Spec.Template.Labels }}
		{{ $key }}: {{ $value }}
	  {{- end }}
	
	  Containers:
	  {{- range .Spec.Template.Spec.Containers }}
		{{ .Name }}:
		  Image: {{ .Image }}
		  Ports:
		  {{- if .Ports }}
			{{- range .Ports }}
			- {{ .ContainerPort }}/{{ .Protocol }}
			{{- end }}
		  {{- else }}
			None
		  {{- end }}
		  
		  Limits:
		  {{- if .Resources.Limits }}
			CPU: {{ .Resources.Limits.Cpu }}
			Memory: {{ .Resources.Limits.Memory }}
		  {{- else }}
			None
		  {{- end }}
	
		  Requests:
		  {{- if .Resources.Requests }}
			CPU: {{ .Resources.Requests.Cpu }}
			Memory: {{ .Resources.Requests.Memory }}
		  {{- else }}
			None
		  {{- end }}
	
		  
	
	  {{- end }}
	
	Volumes:
	{{- range .Spec.Template.Spec.Volumes }}
	  - Name: {{ .Name }}
		{{- if .EmptyDir }}
		  Type: EmptyDir
		{{- else if .Secret }}
		  Type: Secret
		  Secret Name: {{ .Secret.SecretName }}
		{{- else }}
		  Type: Unknown
		{{- end }}
	{{- end }}
	
	Node Selector:
	{{- if .Spec.Template.Spec.NodeSelector }}
	  {{- range $key, $value := .Spec.Template.Spec.NodeSelector }}
	  {{ $key }}: {{ $value }}
	  {{- end }}
	{{- else }}
	  None
	{{- end }}
	
	Tolerations:
	{{- range .Spec.Template.Spec.Tolerations }}
	  - Key: {{ .Key }}
		Operator: {{ .Operator }}
		Value: {{ .Value }}
		Effect: {{ .Effect }}
	{{- else }}
	  None
	{{- end }}
	
	Conditions:
	{{- range .Status.Conditions }}
	  Type: {{ .Type }}
	  Status: {{ .Status }}
	  Reason: {{ .Reason }}
	{{- end }}
	`
	tmpl, err := template.New("describe").Parse(deploymentTemplate)
	if err != nil {

		return bytes.Buffer{}, err
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, deployment)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return output, nil

}

// DeleteDeployment  delete the Deployment and returns the status of deletion
// Parameters:
// - deploymentName : the name of clusterRoleBinding we need to delete
// - namespace: tha namespace to which we need to scope our search
// - clientSet: the  k8sclient need to use to fetch the resources
// Returns:
// - if deletion succeeds returns nil, otherwise returns the error occured
func DeleteDeployment(deploymentName string, namespace string, clientSet *client.K8sConfig) error {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}
	err := clientSet.Client.AppsV1().Deployments(namespace).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
