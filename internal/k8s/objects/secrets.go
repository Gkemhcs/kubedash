package objects 
import (
    "context"
	"time"
   "fmt"
   
	client "github.com/Gkemhcs/kubedash/internal/k8s"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
   
//	"gopkg.in/yaml.v3"
)


func ListSecrets(namespace string , clientSet *client.K8sConfig)([][]string,error){
	
	
	secrets,err:=clientSet.Client.CoreV1().Secrets(namespace).List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		return nil,err
	}
	var secretList [][]string
	for _,secret := range secrets.Items {
	  secretList=append(secretList,[]string{
		secret.Name,
		fmt.Sprintf("%v",secret.Type),
		fmt.Sprintf("%d",len(secret.Data)),
		formatDuration(time.Since(secret.CreationTimestamp.Time)),	
	})
	}
	return secretList,nil 
	
}

func DeleteSecret(secretName string,namespace string, clientSet *client.K8sConfig)(error){
	if namespace==""{
	  namespace=clientSet.DefaultNamespace
	}
	err:=clientSet.Client.CoreV1().Secrets(namespace).Delete(context.TODO(),secretName,metav1.DeleteOptions{})
	if err!=nil {
	  return err 
	}
  
  return nil 
  }
