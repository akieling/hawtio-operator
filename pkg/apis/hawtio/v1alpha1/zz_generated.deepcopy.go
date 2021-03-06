// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hawtconfig) DeepCopyInto(out *Hawtconfig) {
	*out = *in
	out.Branding = in.Branding
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hawtconfig.
func (in *Hawtconfig) DeepCopy() *Hawtconfig {
	if in == nil {
		return nil
	}
	out := new(Hawtconfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hawtio) DeepCopyInto(out *Hawtio) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hawtio.
func (in *Hawtio) DeepCopy() *Hawtio {
	if in == nil {
		return nil
	}
	out := new(Hawtio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Hawtio) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HawtioList) DeepCopyInto(out *HawtioList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Hawtio, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HawtioList.
func (in *HawtioList) DeepCopy() *HawtioList {
	if in == nil {
		return nil
	}
	out := new(HawtioList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HawtioList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HawtioSpec) DeepCopyInto(out *HawtioSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HawtioSpec.
func (in *HawtioSpec) DeepCopy() *HawtioSpec {
	if in == nil {
		return nil
	}
	out := new(HawtioSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HawtioStatus) DeepCopyInto(out *HawtioStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HawtioStatus.
func (in *HawtioStatus) DeepCopy() *HawtioStatus {
	if in == nil {
		return nil
	}
	out := new(HawtioStatus)
	in.DeepCopyInto(out)
	return out
}
