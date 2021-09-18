package v1beta1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_ManagedClusterSet = map[string]string{
	"":       "ManagedClusterSet defines a group of ManagedClusters that user's workload can run on. A workload can be defined to deployed on a ManagedClusterSet, which mean:\n  1. The workload can run on any ManagedCluster in the ManagedClusterSet\n  2. The workload cannot run on any ManagedCluster outside the ManagedClusterSet\n  3. The service exposed by the workload can be shared in any ManagedCluster in the ManagedClusterSet\n\nIn order to assign a ManagedCluster to a certian ManagedClusterSet, add a label with name `cluster.open-cluster-management.io/clusterset` on the ManagedCluster to refers to the ManagedClusterSet. User is not allow to add/remove this label on a ManagedCluster unless they have a RBAC rule to CREATE on a virtual subresource of managedclustersets/join. In order to update this label, user must have the permission on both the old and new ManagedClusterSet.",
	"spec":   "Spec defines the attributes of the ManagedClusterSet",
	"status": "Status represents the current status of the ManagedClusterSet",
}

func (ManagedClusterSet) SwaggerDoc() map[string]string {
	return map_ManagedClusterSet
}

var map_ManagedClusterSetBinding = map[string]string{
	"":     "ManagedClusterSetBinding projects a ManagedClusterSet into a certain namespace. User is able to create a ManagedClusterSetBinding in a namespace and bind it to a ManagedClusterSet if they have an RBAC rule to CREATE on the virtual subresource of managedclustersets/bind. Workloads created in the same namespace can only be distributed to ManagedClusters in ManagedClusterSets bound in this namespace by higher level controllers.",
	"spec": "Spec defines the attributes of ManagedClusterSetBinding.",
}

func (ManagedClusterSetBinding) SwaggerDoc() map[string]string {
	return map_ManagedClusterSetBinding
}

var map_ManagedClusterSetBindingList = map[string]string{
	"":         "ManagedClusterSetBindingList is a collection of ManagedClusterSetBinding.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of ManagedClusterSetBinding.",
}

func (ManagedClusterSetBindingList) SwaggerDoc() map[string]string {
	return map_ManagedClusterSetBindingList
}

var map_ManagedClusterSetBindingSpec = map[string]string{
	"":           "ManagedClusterSetBindingSpec defines the attributes of ManagedClusterSetBinding.",
	"clusterSet": "ClusterSet is the name of the ManagedClusterSet to bind. It must match the instance name of the ManagedClusterSetBinding and cannot change once created. User is allowed to set this field if they have an RBAC rule to CREATE on the virtual subresource of managedclustersets/bind.",
}

func (ManagedClusterSetBindingSpec) SwaggerDoc() map[string]string {
	return map_ManagedClusterSetBindingSpec
}

var map_ManagedClusterSetList = map[string]string{
	"":         "ManagedClusterSetList is a collection of ManagedClusterSet.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of ManagedClusterSet.",
}

func (ManagedClusterSetList) SwaggerDoc() map[string]string {
	return map_ManagedClusterSetList
}

var map_ManagedClusterSetSpec = map[string]string{
	"": "ManagedClusterSetSpec describes the attributes of the ManagedClusterSet",
}

func (ManagedClusterSetSpec) SwaggerDoc() map[string]string {
	return map_ManagedClusterSetSpec
}

var map_ManagedClusterSetStatus = map[string]string{
	"":           "ManagedClusterSetStatus represents the current status of the ManagedClusterSet.",
	"conditions": "Conditions contains the different condition statuses for this ManagedClusterSet.",
}

func (ManagedClusterSetStatus) SwaggerDoc() map[string]string {
	return map_ManagedClusterSetStatus
}

// AUTO-GENERATED FUNCTIONS END HERE