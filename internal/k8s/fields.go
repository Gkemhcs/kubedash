package k8s

// GetPodFields returns the headers  for  table of pods
// Returns:
// - returns the array of  headers
func GetPodFields() []string {
	return []string{"Name", "Ready", "Status", "Restarts", "IP", "Age", "Node Name"}
}

// GetDeploymentFields returns the headers  for  table of deployments
// Returns:
// - returns the array of  headers
func GetDeploymentFields() []string {
	return []string{"Name", "Available", "Updated", "Ready", "Age"}
}

// GetConfigMapFields returns the headers  for  table of configmaps
// Returns:
// - returns the array of  headers
func GetConfigMapFields() []string {
	return []string{"Name", "Data", "Age"}
}

// GetSecretFields returns the headers  for  table of secrets
// Returns:
// - returns the array of  headers
func GetSecretFields() []string {
	return []string{"Name", "Type", "Data", "Age"}
}

// GetDaemonSetFields returns the headers  for  table of daemonSets
// Returns:
// - returns the array of  headers
func GetDaemonSetFields() []string {
	return []string{"Name", "Desired", "Current", "Ready", "Up-To-Date", "Age"}

}

// GetReplicaSetFields returns the headers  for  table of replicaSets
// Returns:
// - returns the array of  headers
func GetReplicaSetFields() []string {
	return []string{"Name", "Ready", "Available", "Total", "Age"}
}

// GetServiceFields returns the headers  for  table of  services
// Returns:
// - returns the array of  headers
func GetServiceFields() []string {
	return []string{"Name", "Type", "Cluster-IP", "LoadBalancer-IP", "Ports", "Age"}
}

// GetStorageClassFields returns the headers  for  table of storageClasses
// Returns:
// - returns the array of  headers
func GetStorageClassFields() []string {
	return []string{"Name", "Provisoner", "ReclaimPolicy", "VolumeBindingMode", "AllowVolumeExpansion", "Age"}
}

// GetJobFields returns the headers  for  table of jobs
// Returns:
// - returns the array of  headers
func GetJobFields() []string {
	return []string{"Name", "Completions", "Duration", "Age"}
}

// GetCronJobFields returns the headers  for  table of cronjobs
// Returns:
// - returns the array of  headers
func GetCronJobFields() []string {
	return []string{"Name", "Schedule", "Suspend", "Active", "Last-Schedule", "Age"}
}

// GetIngressFields returns the headers  for  table of ingresses
// Returns:
// - returns the array of  headers
func GetIngressFields() []string {
	return []string{"Name", "Class", "Hosts", "Address", "Ports", "Age"}
}

// GetPersistentVolumeFields returns the headers  for  table of persistentVolumes
// Returns:
// - returns the array of  headers
func GetPersistentVolumeFields() []string {
	return []string{"Name", "Capacity", "AccessMode", "ReclaimPolicy", "Status", "Claim", "StorageClass", "Reason", "Age"}
}

// GetServiceAccountFields returns the headers  for  table of serviceAccounts
// Returns:
// - returns the array of  headers
func GetServiceAccountFields() []string {
	return []string{"Name", "Age"}
}

// GetRoleFields returns the headers  for  table of roles
// Returns:
// - returns the array of  headers
func GetRoleFields() []string {
	return []string{"Name", "Age"}
}

// GetRoleBindingFields returns the headers  for  table of  roleBindings
// Returns:
// - returns the array of  headers
func GetRoleBindingFields() []string {
	return []string{"Name", "Role", "Kind", "Subject", "Age"}
}

// GetClusterRoleFields returns the headers  for  table of clusterRoles
// Returns:
// - returns the array of  headers
func GetClusterRoleFields() []string {
	return []string{"Name", "Age"}
}

// GetClusterRoleBindingFields returns the headers  for  table of clusterRoleBindings
// Returns:
// - returns the array of  headers
func GetClusterRoleBindingFields() []string {
	return []string{"Name", "Role", "Kind", "Subject", "Age"}
}

// GetEndpointFields returns the headers  for  table of endpoints
// Returns:
// - returns the array of  headers
func GetEndpointFields() []string {
	return []string{"Name", "Addresses", "Age"}
}

// GetNodeFields returns the headers  for  table of nodes
// Returns:
// - returns the array of  headers
func GetNodeFields() []string {
	return []string{"Name", "Status", "Taints", "Version", "Cpu", "Memory", "CpuAllocated", "MemoryAllocated"}
}
