// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	operatorv1 "github.com/openshift/api/operator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOpenShiftAPIServers implements OpenShiftAPIServerInterface
type FakeOpenShiftAPIServers struct {
	Fake *FakeOperatorV1
}

var openshiftapiserversResource = schema.GroupVersionResource{Group: "operator.openshift.io", Version: "v1", Resource: "openshiftapiservers"}

var openshiftapiserversKind = schema.GroupVersionKind{Group: "operator.openshift.io", Version: "v1", Kind: "OpenShiftAPIServer"}

// Get takes name of the openShiftAPIServer, and returns the corresponding openShiftAPIServer object, and an error if there is any.
func (c *FakeOpenShiftAPIServers) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.OpenShiftAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(openshiftapiserversResource, name), &operatorv1.OpenShiftAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.OpenShiftAPIServer), err
}

// List takes label and field selectors, and returns the list of OpenShiftAPIServers that match those selectors.
func (c *FakeOpenShiftAPIServers) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.OpenShiftAPIServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(openshiftapiserversResource, openshiftapiserversKind, opts), &operatorv1.OpenShiftAPIServerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.OpenShiftAPIServerList{ListMeta: obj.(*operatorv1.OpenShiftAPIServerList).ListMeta}
	for _, item := range obj.(*operatorv1.OpenShiftAPIServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested openShiftAPIServers.
func (c *FakeOpenShiftAPIServers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(openshiftapiserversResource, opts))
}

// Create takes the representation of a openShiftAPIServer and creates it.  Returns the server's representation of the openShiftAPIServer, and an error, if there is any.
func (c *FakeOpenShiftAPIServers) Create(ctx context.Context, openShiftAPIServer *operatorv1.OpenShiftAPIServer, opts v1.CreateOptions) (result *operatorv1.OpenShiftAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(openshiftapiserversResource, openShiftAPIServer), &operatorv1.OpenShiftAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.OpenShiftAPIServer), err
}

// Update takes the representation of a openShiftAPIServer and updates it. Returns the server's representation of the openShiftAPIServer, and an error, if there is any.
func (c *FakeOpenShiftAPIServers) Update(ctx context.Context, openShiftAPIServer *operatorv1.OpenShiftAPIServer, opts v1.UpdateOptions) (result *operatorv1.OpenShiftAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(openshiftapiserversResource, openShiftAPIServer), &operatorv1.OpenShiftAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.OpenShiftAPIServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpenShiftAPIServers) UpdateStatus(ctx context.Context, openShiftAPIServer *operatorv1.OpenShiftAPIServer, opts v1.UpdateOptions) (*operatorv1.OpenShiftAPIServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(openshiftapiserversResource, "status", openShiftAPIServer), &operatorv1.OpenShiftAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.OpenShiftAPIServer), err
}

// Delete takes name of the openShiftAPIServer and deletes it. Returns an error if one occurs.
func (c *FakeOpenShiftAPIServers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(openshiftapiserversResource, name, opts), &operatorv1.OpenShiftAPIServer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpenShiftAPIServers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(openshiftapiserversResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.OpenShiftAPIServerList{})
	return err
}

// Patch applies the patch and returns the patched openShiftAPIServer.
func (c *FakeOpenShiftAPIServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.OpenShiftAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(openshiftapiserversResource, name, pt, data, subresources...), &operatorv1.OpenShiftAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.OpenShiftAPIServer), err
}