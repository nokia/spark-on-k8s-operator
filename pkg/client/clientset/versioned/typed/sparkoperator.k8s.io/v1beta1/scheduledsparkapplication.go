// Code generated by k8s code-generator DO NOT EDIT.

/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1beta1 "github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta1"
	scheme "github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ScheduledSparkApplicationsGetter has a method to return a ScheduledSparkApplicationInterface.
// A group's client should implement this interface.
type ScheduledSparkApplicationsGetter interface {
	ScheduledSparkApplications(namespace string) ScheduledSparkApplicationInterface
}

// ScheduledSparkApplicationInterface has methods to work with ScheduledSparkApplication resources.
type ScheduledSparkApplicationInterface interface {
	Create(*v1beta1.ScheduledSparkApplication) (*v1beta1.ScheduledSparkApplication, error)
	Update(*v1beta1.ScheduledSparkApplication) (*v1beta1.ScheduledSparkApplication, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ScheduledSparkApplication, error)
	List(opts v1.ListOptions) (*v1beta1.ScheduledSparkApplicationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ScheduledSparkApplication, err error)
	ScheduledSparkApplicationExpansion
}

// scheduledSparkApplications implements ScheduledSparkApplicationInterface
type scheduledSparkApplications struct {
	client rest.Interface
	ns     string
}

// newScheduledSparkApplications returns a ScheduledSparkApplications
func newScheduledSparkApplications(c *SparkoperatorV1beta1Client, namespace string) *scheduledSparkApplications {
	return &scheduledSparkApplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the scheduledSparkApplication, and returns the corresponding scheduledSparkApplication object, and an error if there is any.
func (c *scheduledSparkApplications) Get(name string, options v1.GetOptions) (result *v1beta1.ScheduledSparkApplication, err error) {
	result = &v1beta1.ScheduledSparkApplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ScheduledSparkApplications that match those selectors.
func (c *scheduledSparkApplications) List(opts v1.ListOptions) (result *v1beta1.ScheduledSparkApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ScheduledSparkApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested scheduledSparkApplications.
func (c *scheduledSparkApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a scheduledSparkApplication and creates it.  Returns the server's representation of the scheduledSparkApplication, and an error, if there is any.
func (c *scheduledSparkApplications) Create(scheduledSparkApplication *v1beta1.ScheduledSparkApplication) (result *v1beta1.ScheduledSparkApplication, err error) {
	result = &v1beta1.ScheduledSparkApplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		Body(scheduledSparkApplication).
		Do().
		Into(result)
	return
}

// Update takes the representation of a scheduledSparkApplication and updates it. Returns the server's representation of the scheduledSparkApplication, and an error, if there is any.
func (c *scheduledSparkApplications) Update(scheduledSparkApplication *v1beta1.ScheduledSparkApplication) (result *v1beta1.ScheduledSparkApplication, err error) {
	result = &v1beta1.ScheduledSparkApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		Name(scheduledSparkApplication.Name).
		Body(scheduledSparkApplication).
		Do().
		Into(result)
	return
}

// Delete takes name of the scheduledSparkApplication and deletes it. Returns an error if one occurs.
func (c *scheduledSparkApplications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *scheduledSparkApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched scheduledSparkApplication.
func (c *scheduledSparkApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ScheduledSparkApplication, err error) {
	result = &v1beta1.ScheduledSparkApplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
