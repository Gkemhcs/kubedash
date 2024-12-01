package objects

import (
	"context"
	"fmt"
	"time"

	"bytes"
	client "github.com/Gkemhcs/kubedash/internal/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"text/template"
	//	"gopkg.in/yaml.v3"
)


// ListConfigMaps  list out the clusterrolebings in cluster and returns it 
// parameters:
// - namespace(string):  the namespace to which  we need to scope  our search
// - clientSet : the kubernetes client which need to use to fetch the resources
// returns :
// - list of configmaps
// - error : if any error occurs returns that otherwise returns nil 

func ListConfigMaps(namespace string, clientSet *client.K8sConfig) ([][]string, error) {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}

	configmaps, err := clientSet.Client.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var configMapList [][]string
	for _, configMap := range configmaps.Items {
		configMapList = append(configMapList, []string{
			configMap.Name,
			fmt.Sprintf("%d", len(configMap.Data)),
			formatDuration(time.Since(configMap.CreationTimestamp.Time)),
		})
	}
	return configMapList, nil

}

func DescribeConfigMap(configMapName string, namespace string, clientSet *client.K8sConfig) (bytes.Buffer, error) {
	configMap, err := clientSet.Client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		return bytes.Buffer{}, err
	}
	configMapTemplate := `
	ConfigMap Name: {{ .ObjectMeta.Name }}
	Namespace: {{ .ObjectMeta.Namespace }}

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

	Data:
	{{- if .Data }}
	{{- range $key, $value := .Data }}
	{{ $key }}:
				{{- $value  }}
	 --------
	{{- end }}
	{{- else }}
	<none>
	{{- end }}

	BinaryData:
	{{- if .BinaryData }}
	{{- range $key, $value := .BinaryData }}
	{{ $key }}: {{ $value }}
	{{- end }}
	{{- else }}
	<none>
	{{- end }}
	`

	tmpl, err := template.New("describe").Parse(configMapTemplate)
	if err != nil {

		return bytes.Buffer{}, err
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, configMap)
	if err != nil {

		return bytes.Buffer{}, err

	}

	return output, nil

}

func DeleteConfigMap(configMapName string, namespace string, clientSet *client.K8sConfig) error {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}
	err := clientSet.Client.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), configMapName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
