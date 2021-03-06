// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistry":                                     schema_pkg_apis_apicur_v1alpha1_ApicurioRegistry(ref),
		"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpec":                                 schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpec(ref),
		"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecConfiguration":                    schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecConfiguration(ref),
		"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecConfigurationDataSource":          schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecConfigurationDataSource(ref),
		"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecConfigurationKafka":               schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecConfigurationKafka(ref),
		"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecConfigurationStreams":             schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecConfigurationStreams(ref),
		"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecDeployment":                       schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecDeployment(ref),
		"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecDeploymentResources":              schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecDeploymentResources(ref),
		"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecDeploymentResourcesRequestsLimit": schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecDeploymentResourcesRequestsLimit(ref),
		"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecImage":                            schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecImage(ref),
		"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistryStatus":                               schema_pkg_apis_apicur_v1alpha1_ApicurioRegistryStatus(ref),
	}
}

func schema_pkg_apis_apicur_v1alpha1_ApicurioRegistry(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ApicurioRegistry is the Schema for the apicurioregistries API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistryStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpec", "github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistryStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ApicurioRegistrySpec defines the desired state of ApicurioRegistry",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecImage"),
						},
					},
					"configuration": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecConfiguration"),
						},
					},
					"deployment": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecDeployment"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecConfiguration", "github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecDeployment", "github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecImage"},
	}
}

func schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecConfiguration(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					"persistence": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"dataSource": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecConfigurationDataSource"),
						},
					},
					"kafka": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecConfigurationKafka"),
						},
					},
					"streams": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecConfigurationStreams"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecConfigurationDataSource", "github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecConfigurationKafka", "github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecConfigurationStreams"},
	}
}

func schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecConfigurationDataSource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					"url": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"userName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"password": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecConfigurationKafka(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					"bootstrapServers": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecConfigurationStreams(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					"bootstrapServers": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"applicationServerPort": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"applicationId": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecDeployment(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"route": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecDeploymentResources"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecDeploymentResources"},
	}
}

func schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecDeploymentResources(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					"cpu": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecDeploymentResourcesRequestsLimit"),
						},
					},
					"memory": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecDeploymentResourcesRequestsLimit"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/Apicurio/apicurio-operators/apicurio-registry/pkg/apis/apicur/v1alpha1.ApicurioRegistrySpecDeploymentResourcesRequestsLimit"},
	}
}

func schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecDeploymentResourcesRequestsLimit(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					"requests": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"limit": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_apicur_v1alpha1_ApicurioRegistrySpecImage(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					"registry": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"override": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_apicur_v1alpha1_ApicurioRegistryStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ApicurioRegistryStatus defines the observed state of ApicurioRegistry",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"deploymentName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"serviceName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"ingressName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"replicaCount": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"route": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
