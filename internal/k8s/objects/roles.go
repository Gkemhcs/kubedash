package objects 
import (
    "context"
	"time"
   
   
	client "github.com/Gkemhcs/kubedash/internal/k8s"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
   
//	"gopkg.in/yaml.v3"
)


func ListRoles(namespace string , clientSet *client.K8sConfig)([][]string,error){
	if namespace == ""{
		namespace=clientSet.DefaultNamespace
	}
	
	roles,err:=clientSet.Client.RbacV1().Roles(namespace).List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		return nil,err
	}
	var roleList [][]string
	for _,role := range roles.Items {
	  roleList=append(roleList,[]string{
		role.Name,
		
		formatDuration(time.Since(role.CreationTimestamp.Time)),	
	})
	}
	return roleList,nil 
	
}

func DeleteRole(roleName string,namespace string, clientSet *client.K8sConfig)(error){
	if namespace==""{
	  namespace=clientSet.DefaultNamespace
	}
	err:=clientSet.Client.RbacV1().Roles(namespace).Delete(context.TODO(),roleName,metav1.DeleteOptions{})
	if err!=nil {
	  return err 
	}
  
  return nil 
  }