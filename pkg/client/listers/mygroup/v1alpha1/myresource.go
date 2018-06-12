// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/sdminonne/crd-webhook/pkg/apis/mygroup/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MyresourceLister helps list Myresources.
type MyresourceLister interface {
	// List lists all Myresources in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Myresource, err error)
	// Myresources returns an object that can list and get Myresources.
	Myresources(namespace string) MyresourceNamespaceLister
	MyresourceListerExpansion
}

// myresourceLister implements the MyresourceLister interface.
type myresourceLister struct {
	indexer cache.Indexer
}

// NewMyresourceLister returns a new MyresourceLister.
func NewMyresourceLister(indexer cache.Indexer) MyresourceLister {
	return &myresourceLister{indexer: indexer}
}

// List lists all Myresources in the indexer.
func (s *myresourceLister) List(selector labels.Selector) (ret []*v1alpha1.Myresource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Myresource))
	})
	return ret, err
}

// Myresources returns an object that can list and get Myresources.
func (s *myresourceLister) Myresources(namespace string) MyresourceNamespaceLister {
	return myresourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MyresourceNamespaceLister helps list and get Myresources.
type MyresourceNamespaceLister interface {
	// List lists all Myresources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Myresource, err error)
	// Get retrieves the Myresource from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Myresource, error)
	MyresourceNamespaceListerExpansion
}

// myresourceNamespaceLister implements the MyresourceNamespaceLister
// interface.
type myresourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Myresources in the indexer for a given namespace.
func (s myresourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Myresource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Myresource))
	})
	return ret, err
}

// Get retrieves the Myresource from the indexer for a given namespace and name.
func (s myresourceNamespaceLister) Get(name string) (*v1alpha1.Myresource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("myresource"), name)
	}
	return obj.(*v1alpha1.Myresource), nil
}
