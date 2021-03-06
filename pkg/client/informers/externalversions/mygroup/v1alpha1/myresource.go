// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	mygroup_v1alpha1 "github.com/sdminonne/crd-webhook/pkg/apis/mygroup/v1alpha1"
	versioned "github.com/sdminonne/crd-webhook/pkg/client/clientset/versioned"
	internalinterfaces "github.com/sdminonne/crd-webhook/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/sdminonne/crd-webhook/pkg/client/listers/mygroup/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MyresourceInformer provides access to a shared informer and lister for
// Myresources.
type MyresourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MyresourceLister
}

type myresourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMyresourceInformer constructs a new informer for Myresource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMyresourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMyresourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMyresourceInformer constructs a new informer for Myresource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMyresourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MygroupV1alpha1().Myresources(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MygroupV1alpha1().Myresources(namespace).Watch(options)
			},
		},
		&mygroup_v1alpha1.Myresource{},
		resyncPeriod,
		indexers,
	)
}

func (f *myresourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMyresourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *myresourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mygroup_v1alpha1.Myresource{}, f.defaultInformer)
}

func (f *myresourceInformer) Lister() v1alpha1.MyresourceLister {
	return v1alpha1.NewMyresourceLister(f.Informer().GetIndexer())
}
