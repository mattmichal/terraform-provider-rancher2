// Copyright 2020 Oracle and/or its affiliates.

package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	clusterOKEKind   = "oke"
	clusterDriverOKE = "oraclecontainerengine"
)

//Types

type OracleKubernetesEngineConfig struct {
	CompartmentID               string `json:"compartmentId,omitempty" yaml:"compartmentId,omitempty"`
	Description                 string `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName                 string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	EnableKubernetesDashboard   bool   `json:"enableKubernetesDashboard,omitempty" yaml:"enableKubernetesDashboard,omitempty"`
	Fingerprint                 string `json:"fingerprint,omitempty" yaml:"fingerprint,omitempty"`
	KubernetesVersion           string `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
	DriverName                  string `json:"driverName,omitempty" yaml:"driverName,omitempty"`
	Name                        string `json:"name,omitempty" yaml:"name,omitempty"`
	NodeImage                   string `json:"nodeImage,omitempty" yaml:"nodeImage,omitempty"`
	NodePoolSubnetDNSDomainName string `json:"nodePoolDnsDomainName,omitempty" yaml:"nodePoolDnsDomainName,omitempty"`
	NodePoolSubnetName          string `json:"nodePoolSubnetName,omitempty" yaml:"nodePoolSubnetName,omitempty"`
	NodePublicSSHKeyContents    string `json:"nodePublicKeyContents,omitempty" yaml:"nodePublicKeyContents,omitempty"`
	NodeShape                   string `json:"nodeShape,omitempty" yaml:"nodeShape,omitempty"`
	PrivateKeyContents          string `json:"privateKeyContents,omitempty" yaml:"privateKeyContents,omitempty"`
	PrivateKeyPassphrase        string `json:"privateKeyPassphrase,omitempty" yaml:"privateKeyPassphrase,omitempty"`
	PrivateNodes                bool   `json:"enablePrivateNodes,omitempty" yaml:"enablePrivateNodes,omitempty"`
	QuantityOfSubnets           int64  `json:"quantityOfNodeSubnets,omitempty" yaml:"quantityOfNodeSubnets,omitempty"`
	QuantityPerSubnet           int64  `json:"quantityPerSubnet,omitempty" yaml:"quantityPerSubnet,omitempty"`
	Region                      string `json:"region,omitempty" yaml:"region,omitempty"`
	ServiceLBSubnet1Name        string `json:"loadBalancerSubnetName1,omitempty" yaml:"loadBalancerSubnetName1,omitempty"`
	ServiceLBSubnet2Name        string `json:"loadBalancerSubnetName2,omitempty" yaml:"loadBalancerSubnetName2,omitempty"`
	ServiceSubnetDNSDomainName  string `json:"serviceDnsDomainName,omitempty" yaml:"serviceDnsDomainName,omitempty"`
	SkipVCNDelete               bool   `json:"skipVcnDelete,omitempty" yaml:"skipVcnDelete,omitempty"`
	TenancyID                   string `json:"tenancyId,omitempty" yaml:"tenancyId,omitempty"`
	UserOCID                    string `json:"userOcid,omitempty" yaml:"userOcid,omitempty"`
	VCNName                     string `json:"vcnName,omitempty" yaml:"vcnName,omitempty"`
	VcnCompartmentID            string `json:"vcnCompartmentId,omitempty" yaml:"vcnCompartmentId,omitempty"`
	WorkerNodeIngressCidr       string `json:"workerNodeIngressCidr,omitempty" yaml:"workerNodeIngressCidr,omitempty"`
}

//Schemas

func clusterOKEConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{

		"compartment_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The OCID of the compartment in which to create resources (VCN, worker nodes, etc.)",
		},
		"fingerprint": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The fingerprint corresponding to the specified user's private API Key",
		},
		"kubernetes_version": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The Kubernetes version that will be used for your master *and* worker nodes e.g. v1.17.9",
		},
		"node_image": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The OS for the node image",
		},
		"node_shape": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The shape of the node (determines number of CPUs and  amount of memory on each node)",
		},
		"private_key_contents": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "The private API key file contents for the specified user, in PEM format",
		},
		"region": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The availability domain within the region to host the OKE cluster",
		},
		"tenancy_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The OCID of the tenancy in which to create resources",
		},
		"user_ocid": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The OCID of a user who has access to the tenancy/compartment",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An optional description of this cluster",
		},
		"enable_kubernetes_dashboard": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Enable the kubernetes dashboard",
		},
		"enable_private_nodes": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether worker nodes are deployed into a new private subnet",
		},
		"load_balancer_subnet_name_1": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the first existing subnet to use for Kubernetes services / LB",
		},
		"load_balancer_subnet_name_2": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The (optional) name of a second existing subnet to use for Kubernetes services / LB",
		},
		"node_pool_dns_domain_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "nodedns",
			Description: "Optional name for DNS domain of node pool subnet",
		},
		"node_pool_subnet_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "nodedns",
			Description: "Optional name for node pool subnet",
		},
		"node_public_key_contents": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The contents of the SSH public key file to use for the nodes",
		},
		"private_key_passphrase": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "The passphrase of the private key for the OKE cluster",
		},
		"quantity_of_node_subnets": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Number of node subnets (defaults to creating 1 regional subnet)",
		},
		"quantity_per_subnet": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     1,
			Description: "Number of worker nodes in each subnet / availability domain",
		},
		"service_dns_domain_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "svcdns",
			Description: "Optional name for DNS domain of service subnet",
		},
		"skip_vcn_delete": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether to skip deleting VCN",
		},
		"vcn_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The optional name of an existing virtual network to use for the cluster creation. A new VCN will be created if not specified.",
		},
		"worker_node_ingress_cidr": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Additional CIDR from which to allow ingress to worker nodes",
		},
	}

	return s
}
