package v1alpha1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	configapi "sigs.k8s.io/kueue/apis/config/v1beta1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Kueue holds cluster-wide information about the Kueue Operator.
//
// Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
// +openshift:compatibility-gen:level=4
// +openshift:file-pattern=cvoRunLevel=0000_00,operatorName=kueue-operator,operatorOrdering=01
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=kueue,scope=Cluster
// +openshift:api-approved.openshift.io=https://github.com/openshift/api/pull/2044
// +openshift:enable:FeatureGate=ClusterVersionOperatorConfiguration
// +kubebuilder:validation:XValidation:rule="self.metadata.name == 'cluster'",message="KueueOperator is a singleton; the .metadata.name field must be 'cluster'"

type Kueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec holds user settable values for configuration
	// +required
	Spec KueueOperandSpec `json:"spec"`
	// status holds observed values from the cluster. They may not be overridden.
	// +optional
	Status KueueStatus `json:"status"`
}

type KueueOperandSpec struct {
	// The config that is persisted to a config map
	Config KueueConfiguration `json:"config"`
	// Image
	Image string `json:"image"`
}

type KueueConfiguration struct {
	// WaitForPodsReady configures gang admission
	// optional
	WaitForPodsReady *configapi.WaitForPodsReady `json:"waitForPodsReady,omitempty"`
	// Integrations are the types of integrations Kueue will manager
	// Required
	Integrations configapi.Integrations `json:"integrations"`
	// Feature gates are advanced features for Kueue
	FeatureGates map[string]bool `json:"featureGates,omitempty"`
}

// KueueStatus defines the observed state of Kueue
type KueueStatus struct {
	operatorv1.OperatorStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KueueList contains a list of SecondaryScheduler
type SecondarySchedulerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kueue `json:"items"`
}
