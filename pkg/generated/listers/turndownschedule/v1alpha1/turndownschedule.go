/* Generated Source: Do Not Modify */
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubecost/cluster-turndown/pkg/apis/turndownschedule/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TurndownScheduleLister helps list TurndownSchedules.
type TurndownScheduleLister interface {
	// List lists all TurndownSchedules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TurndownSchedule, err error)
	// Get retrieves the TurndownSchedule from the index for a given name.
	Get(name string) (*v1alpha1.TurndownSchedule, error)
	TurndownScheduleListerExpansion
}

// turndownScheduleLister implements the TurndownScheduleLister interface.
type turndownScheduleLister struct {
	indexer cache.Indexer
}

// NewTurndownScheduleLister returns a new TurndownScheduleLister.
func NewTurndownScheduleLister(indexer cache.Indexer) TurndownScheduleLister {
	return &turndownScheduleLister{indexer: indexer}
}

// List lists all TurndownSchedules in the indexer.
func (s *turndownScheduleLister) List(selector labels.Selector) (ret []*v1alpha1.TurndownSchedule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TurndownSchedule))
	})
	return ret, err
}

// Get retrieves the TurndownSchedule from the index for a given name.
func (s *turndownScheduleLister) Get(name string) (*v1alpha1.TurndownSchedule, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("turndownschedule"), name)
	}
	return obj.(*v1alpha1.TurndownSchedule), nil
}