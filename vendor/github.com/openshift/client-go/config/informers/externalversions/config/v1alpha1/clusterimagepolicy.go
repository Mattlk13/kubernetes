// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	time "time"

	apiconfigv1alpha1 "github.com/openshift/api/config/v1alpha1"
	versioned "github.com/openshift/client-go/config/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/config/informers/externalversions/internalinterfaces"
	configv1alpha1 "github.com/openshift/client-go/config/listers/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterImagePolicyInformer provides access to a shared informer and lister for
// ClusterImagePolicies.
type ClusterImagePolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() configv1alpha1.ClusterImagePolicyLister
}

type clusterImagePolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterImagePolicyInformer constructs a new informer for ClusterImagePolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterImagePolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterImagePolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterImagePolicyInformer constructs a new informer for ClusterImagePolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterImagePolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().ClusterImagePolicies().List(context.Background(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().ClusterImagePolicies().Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().ClusterImagePolicies().List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().ClusterImagePolicies().Watch(ctx, options)
			},
		},
		&apiconfigv1alpha1.ClusterImagePolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterImagePolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterImagePolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterImagePolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiconfigv1alpha1.ClusterImagePolicy{}, f.defaultInformer)
}

func (f *clusterImagePolicyInformer) Lister() configv1alpha1.ClusterImagePolicyLister {
	return configv1alpha1.NewClusterImagePolicyLister(f.Informer().GetIndexer())
}
