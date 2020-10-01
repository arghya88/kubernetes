/*
Copyright 2020 The Kubernetes Authors.

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

package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LeaderMigrationConfiguration provides versioned configuration for all migrating leader locks.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type LeaderMigrationConfiguration struct {
	metav1.TypeMeta

	// LeaderName is the name of the leader election resource that protects the migration
	// E.g. 1-20-KCM-to-1-21-CCM
	LeaderName string

	// ResourceLock indicates the resource object type that will be used to lock
	// Should be "leases" or "endpoints"
	ResourceLock string

	// ControllerLeaders contains a list of migrating leader lock configurations
	ControllerLeaders []ControllerLeaderConfiguration
}

// ControllerLeaderConfiguration provides the configuration for a migrating leader lock.
type ControllerLeaderConfiguration struct {
	// Name is the name of the controller being migrated
	// E.g. service-controller, route-controller, cloud-node-controller, etc
	Name string

	// Component is the name of the component in which the controller should be running.
	// E.g. kube-controller-manager, cloud-controller-manager, etc
	Component string
}
