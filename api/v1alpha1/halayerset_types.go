/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HALayerSetSpec defines the desired state of HALayerSet
type HALayerSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	//SNO1IP is the IP of the first single node openshift cluster
	SNO1IP string `json:"sno1ip"`

	//SNO2IP is the IP of the second single node openshift cluster
	SNO2IP string `json:"sno2ip"`

	//KubeConfigPath1 is the path to the kubeconfig of the first single node openshift cluster
	KubeConfigPath1 string `json:"kubeconfigpath1"`

	//KubeConfigPath2 is the path to the kubeconfig of the second single node openshift cluster
	KubeConfigPath2 string `json:"kubeconfigpath2"`
}

// HALayerSetStatus defines the observed state of HALayerSet
type HALayerSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HALayerSet is the Schema for the halayersets API
type HALayerSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HALayerSetSpec   `json:"spec,omitempty"`
	Status HALayerSetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HALayerSetList contains a list of HALayerSet
type HALayerSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HALayerSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HALayerSet{}, &HALayerSetList{})
}
