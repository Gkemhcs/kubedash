package objects 
import (
    "context"
	"time"
   
   
	client "github.com/Gkemhcs/kubedash/internal/k8s"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
   
//	"gopkg.in/yaml.v3"
)


func ListRoleBindings(namespace string , clientSet *client.K8sConfig)([][]string,error){
	
	if namespace == ""{
		namespace=clientSet.DefaultNamespace
	}
	roleBindings,err:=clientSet.Client.RbacV1().RoleBindings(namespace).List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		return nil,err
	}
	var roleBindingList [][]string
	for _,roleBinding := range roleBindings.Items {
	  roleBindingList=append(roleBindingList,[]string{
		roleBinding.Name,
		roleBinding.RoleRef.Name,
		roleBinding.Subjects[0].Name,
		roleBinding.Subjects[0].Kind,
		
		formatDuration(time.Since(roleBinding.CreationTimestamp.Time)),	
	})
	}
	return roleBindingList,nil 
	
}

func DeleteRoleBinding(roleBindingName string,namespace string, clientSet *client.K8sConfig)(error){
	if namespace==""{
	  namespace=clientSet.DefaultNamespace
	}
	err:=clientSet.Client.RbacV1().RoleBindings(namespace).Delete(context.TODO(),roleBindingName,metav1.DeleteOptions{})
	if err!=nil {
	  return err 
	}
  
  return nil 
  }