/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	autoscalingxk8siov1 "k8s.io/autoscaler/cluster-autoscaler/apis/provisioningrequest/autoscaling.x-k8s.io/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ProvisioningRequestStatusApplyConfiguration represents a declarative configuration of the ProvisioningRequestStatus type for use
// with apply.
type ProvisioningRequestStatusApplyConfiguration struct {
	Conditions               []v1.ConditionApplyConfiguration      `json:"conditions,omitempty"`
	ProvisioningClassDetails map[string]autoscalingxk8siov1.Detail `json:"provisioningClassDetails,omitempty"`
}

// ProvisioningRequestStatusApplyConfiguration constructs a declarative configuration of the ProvisioningRequestStatus type for use with
// apply.
func ProvisioningRequestStatus() *ProvisioningRequestStatusApplyConfiguration {
	return &ProvisioningRequestStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ProvisioningRequestStatusApplyConfiguration) WithConditions(values ...*v1.ConditionApplyConfiguration) *ProvisioningRequestStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithProvisioningClassDetails puts the entries into the ProvisioningClassDetails field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ProvisioningClassDetails field,
// overwriting an existing map entries in ProvisioningClassDetails field with the same key.
func (b *ProvisioningRequestStatusApplyConfiguration) WithProvisioningClassDetails(entries map[string]autoscalingxk8siov1.Detail) *ProvisioningRequestStatusApplyConfiguration {
	if b.ProvisioningClassDetails == nil && len(entries) > 0 {
		b.ProvisioningClassDetails = make(map[string]autoscalingxk8siov1.Detail, len(entries))
	}
	for k, v := range entries {
		b.ProvisioningClassDetails[k] = v
	}
	return b
}