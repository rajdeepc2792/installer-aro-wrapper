// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/openshift/client-go/machine/clientset/versioned/typed/machine/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMachineV1beta1 struct {
	*testing.Fake
}

func (c *FakeMachineV1beta1) Machines(namespace string) v1beta1.MachineInterface {
	return &FakeMachines{c, namespace}
}

func (c *FakeMachineV1beta1) MachineHealthChecks(namespace string) v1beta1.MachineHealthCheckInterface {
	return &FakeMachineHealthChecks{c, namespace}
}

func (c *FakeMachineV1beta1) MachineSets(namespace string) v1beta1.MachineSetInterface {
	return &FakeMachineSets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMachineV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}