/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by angryjet. DO NOT EDIT.

package v1alpha4

import (
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

// GetBindingPhase of this RouteTable.
func (mg *RouteTable) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this RouteTable.
func (mg *RouteTable) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetClassReference of this RouteTable.
func (mg *RouteTable) GetClassReference() *corev1.ObjectReference {
	return mg.Spec.ClassReference
}

// GetCondition of this RouteTable.
func (mg *RouteTable) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetProviderReference of this RouteTable.
func (mg *RouteTable) GetProviderReference() *corev1.ObjectReference {
	return mg.Spec.ProviderReference
}

// GetReclaimPolicy of this RouteTable.
func (mg *RouteTable) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this RouteTable.
func (mg *RouteTable) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this RouteTable.
func (mg *RouteTable) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this RouteTable.
func (mg *RouteTable) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetClassReference of this RouteTable.
func (mg *RouteTable) SetClassReference(r *corev1.ObjectReference) {
	mg.Spec.ClassReference = r
}

// SetConditions of this RouteTable.
func (mg *RouteTable) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetProviderReference of this RouteTable.
func (mg *RouteTable) SetProviderReference(r *corev1.ObjectReference) {
	mg.Spec.ProviderReference = r
}

// SetReclaimPolicy of this RouteTable.
func (mg *RouteTable) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this RouteTable.
func (mg *RouteTable) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
