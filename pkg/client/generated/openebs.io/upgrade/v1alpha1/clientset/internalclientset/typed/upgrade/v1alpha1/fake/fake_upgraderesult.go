/*
Copyright 2019 The OpenEBS Authors

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
	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/upgrade/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUpgradeResults implements UpgradeResultInterface
type FakeUpgradeResults struct {
	Fake *FakeOpenebsV1alpha1
	ns   string
}

var upgraderesultsResource = schema.GroupVersionResource{Group: "openebs.io", Version: "v1alpha1", Resource: "upgraderesults"}

var upgraderesultsKind = schema.GroupVersionKind{Group: "openebs.io", Version: "v1alpha1", Kind: "UpgradeResult"}

// Get takes name of the upgradeResult, and returns the corresponding upgradeResult object, and an error if there is any.
func (c *FakeUpgradeResults) Get(name string, options v1.GetOptions) (result *v1alpha1.UpgradeResult, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(upgraderesultsResource, c.ns, name), &v1alpha1.UpgradeResult{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UpgradeResult), err
}

// List takes label and field selectors, and returns the list of UpgradeResults that match those selectors.
func (c *FakeUpgradeResults) List(opts v1.ListOptions) (result *v1alpha1.UpgradeResultList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(upgraderesultsResource, upgraderesultsKind, c.ns, opts), &v1alpha1.UpgradeResultList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UpgradeResultList{ListMeta: obj.(*v1alpha1.UpgradeResultList).ListMeta}
	for _, item := range obj.(*v1alpha1.UpgradeResultList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested upgradeResults.
func (c *FakeUpgradeResults) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(upgraderesultsResource, c.ns, opts))

}

// Create takes the representation of a upgradeResult and creates it.  Returns the server's representation of the upgradeResult, and an error, if there is any.
func (c *FakeUpgradeResults) Create(upgradeResult *v1alpha1.UpgradeResult) (result *v1alpha1.UpgradeResult, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(upgraderesultsResource, c.ns, upgradeResult), &v1alpha1.UpgradeResult{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UpgradeResult), err
}

// Update takes the representation of a upgradeResult and updates it. Returns the server's representation of the upgradeResult, and an error, if there is any.
func (c *FakeUpgradeResults) Update(upgradeResult *v1alpha1.UpgradeResult) (result *v1alpha1.UpgradeResult, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(upgraderesultsResource, c.ns, upgradeResult), &v1alpha1.UpgradeResult{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UpgradeResult), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUpgradeResults) UpdateStatus(upgradeResult *v1alpha1.UpgradeResult) (*v1alpha1.UpgradeResult, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(upgraderesultsResource, "status", c.ns, upgradeResult), &v1alpha1.UpgradeResult{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UpgradeResult), err
}

// Delete takes name of the upgradeResult and deletes it. Returns an error if one occurs.
func (c *FakeUpgradeResults) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(upgraderesultsResource, c.ns, name), &v1alpha1.UpgradeResult{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUpgradeResults) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(upgraderesultsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.UpgradeResultList{})
	return err
}

// Patch applies the patch and returns the patched upgradeResult.
func (c *FakeUpgradeResults) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UpgradeResult, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(upgraderesultsResource, c.ns, name, pt, data, subresources...), &v1alpha1.UpgradeResult{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UpgradeResult), err
}
