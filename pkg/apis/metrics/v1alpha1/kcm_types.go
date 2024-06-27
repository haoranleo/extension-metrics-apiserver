/*
Copyright 2024.

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
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KCM
// +k8s:openapi-gen=true
// +resource:path=kcms
// +subresource:request=KCMData,path=data,rest=KCMDataREST
type KCM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KCMSpec   `json:"spec,omitempty"`
	Status KCMStatus `json:"status,omitempty"`
}

// KCMList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KCMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KCM `json:"items"`
}

// KCMSpec defines the desired state of KCM
type KCMSpec struct {
}

var _ resource.Object = &KCM{}
var _ resourcestrategy.Validater = &KCM{}

func (in *KCM) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *KCM) NamespaceScoped() bool {
	return false
}

func (in *KCM) New() runtime.Object {
	return &KCM{}
}

func (in *KCM) NewList() runtime.Object {
	return &KCMList{}
}

func (in *KCM) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "metrics.eks.amazonaws.com",
		Version:  "v1alpha1",
		Resource: "kcm",
	}
}

func (in *KCM) IsStorageVersion() bool {
	return true
}

func (in *KCM) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	return nil
}

var _ resource.ObjectList = &KCMList{}

func (in *KCMList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}

// KCMStatus defines the observed state of KCM
type KCMStatus struct {
}

func (in KCMStatus) SubResourceName() string {
	return "status"
}

// KCM implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &KCM{}

func (in *KCM) GetStatus() resource.StatusSubResource {
	return in.Status
}

// KCMStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &KCMStatus{}

func (in KCMStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*KCM).Status = in
}

var _ resource.ObjectWithArbitrarySubResource = &KCM{}

func (in *KCM) GetArbitrarySubResources() []resource.ArbitrarySubResource {
	return []resource.ArbitrarySubResource{
		// +kubebuilder:scaffold:subresource
		&KCMData{},
	}
}
