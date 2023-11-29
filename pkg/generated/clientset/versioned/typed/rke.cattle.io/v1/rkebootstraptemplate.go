/*
Copyright 2023 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/rancher/rancher/pkg/apis/rke.cattle.io/v1"
	scheme "github.com/rancher/shepherd/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RKEBootstrapTemplatesGetter has a method to return a RKEBootstrapTemplateInterface.
// A group's client should implement this interface.
type RKEBootstrapTemplatesGetter interface {
	RKEBootstrapTemplates(namespace string) RKEBootstrapTemplateInterface
}

// RKEBootstrapTemplateInterface has methods to work with RKEBootstrapTemplate resources.
type RKEBootstrapTemplateInterface interface {
	Create(ctx context.Context, rKEBootstrapTemplate *v1.RKEBootstrapTemplate, opts metav1.CreateOptions) (*v1.RKEBootstrapTemplate, error)
	Update(ctx context.Context, rKEBootstrapTemplate *v1.RKEBootstrapTemplate, opts metav1.UpdateOptions) (*v1.RKEBootstrapTemplate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.RKEBootstrapTemplate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RKEBootstrapTemplateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RKEBootstrapTemplate, err error)
	RKEBootstrapTemplateExpansion
}

// rKEBootstrapTemplates implements RKEBootstrapTemplateInterface
type rKEBootstrapTemplates struct {
	client rest.Interface
	ns     string
}

// newRKEBootstrapTemplates returns a RKEBootstrapTemplates
func newRKEBootstrapTemplates(c *RkeV1Client, namespace string) *rKEBootstrapTemplates {
	return &rKEBootstrapTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the rKEBootstrapTemplate, and returns the corresponding rKEBootstrapTemplate object, and an error if there is any.
func (c *rKEBootstrapTemplates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RKEBootstrapTemplate, err error) {
	result = &v1.RKEBootstrapTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rkebootstraptemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RKEBootstrapTemplates that match those selectors.
func (c *rKEBootstrapTemplates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RKEBootstrapTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RKEBootstrapTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rkebootstraptemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rKEBootstrapTemplates.
func (c *rKEBootstrapTemplates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("rkebootstraptemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a rKEBootstrapTemplate and creates it.  Returns the server's representation of the rKEBootstrapTemplate, and an error, if there is any.
func (c *rKEBootstrapTemplates) Create(ctx context.Context, rKEBootstrapTemplate *v1.RKEBootstrapTemplate, opts metav1.CreateOptions) (result *v1.RKEBootstrapTemplate, err error) {
	result = &v1.RKEBootstrapTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("rkebootstraptemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(rKEBootstrapTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a rKEBootstrapTemplate and updates it. Returns the server's representation of the rKEBootstrapTemplate, and an error, if there is any.
func (c *rKEBootstrapTemplates) Update(ctx context.Context, rKEBootstrapTemplate *v1.RKEBootstrapTemplate, opts metav1.UpdateOptions) (result *v1.RKEBootstrapTemplate, err error) {
	result = &v1.RKEBootstrapTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rkebootstraptemplates").
		Name(rKEBootstrapTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(rKEBootstrapTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the rKEBootstrapTemplate and deletes it. Returns an error if one occurs.
func (c *rKEBootstrapTemplates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rkebootstraptemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rKEBootstrapTemplates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rkebootstraptemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched rKEBootstrapTemplate.
func (c *rKEBootstrapTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RKEBootstrapTemplate, err error) {
	result = &v1.RKEBootstrapTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("rkebootstraptemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
