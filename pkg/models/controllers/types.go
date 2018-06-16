/*
Copyright 2018 The KubeSphere Authors.

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

package controllers

import (
	"time"

	"github.com/jinzhu/gorm"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	stopping                   = "stopped"
	running                    = "running"
	updating                   = "updating"
	tablePods                  = "pods"
	tableDeployments           = "deployments"
	tableDaemonsets            = "daemonsets"
	tableStatefulsets          = "statefulsets"
	tableNamespaces            = "namespaces"
	tableIngresses             = "ingresses"
	tablePersistentVolumeClaim = "pvcs"
	tableRoles                 = "roles"
	tableClusterRoles          = "cluster_roles"
	tableServices              = "services"
	tableStorageClasses        = "storage_classes"

	Pods                  = "pods"
	Deployments           = "deployments"
	Daemonsets            = "daemonsets"
	Statefulsets          = "statefulsets"
	Namespaces            = "namespaces"
	Ingresses             = "ingresses"
	PersistentVolumeClaim = "persistent-volume-claims"
	Roles                 = "roles"
	ClusterRoles          = "cluster-roles"
	Services              = "services"
	StorageClasses        = "storage-classes"
)

var ResourceTable = map[string]string{Deployments: tableDeployments, Statefulsets: tableStatefulsets, Daemonsets: tableDaemonsets,
	Pods: tablePods, Namespaces: tableNamespaces, Ingresses: tableIngresses, PersistentVolumeClaim: tablePersistentVolumeClaim, Roles: tableRoles,
	Services: tableServices, StorageClasses: tableStorageClasses, ClusterRoles: tableClusterRoles}

type Annotation map[string]string

//
//func (annotation *Annotation)Scan(val interface{}) error{
//	switch val := val.(type){
//	case string:
//		return json.Unmarshal([]byte(val), annotation)
//	case []byte:
//		return json.Unmarshal(val, annotation)
//	default:
//		return errors.New("not support")
//	}
//	return nil
//}
//
//func (annotation *Annotation)Value() (driver.Value, error){
//	bytes, err := json.Marshal(annotation)
//	return string(bytes), err
//}

type Deployment struct {
	Name      string `gorm:"primary_key" json:"name"`
	Namespace string `gorm:"primary_key" json:"namespace"`
	App       string `json:"app,omitempty"`

	Available     int32      `json:"available"`
	Desire        int32      `json:"desire"`
	Status        string     `json:"status"`
	AnnotationStr string     `gorm:"type:text" json:"annotationStr,omitempty"`
	Annotation    Annotation `gorm:"-" json:"annotation"`
	UpdateTime    time.Time  `gorm:"column:updateTime" json:"updateTime,omitempty"`
}

func (Deployment) TableName() string {
	return tableDeployments
}

type Statefulset struct {
	Name      string `gorm:"primary_key" json:"name,omitempty"`
	Namespace string `gorm:"primary_key" json:"namespace,omitempty"`
	App       string `json:"app,omitempty"`

	Available     int32      `json:"available"`
	Desire        int32      `json:"desire"`
	Status        string     `json:"status"`
	AnnotationStr string     `gorm:"type:text" json:"annotationStr,omitempty"`
	Annotation    Annotation `gorm:"-" json:"annotation"`
	CreateTime    time.Time  `gorm:"column:createTime" json:"createTime,omitempty"`
}

func (Statefulset) TableName() string {
	return tableStatefulsets
}

type Daemonset struct {
	Name      string `gorm:"primary_key" json:"name,omitempty"`
	Namespace string `gorm:"primary_key" json:"namespace,omitempty"`
	App       string `json:"app,omitempty"`

	Available     int32      `json:"available"`
	Desire        int32      `json:"desire"`
	Status        string     `json:"status"`
	NodeSelector  string     `json:"nodeSelector, omitempty"`
	AnnotationStr string     `gorm:"type:text" json:"annotationStr,omitempty"`
	Annotation    Annotation `gorm:"-" json:"annotation"`
	CreateTime    time.Time  `gorm:"column:createTime" json:"createTime,omitempty"`
}

func (Daemonset) TableName() string {
	return tableDaemonsets
}

type Service struct {
	Name        string `gorm:"primary_key" json:"name"`
	Namespace   string `gorm:"primary_key" json:"namespace"`
	ServiceType string `json:"type,omitempty"`

	VirtualIp  string `json:"virtualIp,omitempty"`
	ExternalIp string `json:"externalIp,omitempty"`

	Ports         string     `json:"ports,omitempty"`
	AnnotationStr string     `gorm:"type:text" json:"annotationStr,omitempty"`
	Annotation    Annotation `gorm:"-" json:"annotation"`
	CreateTime    time.Time  `gorm:"column:createTime" json:"createTime,omitempty"`
}

func (Service) TableName() string {
	return tableServices
}

type Pvc struct {
	Name             string     `gorm:"primary_key" json:"name"`
	Namespace        string     `gorm:"primary_key" json:"namespace"`
	Status           string     `json:"status,omitempty"`
	Capacity         string     `json:"capacity,omitempty"`
	AccessMode       string     `json:"accessMode,omitempty"`
	AnnotationStr    string     `gorm:"type:text" json:"annotationStr,omitempty"`
	Annotation       Annotation `gorm:"-" json:"annotation"`
	CreateTime       time.Time  `gorm:"column:createTime" json:"createTime,omitempty"`
	StorageClassName string     `gorm:"column:storage_class" json:"storage_class,omitempty"`
}

func (Pvc) TableName() string {
	return tablePersistentVolumeClaim
}

type Ingress struct {
	Name           string     `gorm:"primary_key" json:"name"`
	Namespace      string     `gorm:"primary_key" json:"namespace"`
	Ip             string     `json:"ip,omitempty"`
	TlsTermination string     `json:"tlsTermination,omitempty"`
	AnnotationStr  string     `gorm:"type:text" json:"annotationStr,omitempty"`
	Annotation     Annotation `gorm:"-" json:"annotation"`
	CreateTime     time.Time  `gorm:"column:createTime" json:"createTime,omitempty"`
}

func (Ingress) TableName() string {
	return tableIngresses
}

type Pod struct {
	Name          string      `gorm:"primary_key" json:"name"`
	Namespace     string      `gorm:"primary_key" json:"namespace"`
	Status        string      `json:"status,omitempty"`
	Node          string      `json:"node,omitempty"`
	NodeIp        string      `json:"nodeIp,omitempty"`
	PodIp         string      `json:"podIp,omitempty"`
	ContainerStr  string      `gorm:"type:text" json:",omitempty"`
	Containers    []Container `json:"containers,omitempty"`
	AnnotationStr string      `gorm:"type:text" json:"annotationStr,omitempty"`
	Annotation    Annotation  `gorm:"-" json:"annotation"`
	CreateTime    time.Time   `gorm:"column:createTime" json:"createTime,omitempty"`
}

type Container struct {
	Name  string             `json:"name"`
	Ready bool               `json:"ready,omitempty"`
	Image string             `json:"image"`
	Ports []v1.ContainerPort `json:"ports"`
}

func (Pod) TableName() string {
	return tablePods
}

type Role struct {
	Name          string     `gorm:"primary_key" json:"name"`
	Namespace     string     `gorm:"primary_key" json:"namespace"`
	AnnotationStr string     `gorm:"type:text" json:"annotationStr,omitempty"`
	Annotation    Annotation `gorm:"-" json:"annotation"`
	CreateTime    time.Time  `gorm:"column:createTime" json:"createTime,omitempty"`
}

func (Role) TableName() string {
	return tableRoles
}

type ClusterRole struct {
	Name          string     `gorm:"primary_key" json:"name"`
	AnnotationStr string     `gorm:"type:text" json:"annotationStr,omitempty"`
	Annotation    Annotation `gorm:"-" json:"annotation"`
	CreateTime    time.Time  `gorm:"column:createTime" json:"createTime,omitempty"`
}

func (ClusterRole) TableName() string {
	return tableClusterRoles
}

type Namespace struct {
	Name    string `gorm:"primary_key" json:"name"`
	Creator string `json:"creator,omitempty"`
	Status  string `json:"status"`

	Descrition    string     `json:"description,omitempty"`
	AnnotationStr string     `gorm:"type:text" json:"annotationStr,omitempty"`
	Annotation    Annotation `gorm:"-" json:"annotation"`
	CreateTime    time.Time  `gorm:"column:createTime" json:"createTime,omitempty"`
}

func (Namespace) TableName() string {
	return tableNamespaces
}

type StorageClass struct {
	Name          string     `gorm:"primary_key" json:"name"`
	Creator       string     `json:"creator,omitempty"`
	AnnotationStr string     `gorm:"type:text" json:"annotationStr,omitempty"`
	Annotation    Annotation `gorm:"-" json:"annotation"`
	CreateTime    time.Time  `gorm:"column:createTime" json:"createTime,omitempty"`
	IsDefault     bool       `json:"default"`
	Count         int        `json:"count"`
}

func (StorageClass) TableName() string {
	return tableStorageClasses
}

type Paging struct {
	Limit, Offset int
}

type Controller interface {
	listAndWatch()
	chanStop() chan struct{}
	chanAlive() chan struct{}
	Count(conditions string) int
	ListWithConditions(condition string, paging *Paging) (int, interface{}, error)
}

type CommonAttribute struct {
	K8sClient *kubernetes.Clientset
	DB        *gorm.DB
	stopChan  chan struct{}
	aliveChan chan struct{}
}

func (ca *CommonAttribute) chanStop() chan struct{} {

	return ca.stopChan
}

func (ca *CommonAttribute) chanAlive() chan struct{} {

	return ca.aliveChan
}

type DeploymentCtl struct {
	CommonAttribute
}

type StatefulsetCtl struct {
	CommonAttribute
}

type DaemonsetCtl struct {
	CommonAttribute
}

type ServiceCtl struct {
	CommonAttribute
}

type PvcCtl struct {
	CommonAttribute
}

type PodCtl struct {
	CommonAttribute
}

type IngressCtl struct {
	CommonAttribute
}

type NamespaceCtl struct {
	CommonAttribute
}

type StorageClassCtl struct {
	CommonAttribute
}

type RoleCtl struct {
	CommonAttribute
}

type ClusterRoleCtl struct {
	CommonAttribute
}

func listWithConditions(db *gorm.DB, total *int, object, list interface{}, conditions string, paging *Paging, order string) {
	if len(conditions) == 0 {
		db.Model(object).Count(total)
	} else {
		db.Model(object).Where(conditions).Count(total)
	}

	if paging != nil {
		if len(conditions) > 0 {
			db.Where(conditions).Order(order).Limit(paging.Limit).Offset(paging.Offset).Find(list)
		} else {
			db.Order(order).Limit(paging.Limit).Offset(paging.Offset).Find(list)
		}

	} else {
		if len(conditions) > 0 {
			db.Where(conditions).Order(order).Find(list)
		} else {
			db.Order(order).Find(list)
		}
	}
}

func countWithConditions(db *gorm.DB, conditions string, object interface{}) int {
	var count int
	if len(conditions) == 0 {
		db.Model(object).Count(&count)
	} else {
		db.Model(object).Where(conditions).Count(&count)
	}
	return count
}
