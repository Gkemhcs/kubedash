package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// K8sConfig contains k8s clientSet and defaultNamespace
type K8sConfig struct {
	Client           kubernetes.Interface
	DefaultNamespace string
}

// InitClient function  initializes the k8s client
func (k8sConfig *K8sConfig) InitClient() error {
	kubeconfig := clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		return err
	}

	// Create a clientset
	k8sConfig.Client, err = kubernetes.NewForConfig(config)

	if err != nil {
		return err
	}
	k8sConfig.DefaultNamespace = "monitoring"
	return nil

}

// GetClient function returns the k8sConfig pointer
func (k8sConfig *K8sConfig) GetClient() *K8sConfig {
	return k8sConfig
}
