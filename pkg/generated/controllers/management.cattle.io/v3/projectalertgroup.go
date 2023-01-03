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

package v3

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type ProjectAlertGroupHandler func(string, *v3.ProjectAlertGroup) (*v3.ProjectAlertGroup, error)

type ProjectAlertGroupController interface {
	generic.ControllerMeta
	ProjectAlertGroupClient

	OnChange(ctx context.Context, name string, sync ProjectAlertGroupHandler)
	OnRemove(ctx context.Context, name string, sync ProjectAlertGroupHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() ProjectAlertGroupCache
}

type ProjectAlertGroupClient interface {
	Create(*v3.ProjectAlertGroup) (*v3.ProjectAlertGroup, error)
	Update(*v3.ProjectAlertGroup) (*v3.ProjectAlertGroup, error)
	UpdateStatus(*v3.ProjectAlertGroup) (*v3.ProjectAlertGroup, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v3.ProjectAlertGroup, error)
	List(namespace string, opts metav1.ListOptions) (*v3.ProjectAlertGroupList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.ProjectAlertGroup, err error)
}

type ProjectAlertGroupCache interface {
	Get(namespace, name string) (*v3.ProjectAlertGroup, error)
	List(namespace string, selector labels.Selector) ([]*v3.ProjectAlertGroup, error)

	AddIndexer(indexName string, indexer ProjectAlertGroupIndexer)
	GetByIndex(indexName, key string) ([]*v3.ProjectAlertGroup, error)
}

type ProjectAlertGroupIndexer func(obj *v3.ProjectAlertGroup) ([]string, error)

type projectAlertGroupController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewProjectAlertGroupController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ProjectAlertGroupController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &projectAlertGroupController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromProjectAlertGroupHandlerToHandler(sync ProjectAlertGroupHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.ProjectAlertGroup
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.ProjectAlertGroup))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *projectAlertGroupController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.ProjectAlertGroup))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateProjectAlertGroupDeepCopyOnChange(client ProjectAlertGroupClient, obj *v3.ProjectAlertGroup, handler func(obj *v3.ProjectAlertGroup) (*v3.ProjectAlertGroup, error)) (*v3.ProjectAlertGroup, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *projectAlertGroupController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *projectAlertGroupController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *projectAlertGroupController) OnChange(ctx context.Context, name string, sync ProjectAlertGroupHandler) {
	c.AddGenericHandler(ctx, name, FromProjectAlertGroupHandlerToHandler(sync))
}

func (c *projectAlertGroupController) OnRemove(ctx context.Context, name string, sync ProjectAlertGroupHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromProjectAlertGroupHandlerToHandler(sync)))
}

func (c *projectAlertGroupController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *projectAlertGroupController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *projectAlertGroupController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *projectAlertGroupController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *projectAlertGroupController) Cache() ProjectAlertGroupCache {
	return &projectAlertGroupCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *projectAlertGroupController) Create(obj *v3.ProjectAlertGroup) (*v3.ProjectAlertGroup, error) {
	result := &v3.ProjectAlertGroup{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *projectAlertGroupController) Update(obj *v3.ProjectAlertGroup) (*v3.ProjectAlertGroup, error) {
	result := &v3.ProjectAlertGroup{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *projectAlertGroupController) UpdateStatus(obj *v3.ProjectAlertGroup) (*v3.ProjectAlertGroup, error) {
	result := &v3.ProjectAlertGroup{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *projectAlertGroupController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *projectAlertGroupController) Get(namespace, name string, options metav1.GetOptions) (*v3.ProjectAlertGroup, error) {
	result := &v3.ProjectAlertGroup{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *projectAlertGroupController) List(namespace string, opts metav1.ListOptions) (*v3.ProjectAlertGroupList, error) {
	result := &v3.ProjectAlertGroupList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *projectAlertGroupController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *projectAlertGroupController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v3.ProjectAlertGroup, error) {
	result := &v3.ProjectAlertGroup{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type projectAlertGroupCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *projectAlertGroupCache) Get(namespace, name string) (*v3.ProjectAlertGroup, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.ProjectAlertGroup), nil
}

func (c *projectAlertGroupCache) List(namespace string, selector labels.Selector) (ret []*v3.ProjectAlertGroup, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.ProjectAlertGroup))
	})

	return ret, err
}

func (c *projectAlertGroupCache) AddIndexer(indexName string, indexer ProjectAlertGroupIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.ProjectAlertGroup))
		},
	}))
}

func (c *projectAlertGroupCache) GetByIndex(indexName, key string) (result []*v3.ProjectAlertGroup, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.ProjectAlertGroup, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.ProjectAlertGroup))
	}
	return result, nil
}

type ProjectAlertGroupStatusHandler func(obj *v3.ProjectAlertGroup, status v3.AlertStatus) (v3.AlertStatus, error)

type ProjectAlertGroupGeneratingHandler func(obj *v3.ProjectAlertGroup, status v3.AlertStatus) ([]runtime.Object, v3.AlertStatus, error)

func RegisterProjectAlertGroupStatusHandler(ctx context.Context, controller ProjectAlertGroupController, condition condition.Cond, name string, handler ProjectAlertGroupStatusHandler) {
	statusHandler := &projectAlertGroupStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromProjectAlertGroupHandlerToHandler(statusHandler.sync))
}

func RegisterProjectAlertGroupGeneratingHandler(ctx context.Context, controller ProjectAlertGroupController, apply apply.Apply,
	condition condition.Cond, name string, handler ProjectAlertGroupGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &projectAlertGroupGeneratingHandler{
		ProjectAlertGroupGeneratingHandler: handler,
		apply:                              apply,
		name:                               name,
		gvk:                                controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterProjectAlertGroupStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type projectAlertGroupStatusHandler struct {
	client    ProjectAlertGroupClient
	condition condition.Cond
	handler   ProjectAlertGroupStatusHandler
}

func (a *projectAlertGroupStatusHandler) sync(key string, obj *v3.ProjectAlertGroup) (*v3.ProjectAlertGroup, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type projectAlertGroupGeneratingHandler struct {
	ProjectAlertGroupGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *projectAlertGroupGeneratingHandler) Remove(key string, obj *v3.ProjectAlertGroup) (*v3.ProjectAlertGroup, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v3.ProjectAlertGroup{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *projectAlertGroupGeneratingHandler) Handle(obj *v3.ProjectAlertGroup, status v3.AlertStatus) (v3.AlertStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.ProjectAlertGroupGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
