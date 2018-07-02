// +build !ignore_autogenerated

/*
 * licensed to vmware.
 */

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	kuber "apiserver-builder/pkg/apis/kuber"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_MyKind_To_kuber_MyKind,
		Convert_kuber_MyKind_To_v1_MyKind,
		Convert_v1_MyKindList_To_kuber_MyKindList,
		Convert_kuber_MyKindList_To_v1_MyKindList,
		Convert_v1_MyKindSpec_To_kuber_MyKindSpec,
		Convert_kuber_MyKindSpec_To_v1_MyKindSpec,
		Convert_v1_MyKindStatus_To_kuber_MyKindStatus,
		Convert_kuber_MyKindStatus_To_v1_MyKindStatus,
		Convert_v1_MyKindStatusStrategy_To_kuber_MyKindStatusStrategy,
		Convert_kuber_MyKindStatusStrategy_To_v1_MyKindStatusStrategy,
		Convert_v1_MyKindStrategy_To_kuber_MyKindStrategy,
		Convert_kuber_MyKindStrategy_To_v1_MyKindStrategy,
	)
}

func autoConvert_v1_MyKind_To_kuber_MyKind(in *MyKind, out *kuber.MyKind, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_MyKindSpec_To_kuber_MyKindSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_MyKindStatus_To_kuber_MyKindStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_MyKind_To_kuber_MyKind is an autogenerated conversion function.
func Convert_v1_MyKind_To_kuber_MyKind(in *MyKind, out *kuber.MyKind, s conversion.Scope) error {
	return autoConvert_v1_MyKind_To_kuber_MyKind(in, out, s)
}

func autoConvert_kuber_MyKind_To_v1_MyKind(in *kuber.MyKind, out *MyKind, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_kuber_MyKindSpec_To_v1_MyKindSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_kuber_MyKindStatus_To_v1_MyKindStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_kuber_MyKind_To_v1_MyKind is an autogenerated conversion function.
func Convert_kuber_MyKind_To_v1_MyKind(in *kuber.MyKind, out *MyKind, s conversion.Scope) error {
	return autoConvert_kuber_MyKind_To_v1_MyKind(in, out, s)
}

func autoConvert_v1_MyKindList_To_kuber_MyKindList(in *MyKindList, out *kuber.MyKindList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]kuber.MyKind)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_MyKindList_To_kuber_MyKindList is an autogenerated conversion function.
func Convert_v1_MyKindList_To_kuber_MyKindList(in *MyKindList, out *kuber.MyKindList, s conversion.Scope) error {
	return autoConvert_v1_MyKindList_To_kuber_MyKindList(in, out, s)
}

func autoConvert_kuber_MyKindList_To_v1_MyKindList(in *kuber.MyKindList, out *MyKindList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]MyKind)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_kuber_MyKindList_To_v1_MyKindList is an autogenerated conversion function.
func Convert_kuber_MyKindList_To_v1_MyKindList(in *kuber.MyKindList, out *MyKindList, s conversion.Scope) error {
	return autoConvert_kuber_MyKindList_To_v1_MyKindList(in, out, s)
}

func autoConvert_v1_MyKindSpec_To_kuber_MyKindSpec(in *MyKindSpec, out *kuber.MyKindSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1_MyKindSpec_To_kuber_MyKindSpec is an autogenerated conversion function.
func Convert_v1_MyKindSpec_To_kuber_MyKindSpec(in *MyKindSpec, out *kuber.MyKindSpec, s conversion.Scope) error {
	return autoConvert_v1_MyKindSpec_To_kuber_MyKindSpec(in, out, s)
}

func autoConvert_kuber_MyKindSpec_To_v1_MyKindSpec(in *kuber.MyKindSpec, out *MyKindSpec, s conversion.Scope) error {
	return nil
}

// Convert_kuber_MyKindSpec_To_v1_MyKindSpec is an autogenerated conversion function.
func Convert_kuber_MyKindSpec_To_v1_MyKindSpec(in *kuber.MyKindSpec, out *MyKindSpec, s conversion.Scope) error {
	return autoConvert_kuber_MyKindSpec_To_v1_MyKindSpec(in, out, s)
}

func autoConvert_v1_MyKindStatus_To_kuber_MyKindStatus(in *MyKindStatus, out *kuber.MyKindStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1_MyKindStatus_To_kuber_MyKindStatus is an autogenerated conversion function.
func Convert_v1_MyKindStatus_To_kuber_MyKindStatus(in *MyKindStatus, out *kuber.MyKindStatus, s conversion.Scope) error {
	return autoConvert_v1_MyKindStatus_To_kuber_MyKindStatus(in, out, s)
}

func autoConvert_kuber_MyKindStatus_To_v1_MyKindStatus(in *kuber.MyKindStatus, out *MyKindStatus, s conversion.Scope) error {
	return nil
}

// Convert_kuber_MyKindStatus_To_v1_MyKindStatus is an autogenerated conversion function.
func Convert_kuber_MyKindStatus_To_v1_MyKindStatus(in *kuber.MyKindStatus, out *MyKindStatus, s conversion.Scope) error {
	return autoConvert_kuber_MyKindStatus_To_v1_MyKindStatus(in, out, s)
}

func autoConvert_v1_MyKindStatusStrategy_To_kuber_MyKindStatusStrategy(in *MyKindStatusStrategy, out *kuber.MyKindStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1_MyKindStatusStrategy_To_kuber_MyKindStatusStrategy is an autogenerated conversion function.
func Convert_v1_MyKindStatusStrategy_To_kuber_MyKindStatusStrategy(in *MyKindStatusStrategy, out *kuber.MyKindStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1_MyKindStatusStrategy_To_kuber_MyKindStatusStrategy(in, out, s)
}

func autoConvert_kuber_MyKindStatusStrategy_To_v1_MyKindStatusStrategy(in *kuber.MyKindStatusStrategy, out *MyKindStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_kuber_MyKindStatusStrategy_To_v1_MyKindStatusStrategy is an autogenerated conversion function.
func Convert_kuber_MyKindStatusStrategy_To_v1_MyKindStatusStrategy(in *kuber.MyKindStatusStrategy, out *MyKindStatusStrategy, s conversion.Scope) error {
	return autoConvert_kuber_MyKindStatusStrategy_To_v1_MyKindStatusStrategy(in, out, s)
}

func autoConvert_v1_MyKindStrategy_To_kuber_MyKindStrategy(in *MyKindStrategy, out *kuber.MyKindStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1_MyKindStrategy_To_kuber_MyKindStrategy is an autogenerated conversion function.
func Convert_v1_MyKindStrategy_To_kuber_MyKindStrategy(in *MyKindStrategy, out *kuber.MyKindStrategy, s conversion.Scope) error {
	return autoConvert_v1_MyKindStrategy_To_kuber_MyKindStrategy(in, out, s)
}

func autoConvert_kuber_MyKindStrategy_To_v1_MyKindStrategy(in *kuber.MyKindStrategy, out *MyKindStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_kuber_MyKindStrategy_To_v1_MyKindStrategy is an autogenerated conversion function.
func Convert_kuber_MyKindStrategy_To_v1_MyKindStrategy(in *kuber.MyKindStrategy, out *MyKindStrategy, s conversion.Scope) error {
	return autoConvert_kuber_MyKindStrategy_To_v1_MyKindStrategy(in, out, s)
}
