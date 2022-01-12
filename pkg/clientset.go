package pkg

import (
	"fmt"
	"net/http"

	"github.com/kiragoo/grpc-asynq-crd-console/pkg/clientset/v1beta1"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	EmqxBrokerV1Beta1Interface() v1beta1.EmqxBrokerInterface
}

type Clientset struct {
	emqxbrokerv1beta1 *v1beta1.EmqxBrokerV1Beta1Client
}

func (c *Clientset) EmqxBrokersV1Beta1() v1beta1.EmqxBrokerV1Beta1Interface {
	return c.emqxbrokerv1beta1
}

func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.emqxbrokerv1beta1, err = v1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}
