package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sConfig struct {
	Client           kubernetes.Interface
	DefaultNamespace string
}

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

func (k8sConfig *K8sConfig) GetClient() *K8sConfig {
	return k8sConfig
}

/*
var K8sClient *kubernetes.Clientset

func InitK8sClient(context string) (error) {
    // Use the default kubeconfig file from ~/.kube/config

    kubeconfig := clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()

    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

    if err != nil {
        return  err
    }


    // Create a clientset
    K8sClient, err = kubernetes.NewForConfig(config)
    if err != nil {
        return  err
    }
    return nil


}
*/
