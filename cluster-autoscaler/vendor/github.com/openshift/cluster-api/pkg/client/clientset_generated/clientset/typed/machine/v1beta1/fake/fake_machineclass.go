/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachineClasses implements MachineClassInterface
type FakeMachineClasses struct {
	Fake *FakeMachineV1beta1
	ns   string
}

var machineclassesResource = schema.GroupVersionResource{Group: "machine.openshift.io", Version: "v1beta1", Resource: "machineclasses"}

var machineclassesKind = schema.GroupVersionKind{Group: "machine.openshift.io", Version: "v1beta1", Kind: "MachineClass"}

// Get takes name of the machineClass, and returns the corresponding machineClass object, and an error if there is any.
func (c *FakeMachineClasses) Get(name string, options v1.GetOptions) (result *v1beta1.MachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(machineclassesResource, c.ns, name), &v1beta1.MachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineClass), err
}

// List takes label and field selectors, and returns the list of MachineClasses that match those selectors.
func (c *FakeMachineClasses) List(opts v1.ListOptions) (result *v1beta1.MachineClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(machineclassesResource, machineclassesKind, c.ns, opts), &v1beta1.MachineClassList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MachineClassList{ListMeta: obj.(*v1beta1.MachineClassList).ListMeta}
	for _, item := range obj.(*v1beta1.MachineClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineClasses.
func (c *FakeMachineClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(machineclassesResource, c.ns, opts))

}

// Create takes the representation of a machineClass and creates it.  Returns the server's representation of the machineClass, and an error, if there is any.
func (c *FakeMachineClasses) Create(machineClass *v1beta1.MachineClass) (result *v1beta1.MachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(machineclassesResource, c.ns, machineClass), &v1beta1.MachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineClass), err
}

// Update takes the representation of a machineClass and updates it. Returns the server's representation of the machineClass, and an error, if there is any.
func (c *FakeMachineClasses) Update(machineClass *v1beta1.MachineClass) (result *v1beta1.MachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(machineclassesResource, c.ns, machineClass), &v1beta1.MachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineClass), err
}

// Delete takes name of the machineClass and deletes it. Returns an error if one occurs.
func (c *FakeMachineClasses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(machineclassesResource, c.ns, name), &v1beta1.MachineClass{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(machineclassesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.MachineClassList{})
	return err
}

// Patch applies the patch and returns the patched machineClass.
func (c *FakeMachineClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.MachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machineclassesResource, c.ns, name, data, subresources...), &v1beta1.MachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineClass), err
}