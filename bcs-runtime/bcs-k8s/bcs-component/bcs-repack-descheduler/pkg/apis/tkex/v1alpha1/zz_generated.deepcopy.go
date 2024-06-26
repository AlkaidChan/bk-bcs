//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DescheduleBalanceStrategy) DeepCopyInto(out *DescheduleBalanceStrategy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DescheduleBalanceStrategy.
func (in *DescheduleBalanceStrategy) DeepCopy() *DescheduleBalanceStrategy {
	if in == nil {
		return nil
	}
	out := new(DescheduleBalanceStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DescheduleConvergeStrategy) DeepCopyInto(out *DescheduleConvergeStrategy) {
	*out = *in
	out.ProfitTarget = in.ProfitTarget
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DescheduleConvergeStrategy.
func (in *DescheduleConvergeStrategy) DeepCopy() *DescheduleConvergeStrategy {
	if in == nil {
		return nil
	}
	out := new(DescheduleConvergeStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeschedulePolicy) DeepCopyInto(out *DeschedulePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeschedulePolicy.
func (in *DeschedulePolicy) DeepCopy() *DeschedulePolicy {
	if in == nil {
		return nil
	}
	out := new(DeschedulePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeschedulePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeschedulePolicyAnnotator) DeepCopyInto(out *DeschedulePolicyAnnotator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeschedulePolicyAnnotator.
func (in *DeschedulePolicyAnnotator) DeepCopy() *DeschedulePolicyAnnotator {
	if in == nil {
		return nil
	}
	out := new(DeschedulePolicyAnnotator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeschedulePolicyList) DeepCopyInto(out *DeschedulePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeschedulePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeschedulePolicyList.
func (in *DeschedulePolicyList) DeepCopy() *DeschedulePolicyList {
	if in == nil {
		return nil
	}
	out := new(DeschedulePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeschedulePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DescheduleSpec) DeepCopyInto(out *DescheduleSpec) {
	*out = *in
	out.Converge = in.Converge
	out.Balance = in.Balance
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DescheduleSpec.
func (in *DescheduleSpec) DeepCopy() *DescheduleSpec {
	if in == nil {
		return nil
	}
	out := new(DescheduleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfitTarget) DeepCopyInto(out *ProfitTarget) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfitTarget.
func (in *ProfitTarget) DeepCopy() *ProfitTarget {
	if in == nil {
		return nil
	}
	out := new(ProfitTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Sequence) DeepCopyInto(out *Sequence) {
	{
		in := &in
		*out = make(Sequence, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sequence.
func (in Sequence) DeepCopy() Sequence {
	if in == nil {
		return nil
	}
	out := new(Sequence)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceItem) DeepCopyInto(out *SequenceItem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceItem.
func (in *SequenceItem) DeepCopy() *SequenceItem {
	if in == nil {
		return nil
	}
	out := new(SequenceItem)
	in.DeepCopyInto(out)
	return out
}
