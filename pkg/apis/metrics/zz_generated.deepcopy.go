//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package metrics

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsData) DeepCopyInto(out *MetricsData) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsData.
func (in *MetricsData) DeepCopy() *MetricsData {
	if in == nil {
		return nil
	}
	out := new(MetricsData)
	in.DeepCopyInto(out)
	return out
}
