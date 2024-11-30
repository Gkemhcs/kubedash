package objects

import (
	"context"
	"fmt"

	"time"
	client "github.com/Gkemhcs/kubedash/internal/k8s"
	

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)







func  ListStorageClasses(namespace string, clientSet *client.K8sConfig)([][]string,error) {


storageClassments,err:=clientSet.Client.StorageV1().StorageClasses().List(context.TODO(),metav1.ListOptions{})
if(err != nil ){
fmt.Print(err)
return [][]string{},err
}
var storageClassList [][]string
for _,storageClass := range storageClassments.Items {
	
	
 	storageClassList=append(storageClassList,[]string{
	storageClass.Name,
	storageClass.Provisioner,
	string(*storageClass.ReclaimPolicy),
	string(*storageClass.VolumeBindingMode),
	fmt.Sprintf("%v",*storageClass.AllowVolumeExpansion),
	formatDuration(time.Since(storageClass.CreationTimestamp.Time)),		
	})
	
}
return storageClassList,nil
}

func DeleteStorageClass(storageClassName string,namespace string, clientSet *client.K8sConfig)(error){
	
	err:=clientSet.Client.StorageV1().StorageClasses().Delete(context.TODO(),storageClassName,metav1.DeleteOptions{})
	if err!=nil {
	  return err 
	}
  
  return nil 
  }
