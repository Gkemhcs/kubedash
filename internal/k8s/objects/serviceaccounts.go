package objects 
import (
    "context"
	"time"

   
	client "github.com/Gkemhcs/kubedash/internal/k8s"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
   
//	"gopkg.in/yaml.v3"
)


func ListServiceAccounts(namespace string , clientSet *client.K8sConfig)([][]string,error){
	if namespace == ""{
	namespace=clientSet.DefaultNamespace
}
	
	serviceAccounts,err:=clientSet.Client.CoreV1().ServiceAccounts(namespace).List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		return nil,err
	}
	var serviceAccountList [][]string
	for _,serviceAccount := range serviceAccounts.Items {
	  serviceAccountList=append(serviceAccountList,[]string{
		serviceAccount.Name,
		
		
		formatDuration(time.Since(serviceAccount.CreationTimestamp.Time)),	
	})
	}
	return serviceAccountList,nil 
	
}
func DeleteServiceAccount(serviceAccountName string,namespace string, clientSet *client.K8sConfig)(error){
	if namespace==""{
	  namespace=clientSet.DefaultNamespace
	}
	err:=clientSet.Client.CoreV1().ServiceAccounts(namespace).Delete(context.TODO(),serviceAccountName,metav1.DeleteOptions{})
	if err!=nil {
	  return err 
	}
  
  return nil 
  }
