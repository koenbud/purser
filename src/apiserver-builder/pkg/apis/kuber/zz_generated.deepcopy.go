// +build !ignore_autogenerated

/*
 * licensed to vmware.
 */

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package kuber

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MyKind) DeepCopyInto(out *MyKind) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MyKind.
func (in *MyKind) DeepCopy() *MyKind {
	if in == nil {
		return nil
	}
	out := new(MyKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MyKind) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MyKindList) DeepCopyInto(out *MyKindList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MyKind, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MyKindList.
func (in *MyKindList) DeepCopy() *MyKindList {
	if in == nil {
		return nil
	}
	out := new(MyKindList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MyKindList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MyKindSpec) DeepCopyInto(out *MyKindSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MyKindSpec.
func (in *MyKindSpec) DeepCopy() *MyKindSpec {
	if in == nil {
		return nil
	}
	out := new(MyKindSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MyKindStatus) DeepCopyInto(out *MyKindStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MyKindStatus.
func (in *MyKindStatus) DeepCopy() *MyKindStatus {
	if in == nil {
		return nil
	}
	out := new(MyKindStatus)
	in.DeepCopyInto(out)
	return out
}
