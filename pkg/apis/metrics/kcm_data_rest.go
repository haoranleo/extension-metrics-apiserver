package metrics

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/client-go/kubernetes"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	restclient "k8s.io/client-go/rest"
)

// +k8s:deepcopy-gen=false
type KCMDataREST struct {
	podClient corev1client.CoreV1Interface
}

// Status Subresource
var _ rest.Getter = &KCMDataREST{}

func NewKCMDataREST(getter generic.RESTOptionsGetter) rest.Storage {
	inClusterClientConfig, err := restclient.InClusterConfig()
	if err != nil {
		panic(err)
	}
	client := kubernetes.NewForConfigOrDie(inClusterClientConfig)

	return &KCMDataREST{
		podClient: client.CoreV1(),
	}
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *KCMDataREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return &MetricsData{
		data: "Hello, World!",
	}, nil
}

func (r *KCMDataREST) New() runtime.Object {
	return &MetricsData{}
}

type MetricsData struct {
	data string
}

func (obj *MetricsData) GetObjectKind() schema.ObjectKind {
	return schema.EmptyObjectKind
}
func (obj *MetricsData) DeepCopyObject() runtime.Object {
	m := *obj
	return &m
}
