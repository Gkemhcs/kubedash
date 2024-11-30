package objects

import (
	"context"
	"fmt"

	"time"
	client "github.com/Gkemhcs/kubedash/internal/k8s"
	

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)








func  ListCronJobs(namespace string, clientSet *client.K8sConfig)([][]string,error) {
if namespace == ""{
	namespace=clientSet.DefaultNamespace
}

cronJobs,err:=clientSet.Client.BatchV1().CronJobs(namespace).List(context.TODO(),metav1.ListOptions{})
if(err != nil ){
fmt.Print(err)
return [][]string{},err
}
var cronJobsList [][]string
for _,cronJob := range cronJobs.Items {
	
	
 	cronJobsList=append(cronJobsList,[]string{
	cronJob.Name,
	cronJob.Spec.Schedule,
	fmt.Sprintf("%b",cronJob.Spec.Suspend),
	fmt.Sprintf("%d",len(cronJob.Status.Active)),
	formatDuration(time.Since(cronJob.Status.LastScheduleTime.Time)),
	formatDuration(time.Since(cronJob.CreationTimestamp.Time)),		
	})
	
}
return cronJobsList,nil
}


func DeleteCronJob(cronJobName string,namespace string, clientSet *client.K8sConfig)(error){
	if namespace==""{
	  namespace=clientSet.DefaultNamespace
	}
	err:=clientSet.Client.BatchV1().CronJobs(namespace).Delete(context.TODO(),cronJobName,metav1.DeleteOptions{})
	if err!=nil {
	  return err 
	}
  
  return nil 
  }

