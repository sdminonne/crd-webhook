// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/sdminonne/crd-webhook/pkg/apis/mygroup/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMyresources implements MyresourceInterface
type FakeMyresources struct {
	Fake *FakeMygroupV1alpha1
	ns   string
}

var myresourcesResource = schema.GroupVersionResource{Group: "mygroup.amadeus.io", Version: "v1alpha1", Resource: "myresources"}

var myresourcesKind = schema.GroupVersionKind{Group: "mygroup.amadeus.io", Version: "v1alpha1", Kind: "Myresource"}

// Get takes name of the myresource, and returns the corresponding myresource object, and an error if there is any.
func (c *FakeMyresources) Get(name string, options v1.GetOptions) (result *v1alpha1.Myresource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(myresourcesResource, c.ns, name), &v1alpha1.Myresource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Myresource), err
}

// List takes label and field selectors, and returns the list of Myresources that match those selectors.
func (c *FakeMyresources) List(opts v1.ListOptions) (result *v1alpha1.MyresourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(myresourcesResource, myresourcesKind, c.ns, opts), &v1alpha1.MyresourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MyresourceList{}
	for _, item := range obj.(*v1alpha1.MyresourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested myresources.
func (c *FakeMyresources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(myresourcesResource, c.ns, opts))

}

// Create takes the representation of a myresource and creates it.  Returns the server's representation of the myresource, and an error, if there is any.
func (c *FakeMyresources) Create(myresource *v1alpha1.Myresource) (result *v1alpha1.Myresource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(myresourcesResource, c.ns, myresource), &v1alpha1.Myresource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Myresource), err
}

// Update takes the representation of a myresource and updates it. Returns the server's representation of the myresource, and an error, if there is any.
func (c *FakeMyresources) Update(myresource *v1alpha1.Myresource) (result *v1alpha1.Myresource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(myresourcesResource, c.ns, myresource), &v1alpha1.Myresource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Myresource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMyresources) UpdateStatus(myresource *v1alpha1.Myresource) (*v1alpha1.Myresource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(myresourcesResource, "status", c.ns, myresource), &v1alpha1.Myresource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Myresource), err
}

// Delete takes name of the myresource and deletes it. Returns an error if one occurs.
func (c *FakeMyresources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(myresourcesResource, c.ns, name), &v1alpha1.Myresource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMyresources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(myresourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MyresourceList{})
	return err
}

// Patch applies the patch and returns the patched myresource.
func (c *FakeMyresources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Myresource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(myresourcesResource, c.ns, name, data, subresources...), &v1alpha1.Myresource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Myresource), err
}
