package v1beta1

import (
	"context"
	"time"

	"github.com/emqx/emqx-operator/api/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/kubectl/pkg/scheme"
)

type EmqxBrokersGetter interface {
	EmqxBrokers(namespace string) EmqxBrokerInterface
}
type EmqxBrokerInterface interface {
	List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.EmqxBrokerList, error)
	Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta1.EmqxBroker, error)
	Create(ctx context.Context, emqxbroker *v1beta1.EmqxBroker, opts metav1.CreateOptions) (*v1beta1.EmqxBroker, error)
	Delete(ctx context.Context, name string, options metav1.DeleteOptions) error
	// Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type emqxBrokers struct {
	client rest.Interface
	ns     string
}

func NewEmqxBrokers(c *EmqxBrokerV1Beta1Client, ns string) *emqxBrokers {
	return &emqxBrokers{
		client: c.RESTClient(),
		ns:     ns,
	}
}

func (c *emqxBrokers) List(ctx context.Context, opts metav1.ListOptions) (result *v1beta1.EmqxBrokerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.EmqxBrokerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("emqxbrokers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *emqxBrokers) Get(ctx context.Context, name string, opts metav1.GetOptions) (result *v1beta1.EmqxBroker, err error) {
	result = &v1beta1.EmqxBroker{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("emqxbrokers").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *emqxBrokers) Create(ctx context.Context, emqxbroker *v1beta1.EmqxBroker, opts metav1.CreateOptions) (result *v1beta1.EmqxBroker, err error) {
	result = &v1beta1.EmqxBroker{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("emqxbrokers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(emqxbroker).
		Do(ctx).
		Into(result)
	return
}

func (c *emqxBrokers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("emqxbrokers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}
