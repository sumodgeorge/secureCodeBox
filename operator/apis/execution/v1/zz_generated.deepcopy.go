// +build !ignore_autogenerated

/*
Copyright 2020 iteratec GmbH.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtractResults) DeepCopyInto(out *ExtractResults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtractResults.
func (in *ExtractResults) DeepCopy() *ExtractResults {
	if in == nil {
		return nil
	}
	out := new(ExtractResults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FindingSeverities) DeepCopyInto(out *FindingSeverities) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FindingSeverities.
func (in *FindingSeverities) DeepCopy() *FindingSeverities {
	if in == nil {
		return nil
	}
	out := new(FindingSeverities)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FindingStats) DeepCopyInto(out *FindingStats) {
	*out = *in
	out.FindingSeverities = in.FindingSeverities
	if in.FindingCategories != nil {
		in, out := &in.FindingCategories, &out.FindingCategories
		*out = make(map[string]uint64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FindingStats.
func (in *FindingStats) DeepCopy() *FindingStats {
	if in == nil {
		return nil
	}
	out := new(FindingStats)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HookStatus) DeepCopyInto(out *HookStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HookStatus.
func (in *HookStatus) DeepCopy() *HookStatus {
	if in == nil {
		return nil
	}
	out := new(HookStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParseDefinition) DeepCopyInto(out *ParseDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParseDefinition.
func (in *ParseDefinition) DeepCopy() *ParseDefinition {
	if in == nil {
		return nil
	}
	out := new(ParseDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParseDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParseDefinitionList) DeepCopyInto(out *ParseDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ParseDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParseDefinitionList.
func (in *ParseDefinitionList) DeepCopy() *ParseDefinitionList {
	if in == nil {
		return nil
	}
	out := new(ParseDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParseDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParseDefinitionSpec) DeepCopyInto(out *ParseDefinitionSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParseDefinitionSpec.
func (in *ParseDefinitionSpec) DeepCopy() *ParseDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(ParseDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParseDefinitionStatus) DeepCopyInto(out *ParseDefinitionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParseDefinitionStatus.
func (in *ParseDefinitionStatus) DeepCopy() *ParseDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(ParseDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scan) DeepCopyInto(out *Scan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scan.
func (in *Scan) DeepCopy() *Scan {
	if in == nil {
		return nil
	}
	out := new(Scan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Scan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanCompletionHook) DeepCopyInto(out *ScanCompletionHook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanCompletionHook.
func (in *ScanCompletionHook) DeepCopy() *ScanCompletionHook {
	if in == nil {
		return nil
	}
	out := new(ScanCompletionHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScanCompletionHook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanCompletionHookList) DeepCopyInto(out *ScanCompletionHookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScanCompletionHook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanCompletionHookList.
func (in *ScanCompletionHookList) DeepCopy() *ScanCompletionHookList {
	if in == nil {
		return nil
	}
	out := new(ScanCompletionHookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScanCompletionHookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanCompletionHookSpec) DeepCopyInto(out *ScanCompletionHookSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanCompletionHookSpec.
func (in *ScanCompletionHookSpec) DeepCopy() *ScanCompletionHookSpec {
	if in == nil {
		return nil
	}
	out := new(ScanCompletionHookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanCompletionHookStatus) DeepCopyInto(out *ScanCompletionHookStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanCompletionHookStatus.
func (in *ScanCompletionHookStatus) DeepCopy() *ScanCompletionHookStatus {
	if in == nil {
		return nil
	}
	out := new(ScanCompletionHookStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanList) DeepCopyInto(out *ScanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Scan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanList.
func (in *ScanList) DeepCopy() *ScanList {
	if in == nil {
		return nil
	}
	out := new(ScanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanSpec) DeepCopyInto(out *ScanSpec) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Cascades != nil {
		in, out := &in.Cascades, &out.Cascades
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanSpec.
func (in *ScanSpec) DeepCopy() *ScanSpec {
	if in == nil {
		return nil
	}
	out := new(ScanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanStatus) DeepCopyInto(out *ScanStatus) {
	*out = *in
	in.Findings.DeepCopyInto(&out.Findings)
	if in.ReadAndWriteHookStatus != nil {
		in, out := &in.ReadAndWriteHookStatus, &out.ReadAndWriteHookStatus
		*out = make([]HookStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanStatus.
func (in *ScanStatus) DeepCopy() *ScanStatus {
	if in == nil {
		return nil
	}
	out := new(ScanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanType) DeepCopyInto(out *ScanType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanType.
func (in *ScanType) DeepCopy() *ScanType {
	if in == nil {
		return nil
	}
	out := new(ScanType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScanType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanTypeList) DeepCopyInto(out *ScanTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScanType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanTypeList.
func (in *ScanTypeList) DeepCopy() *ScanTypeList {
	if in == nil {
		return nil
	}
	out := new(ScanTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScanTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanTypeSpec) DeepCopyInto(out *ScanTypeSpec) {
	*out = *in
	out.ExtractResults = in.ExtractResults
	in.JobTemplate.DeepCopyInto(&out.JobTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanTypeSpec.
func (in *ScanTypeSpec) DeepCopy() *ScanTypeSpec {
	if in == nil {
		return nil
	}
	out := new(ScanTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanTypeStatus) DeepCopyInto(out *ScanTypeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanTypeStatus.
func (in *ScanTypeStatus) DeepCopy() *ScanTypeStatus {
	if in == nil {
		return nil
	}
	out := new(ScanTypeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledScan) DeepCopyInto(out *ScheduledScan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledScan.
func (in *ScheduledScan) DeepCopy() *ScheduledScan {
	if in == nil {
		return nil
	}
	out := new(ScheduledScan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledScan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledScanList) DeepCopyInto(out *ScheduledScanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScheduledScan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledScanList.
func (in *ScheduledScanList) DeepCopy() *ScheduledScanList {
	if in == nil {
		return nil
	}
	out := new(ScheduledScanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledScanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledScanSpec) DeepCopyInto(out *ScheduledScanSpec) {
	*out = *in
	out.Interval = in.Interval
	if in.ScanSpec != nil {
		in, out := &in.ScanSpec, &out.ScanSpec
		*out = new(ScanSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledScanSpec.
func (in *ScheduledScanSpec) DeepCopy() *ScheduledScanSpec {
	if in == nil {
		return nil
	}
	out := new(ScheduledScanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledScanStatus) DeepCopyInto(out *ScheduledScanStatus) {
	*out = *in
	if in.LastScheduleTime != nil {
		in, out := &in.LastScheduleTime, &out.LastScheduleTime
		*out = (*in).DeepCopy()
	}
	in.Findings.DeepCopyInto(&out.Findings)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledScanStatus.
func (in *ScheduledScanStatus) DeepCopy() *ScheduledScanStatus {
	if in == nil {
		return nil
	}
	out := new(ScheduledScanStatus)
	in.DeepCopyInto(out)
	return out
}
