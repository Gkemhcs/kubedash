package objects

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"text/template"
	"time"

	client "github.com/Gkemhcs/kubedash/internal/k8s"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type podMetadata struct {
	Name            string                  `json:"name"`
	Namespace       string                  `json:"namespace"`
	Labels          map[string]string       `json:"labels"`
	Annotations     map[string]string       `json:"annotations"`
	OwnerReferences []metav1.OwnerReference `json:"ownerReferences"`
}
type podSpec struct {
	Node           string              `json:"nodeName"`
	Containers     []corev1.Container  `json:"containers"`
	Volumes        []corev1.Volume     `json:"volumes"`
	ServiceAccount string              `json:"serviceAccountName"`
	NodeSelector   corev1.NodeSelector `json:"nodeSelector"`
	Tolerations    []corev1.Toleration `json:"tolerations"`
}
type podStatus struct {
	Phase      string                `json:"phase"`
	Ip         string                `json:"podIP"`
	Conditions []corev1.PodCondition `json:"conditions"`
	QOSClass   corev1.PodQOSClass    `json:"qosClass"`
}
type Pod struct {
	Metadata podMetadata `json:"metadata"`
	Spec     podSpec     `json:"spec"`
	Status   podStatus   `json:"status"`
}

func formatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	} else if d < time.Hour {
		return fmt.Sprintf("%dm", int(d.Minutes()))
	} else if d < 24*time.Hour {
		return fmt.Sprintf("%dh", int(d.Hours()))
	}
	return fmt.Sprintf("%dd", int(d.Hours()/24))
}

func ListPods(namespace string, clientSet *client.K8sConfig) ([][]string, error) {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}

	pods, err := clientSet.Client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Print(err)
		return [][]string{}, err
	}
	var podList [][]string
	for _, pod := range pods.Items {

		var restartCount int32
		var readyCount int64
		for _, container := range pod.Status.ContainerStatuses {
			restartCount += container.RestartCount
			if container.Ready {
				readyCount += 1
			}
		}

		podList = append(podList, []string{
			pod.Name,
			fmt.Sprintf("%d/%d", readyCount, len(pod.Status.ContainerStatuses)),
			string(pod.Status.Phase),
			fmt.Sprintf("%d", restartCount),
			pod.Status.PodIP,
			formatDuration(time.Since(pod.Status.StartTime.Time)),
			pod.Spec.NodeName,
		})

	}
	return podList, nil
}

func DescribePod(podName string, namespace string, clientSet *client.K8sConfig) (bytes.Buffer, error) {

	resp, err := clientSet.Client.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {

		return bytes.Buffer{}, err
	}
	var podDetails Pod
	jsonData, err := json.Marshal(resp)
	if err != nil {

		return bytes.Buffer{}, err
	}
	if err := json.Unmarshal(jsonData, &podDetails); err != nil {

		return bytes.Buffer{}, err
	}
	podTemplate := `
    Pod Name: {{.Metadata.Name }}
    Namespace: {{ .Metadata.Namespace }}
	{{- if .Spec.ServiceAccount }}
    Service Account: {{ .Spec.ServiceAccount }}
	{{- end }}
    Node: {{ .Spec.Node }}
    
    Annotations:
    {{- if .Metadata.Annotations }}
      {{- range $key, $value := .Metadata.Annotations }}
        {{ $key }}: {{ $value }}
      {{- end }}
    {{- else }}
      None
    {{- end }}
    
    Labels:
    {{- if .Metadata.Labels }}
      {{- range $key, $value := .Metadata.Labels }}
        {{ $key }}: {{ $value }}
      {{- end }}
    {{- else }}
      None
    {{- end }}
    
    Containers:
    {{- range .Spec.Containers }}
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
        
        Command: 
        {{- if .Command }}
          {{- range .Command }}
            - {{ . }}
          {{- end }}
        {{- else }}
          None
        {{- end }}
    
        Args:
        {{- if .Args }}
          {{- range .Args }}
            - {{ . }}
          {{- end }}
        {{- else }}
          None
        {{- end }}
    {{- end }}
    Volumes:
        {{- range .Spec.Volumes  }}

        {{- end }}
    QOSClass: {{ .Status.QOSClass }}
`

	tmpl, err := template.New("describe").Parse(podTemplate)
	if err != nil {

		return bytes.Buffer{}, err
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, podDetails)
	if err != nil {

		return bytes.Buffer{}, err

	}

	return output, nil

}

func DeletePod(podName string, namespace string, clientSet *client.K8sConfig) error {
	if namespace == "" {
		namespace = clientSet.DefaultNamespace
	}
	err := clientSet.Client.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
