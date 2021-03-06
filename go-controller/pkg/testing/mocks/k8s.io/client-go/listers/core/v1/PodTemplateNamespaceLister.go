// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	labels "k8s.io/apimachinery/pkg/labels"

	v1 "k8s.io/api/core/v1"
)

// PodTemplateNamespaceLister is an autogenerated mock type for the PodTemplateNamespaceLister type
type PodTemplateNamespaceLister struct {
	mock.Mock
}

// Get provides a mock function with given fields: name
func (_m *PodTemplateNamespaceLister) Get(name string) (*v1.PodTemplate, error) {
	ret := _m.Called(name)

	var r0 *v1.PodTemplate
	if rf, ok := ret.Get(0).(func(string) *v1.PodTemplate); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.PodTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: selector
func (_m *PodTemplateNamespaceLister) List(selector labels.Selector) ([]*v1.PodTemplate, error) {
	ret := _m.Called(selector)

	var r0 []*v1.PodTemplate
	if rf, ok := ret.Get(0).(func(labels.Selector) []*v1.PodTemplate); ok {
		r0 = rf(selector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.PodTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(labels.Selector) error); ok {
		r1 = rf(selector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
