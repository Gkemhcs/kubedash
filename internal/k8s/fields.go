package k8s 


func GetPodFields() []string {
	return []string{"Name", "Ready", "Status", "Restarts", "IP", "Age", "Node Name"}
}
func GetDeploymentFields() []string{
	return []string {"Name" ,"Available", "Updated","Ready","Age" }
}

func GetConfigMapFields()[]string{
	return []string{"Name","Data","Age"}
}
func GetSecretFields()[]string{
	return []string{"Name","Type","Data","Age"}
}
func GetDaemonSetFields()[]string{
	return  []string{"Name","Desired","Current","Ready","Up-To-Date","Age"}

}
func GetReplicaSetFields()[]string{
	return  []string{"Name","Ready","Available","Total","Age"}
}
func GetServiceFields()[]string{
	return  []string{"Name","Type","Cluster-IP","LoadBalancer-IP","Ports","Age"}
}
func GetStorageClassFields()[]string{
	return []string{"Name","Provisoner","ReclaimPolicy","VolumeBindingMode","AllowVolumeExpansion","Age"}
}

func GetJobFields()[]string{
	return []string {"Name","Completions","Duration","Age"}
}
func GetCronJobFields()[]string{
	return []string{"Name","Schedule","Suspend","Active","Last-Schedule","Age"}
}

func GetIngressFields()[]string{
	return []string{"Name","Class","Hosts","Address","Ports","Age"}
}

func GetPersistentVolumeFields()[]string{
	return []string{"Name","Capacity","AccessMode","ReclaimPolicy","Status","Claim","StorageClass","Reason","Age"}
}
func GetServiceAccountFields()[]string{
	return []string{"Name","Age"}
}
func GetRoleFields()[]string{
	return []string{"Name","Age"}
}
func GetRoleBindingFields()[]string{
	return []string{"Name","Role","Kind","Subject","Age"}
}
func GetClusterRoleFields()[]string{
	return []string{"Name","Age"}
}
func GetClusterRoleBindingFields()[]string{
	return []string{"Name","Role","Kind","Subject","Age"}
}
func GetEndpointFields()[]string{
	return []string{"Name","Addresses","Age"}
}
func GetNodeFields()[]string {
	return []string{"Name","Status","Taints","Version","Cpu","Memory","CpuAllocated","MemoryAllocated"}
}

