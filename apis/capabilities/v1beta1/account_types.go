/*
Copyright 2020 Red Hat.

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

package v1beta1

import (
	"reflect"

	"github.com/3scale/3scale-operator/pkg/common"

	"github.com/go-logr/logr"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	AccountKind = "Account"

	// AccountInvalidConditionType represents that the combination of configuration
	// in the AccountSpec is not supported. This is not a transient error, but
	// indicates a state that must be fixed before progress can be made.
	AccountInvalidConditionType common.ConditionType = "Invalid"

	// AccountReadyConditionType indicates the policy has been successfully synchronized.
	// Steady state
	AccountReadyConditionType common.ConditionType = "Ready"

	// AccountFailedConditionType indicates that an error occurred during synchronization.
	// The operator will retry.
	AccountFailedConditionType common.ConditionType = "Failed"
)

// AccountSpec defines the desired state of Account
type AccountSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// OrgName is the organization name
	OrgName string `json:"org_name"`

	// MonthlyBilling
	// +optional
	MonthlyBilling *bool `json:"monthly_billing_enabled,omitempty"`

	// MonthlyCharging
	// +optional
	MonthlyCharging *bool `json:"monthly_charging_enabled,omitempty"`
}

// AccountStatus defines the observed state of Account
type AccountStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	ID *int64 `json:"accountID,omitempty"`

	// ProviderAccountHost contains the 3scale account's provider URL
	// +optional
	ProviderAccountHost string `json:"providerAccountHost,omitempty"`

	// ObservedGeneration reflects the generation of the most recently observed Backend Spec.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Current state of the policy resource.
	// Conditions represent the latest available observations of an object's state
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions common.Conditions `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,2,rep,name=conditions"`
}

func (a *AccountStatus) Equals(other *AccountStatus, logger logr.Logger) bool {
	if !reflect.DeepEqual(a.ID, other.ID) {
		diff := cmp.Diff(a.ID, other.ID)
		logger.V(1).Info("ID not equal", "difference", diff)
		return false
	}

	if a.ProviderAccountHost != other.ProviderAccountHost {
		diff := cmp.Diff(a.ProviderAccountHost, other.ProviderAccountHost)
		logger.V(1).Info("ProviderAccountHost not equal", "difference", diff)
		return false
	}

	if a.ObservedGeneration != other.ObservedGeneration {
		diff := cmp.Diff(a.ObservedGeneration, other.ObservedGeneration)
		logger.V(1).Info("ObservedGeneration not equal", "difference", diff)
		return false
	}

	// Marshalling sorts by condition type
	currentMarshaledJSON, _ := a.Conditions.MarshalJSON()
	otherMarshaledJSON, _ := other.Conditions.MarshalJSON()
	if string(currentMarshaledJSON) != string(otherMarshaledJSON) {
		diff := cmp.Diff(string(currentMarshaledJSON), string(otherMarshaledJSON))
		logger.V(1).Info("Conditions not equal", "difference", diff)
		return false
	}

	return true
}

func (a *AccountStatus) Validate() field.ErrorList {
	errors := field.ErrorList{}
	return errors
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Account is the Schema for the accounts API
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccountSpec   `json:"spec,omitempty"`
	Status AccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountList contains a list of Account
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}
