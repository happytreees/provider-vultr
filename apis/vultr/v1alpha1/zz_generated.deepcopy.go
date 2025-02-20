//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kubernetes) DeepCopyInto(out *Kubernetes) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kubernetes.
func (in *Kubernetes) DeepCopy() *Kubernetes {
	if in == nil {
		return nil
	}
	out := new(Kubernetes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kubernetes) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesList) DeepCopyInto(out *KubernetesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kubernetes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesList.
func (in *KubernetesList) DeepCopy() *KubernetesList {
	if in == nil {
		return nil
	}
	out := new(KubernetesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubernetesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesObservation) DeepCopyInto(out *KubernetesObservation) {
	*out = *in
	if in.ClusterSubnet != nil {
		in, out := &in.ClusterSubnet, &out.ClusterSubnet
		*out = new(string)
		**out = **in
	}
	if in.DateCreated != nil {
		in, out := &in.DateCreated, &out.DateCreated
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make([]NodePoolsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ServiceSubnet != nil {
		in, out := &in.ServiceSubnet, &out.ServiceSubnet
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesObservation.
func (in *KubernetesObservation) DeepCopy() *KubernetesObservation {
	if in == nil {
		return nil
	}
	out := new(KubernetesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesParameters) DeepCopyInto(out *KubernetesParameters) {
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make([]NodePoolsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesParameters.
func (in *KubernetesParameters) DeepCopy() *KubernetesParameters {
	if in == nil {
		return nil
	}
	out := new(KubernetesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesSpec) DeepCopyInto(out *KubernetesSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesSpec.
func (in *KubernetesSpec) DeepCopy() *KubernetesSpec {
	if in == nil {
		return nil
	}
	out := new(KubernetesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesStatus) DeepCopyInto(out *KubernetesStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesStatus.
func (in *KubernetesStatus) DeepCopy() *KubernetesStatus {
	if in == nil {
		return nil
	}
	out := new(KubernetesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolsObservation) DeepCopyInto(out *NodePoolsObservation) {
	*out = *in
	if in.AutoScaler != nil {
		in, out := &in.AutoScaler, &out.AutoScaler
		*out = new(bool)
		**out = **in
	}
	if in.DateCreated != nil {
		in, out := &in.DateCreated, &out.DateCreated
		*out = new(string)
		**out = **in
	}
	if in.DateUpdated != nil {
		in, out := &in.DateUpdated, &out.DateUpdated
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.MaxNodes != nil {
		in, out := &in.MaxNodes, &out.MaxNodes
		*out = new(float64)
		**out = **in
	}
	if in.MinNodes != nil {
		in, out := &in.MinNodes, &out.MinNodes
		*out = new(float64)
		**out = **in
	}
	if in.NodeQuantity != nil {
		in, out := &in.NodeQuantity, &out.NodeQuantity
		*out = new(float64)
		**out = **in
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]NodesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolsObservation.
func (in *NodePoolsObservation) DeepCopy() *NodePoolsObservation {
	if in == nil {
		return nil
	}
	out := new(NodePoolsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolsParameters) DeepCopyInto(out *NodePoolsParameters) {
	*out = *in
	if in.AutoScaler != nil {
		in, out := &in.AutoScaler, &out.AutoScaler
		*out = new(bool)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.MaxNodes != nil {
		in, out := &in.MaxNodes, &out.MaxNodes
		*out = new(float64)
		**out = **in
	}
	if in.MinNodes != nil {
		in, out := &in.MinNodes, &out.MinNodes
		*out = new(float64)
		**out = **in
	}
	if in.NodeQuantity != nil {
		in, out := &in.NodeQuantity, &out.NodeQuantity
		*out = new(float64)
		**out = **in
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolsParameters.
func (in *NodePoolsParameters) DeepCopy() *NodePoolsParameters {
	if in == nil {
		return nil
	}
	out := new(NodePoolsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodesObservation) DeepCopyInto(out *NodesObservation) {
	*out = *in
	if in.DateCreated != nil {
		in, out := &in.DateCreated, &out.DateCreated
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodesObservation.
func (in *NodesObservation) DeepCopy() *NodesObservation {
	if in == nil {
		return nil
	}
	out := new(NodesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodesParameters) DeepCopyInto(out *NodesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodesParameters.
func (in *NodesParameters) DeepCopy() *NodesParameters {
	if in == nil {
		return nil
	}
	out := new(NodesParameters)
	in.DeepCopyInto(out)
	return out
}
