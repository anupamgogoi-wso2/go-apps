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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HelloAppSpec defines the desired state of HelloApp
type HelloAppSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of HelloApp. Edit helloapp_types.go to remove/update
	Name string `json:"name,omitempty"`
	Age  int    `json:"Age,omitempty"`
	Size int32  `json:"size"`
}

// HelloAppStatus defines the observed state of HelloApp
type HelloAppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HelloApp is the Schema for the helloapps API
type HelloApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelloAppSpec   `json:"spec,omitempty"`
	Status HelloAppStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HelloAppList contains a list of HelloApp
type HelloAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelloApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HelloApp{}, &HelloAppList{})
}
