/*
Copyright The Kube Bind Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/kube-bind/kube-bind/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// APIServiceBindings returns a APIServiceBindingInformer.
	APIServiceBindings() APIServiceBindingInformer
	// APIServiceExports returns a APIServiceExportInformer.
	APIServiceExports() APIServiceExportInformer
	// APIServiceExportRequests returns a APIServiceExportRequestInformer.
	APIServiceExportRequests() APIServiceExportRequestInformer
	// APIServiceExportTemplates returns a APIServiceExportTemplateInformer.
	APIServiceExportTemplates() APIServiceExportTemplateInformer
	// APIServiceNamespaces returns a APIServiceNamespaceInformer.
	APIServiceNamespaces() APIServiceNamespaceInformer
	// ClusterBindings returns a ClusterBindingInformer.
	ClusterBindings() ClusterBindingInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// APIServiceBindings returns a APIServiceBindingInformer.
func (v *version) APIServiceBindings() APIServiceBindingInformer {
	return &aPIServiceBindingInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// APIServiceExports returns a APIServiceExportInformer.
func (v *version) APIServiceExports() APIServiceExportInformer {
	return &aPIServiceExportInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// APIServiceExportRequests returns a APIServiceExportRequestInformer.
func (v *version) APIServiceExportRequests() APIServiceExportRequestInformer {
	return &aPIServiceExportRequestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// APIServiceExportTemplates returns a APIServiceExportTemplateInformer.
func (v *version) APIServiceExportTemplates() APIServiceExportTemplateInformer {
	return &aPIServiceExportTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// APIServiceNamespaces returns a APIServiceNamespaceInformer.
func (v *version) APIServiceNamespaces() APIServiceNamespaceInformer {
	return &aPIServiceNamespaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterBindings returns a ClusterBindingInformer.
func (v *version) ClusterBindings() ClusterBindingInformer {
	return &clusterBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
