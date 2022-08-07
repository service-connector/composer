/*
Copyright 2022 The Service Connector Authors.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Ref represents a Kubernetes resource reference
type Ref struct {
	// API version of the Resource
	APIVersion string `json:"apiVersion"`

	// Kind of the Resource
	Kind string `json:"kind"`

	// Name of the Resource
	// Mutually exclusive with a combination of SourceID and TargetPath
	Name string `json:"name,omitempty"`

	// Name of the source resource
	// A combination of SourceID and TargetPath is mutually exclusive with Name
	ResouceID string `json:"resourceID,omitempty"`

	// Path is a JSONPath expression to target resource
	// A combination of SourceID and TargetPath is mutually exclusive with Name
	Path string `json:"path,omitempty"`
}

// Resource represents a Kubernetes resource with an ID
type Resource struct {
	// ID to indentify the resource reference
	ID string `json:"id"`

	// Ref points to a Kubernetes resource reference
	Ref Ref `json:"ref"`
}

// Mapping represents JSONPath expression for a field in the resource
type Mapping struct {
	// target key name
	Name string `json:"name"`

	// JSONPath expression
	Path string `json:"path"`

	// Resource ID
	ResouceID string `json:"resouceID"`

	// Decode Base64
	// +optional
	Base64 bool `json:"base64,omitempty"`

	// Ignore the entry from the target
	// +optional
	Ignore bool `json:"ignore,omitempty"`
}

// Template represents a Go template with access to all the path values
type Template struct {
	Name     string `json:"name"`
	Template string `json:"template"`
}

// ComposerSpec defines the desired state of Composer
type ComposerSpec struct {
	// Type is the type of the service
	// +optional
	Type string `json:"type,omitempty"`
	// Provider is the provider of the service
	// +optional
	Provider string `json:"provider,omitempty"`

	// ServiceAccountName refer the name of the service account used
	// for all reconciliation
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// Resources is a collection of Resouce objects
	// where each resource represents a Kubernetes resource including custom resources
	Resources []Resource `json:"resources"`

	// Mappings is a collection of Path values
	// where each path represents a path represents JSONPath expression for a field in the resource
	Mappings []Mapping `json:"mappings"`

	// Templates is a collection of Template values
	// where each path represents a Go template with access to all the path values
	// +optional
	Templates []Template `json:"templates,omitempty"`
}

// ComposerStatus defines the observed state of Composer
type ComposerStatus struct {
	// Binding exposes the constructed Secret from the backing services resources
	// and comform to Provisioned Service duck-type of Service Binding Specification
	Binding *corev1.LocalObjectReference `json:"binding"`

	// ObservedGeneration is the 'Generation' of the CompositeService that
	// was last processed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Represents the observations of a Composer's current state.
	// Composer.status.conditions.type are: "Available", "Progressing", and "Degraded"
	// Composer.status.conditions.status are one of True, False, Unknown.
	// Composer.status.conditions.reason the value should be a CamelCase string and producers of specific
	// condition types may define expected values and meanings for this field, and whether the values
	// are considered a guaranteed API.
	// Composer.status.conditions.Message is a human readable message indicating details about the transition.
	// For further information see: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Composer is the Schema for the composers API
type Composer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComposerSpec   `json:"spec,omitempty"`
	Status ComposerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ComposerList contains a list of Composer
type ComposerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Composer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Composer{}, &ComposerList{})
}
