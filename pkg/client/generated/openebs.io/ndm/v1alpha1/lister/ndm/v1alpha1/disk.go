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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/ndm/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DiskLister helps list Disks.
type DiskLister interface {
	// List lists all Disks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Disk, err error)
	// Disks returns an object that can list and get Disks.
	Disks(namespace string) DiskNamespaceLister
	DiskListerExpansion
}

// diskLister implements the DiskLister interface.
type diskLister struct {
	indexer cache.Indexer
}

// NewDiskLister returns a new DiskLister.
func NewDiskLister(indexer cache.Indexer) DiskLister {
	return &diskLister{indexer: indexer}
}

// List lists all Disks in the indexer.
func (s *diskLister) List(selector labels.Selector) (ret []*v1alpha1.Disk, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Disk))
	})
	return ret, err
}

// Disks returns an object that can list and get Disks.
func (s *diskLister) Disks(namespace string) DiskNamespaceLister {
	return diskNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DiskNamespaceLister helps list and get Disks.
type DiskNamespaceLister interface {
	// List lists all Disks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Disk, err error)
	// Get retrieves the Disk from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Disk, error)
	DiskNamespaceListerExpansion
}

// diskNamespaceLister implements the DiskNamespaceLister
// interface.
type diskNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Disks in the indexer for a given namespace.
func (s diskNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Disk, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Disk))
	})
	return ret, err
}

// Get retrieves the Disk from the indexer for a given namespace and name.
func (s diskNamespaceLister) Get(name string) (*v1alpha1.Disk, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("disk"), name)
	}
	return obj.(*v1alpha1.Disk), nil
}
