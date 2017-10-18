/*
Copyright (c) 2017 GigaSpaces Technologies Ltd. All rights reserved

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

package cloudifyprovider

import (
	cloudify "github.com/cloudify-incubator/cloudify-rest-go-client/cloudify"
	"github.com/golang/glog"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
)

type CloudifyScaleProvider struct {
	client       *cloudify.CloudifyClient
	deploymentID string
}

// Name returns name of the cloud provider.
func (clsp *CloudifyScaleProvider) Name() string {
	glog.Warning("Name")
	return "cloudify"
}

// NodeGroups returns all node groups configured for this cloud provider.
func (clsp *CloudifyScaleProvider) NodeGroups() []cloudprovider.NodeGroup {
	glog.Warning("NodeGroups")
	nodes := []cloudprovider.NodeGroup{}

	// get all nodes with type=="kubernetes_host"
	params := map[string]string{}
	params["deployment_id"] = clsp.deploymentID
	cloud_nodes := clsp.client.GetNodes(params)
	for _, node := range cloud_nodes.Items {
		var not_kubernetes_host bool = true
		for _, type_name := range node.TypeHierarchy {
			if type_name == "kubernetes_host" {
				not_kubernetes_host = false
				break
			}
		}

		if not_kubernetes_host {
			continue
		}

		nodes = append(nodes, CreateNodeGroup(clsp.client, clsp.deploymentID, node.Id))
	}
	glog.Warningf("Nodes result %+v", nodes)
	return nodes
}

// NodeGroupForNode returns the node group for the given node.
func (clsp *CloudifyScaleProvider) NodeGroupForNode(node *apiv1.Node) (cloudprovider.NodeGroup, error) {
	glog.Warningf("NodeGroupForNode %+v", node)
	return nil, cloudprovider.ErrNotImplemented
}

// Pricing returns pricing model for this cloud provider or error if not available.
func (clsp *CloudifyScaleProvider) Pricing() (cloudprovider.PricingModel, errors.AutoscalerError) {
	glog.Warning("Pricing")
	return nil, cloudprovider.ErrNotImplemented
}

// GetAvailableMachineTypes get all machine types that can be requested from the cloud provider.
// Implementation optional.
func (clsp *CloudifyScaleProvider) GetAvailableMachineTypes() ([]string, error) {
	glog.Warning("GetAvailableMachineTypes")
	return []string{}, cloudprovider.ErrNotImplemented
}

// NewNodeGroup builds a theoretical node group based on the node definition provided. The node group is not automatically
// created on the cloud provider side. The node group is not returned by NodeGroups() until it is created.
func (clsp *CloudifyScaleProvider) NewNodeGroup(machineType string, labels map[string]string, extraResources map[string]resource.Quantity) (cloudprovider.NodeGroup, error) {
	glog.Warningf("NewNodeGroup: %+v %+v %+v", machineType, labels, extraResources)
	return nil, cloudprovider.ErrNotImplemented
}
