// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/sdminonne/crd-webhook/pkg/apis/mygroup/v1alpha1"
	scheme "github.com/sdminonne/crd-webhook/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MyresourcesGetter has a method to return a MyresourceInterface.
// A group's client should implement this interface.
type MyresourcesGetter interface {
	Myresources(namespace string) MyresourceInterface
}

// MyresourceInterface has methods to work with Myresource resources.
type MyresourceInterface interface {
	Create(*v1alpha1.Myresource) (*v1alpha1.Myresource, error)
	Update(*v1alpha1.Myresource) (*v1alpha1.Myresource, error)
	UpdateStatus(*v1alpha1.Myresource) (*v1alpha1.Myresource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Myresource, error)
	List(opts v1.ListOptions) (*v1alpha1.MyresourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Myresource, err error)
	MyresourceExpansion
}

// myresources implements MyresourceInterface
type myresources struct {
	client rest.Interface
	ns     string
}

// newMyresources returns a Myresources
func newMyresources(c *MygroupV1alpha1Client, namespace string) *myresources {
	return &myresources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the myresource, and returns the corresponding myresource object, and an error if there is any.
func (c *myresources) Get(name string, options v1.GetOptions) (result *v1alpha1.Myresource, err error) {
	result = &v1alpha1.Myresource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("myresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Myresources that match those selectors.
func (c *myresources) List(opts v1.ListOptions) (result *v1alpha1.MyresourceList, err error) {
	result = &v1alpha1.MyresourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("myresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested myresources.
func (c *myresources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("myresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a myresource and creates it.  Returns the server's representation of the myresource, and an error, if there is any.
func (c *myresources) Create(myresource *v1alpha1.Myresource) (result *v1alpha1.Myresource, err error) {
	result = &v1alpha1.Myresource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("myresources").
		Body(myresource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a myresource and updates it. Returns the server's representation of the myresource, and an error, if there is any.
func (c *myresources) Update(myresource *v1alpha1.Myresource) (result *v1alpha1.Myresource, err error) {
	result = &v1alpha1.Myresource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("myresources").
		Name(myresource.Name).
		Body(myresource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *myresources) UpdateStatus(myresource *v1alpha1.Myresource) (result *v1alpha1.Myresource, err error) {
	result = &v1alpha1.Myresource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("myresources").
		Name(myresource.Name).
		SubResource("status").
		Body(myresource).
		Do().
		Into(result)
	return
}

// Delete takes name of the myresource and deletes it. Returns an error if one occurs.
func (c *myresources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("myresources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *myresources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("myresources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched myresource.
func (c *myresources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Myresource, err error) {
	result = &v1alpha1.Myresource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("myresources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
