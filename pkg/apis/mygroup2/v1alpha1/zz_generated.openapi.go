// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"mydomain.com/mygroup2/kata-operator/pkg/apis/mygroup2/v1alpha1.KataSet":       schema_pkg_apis_mygroup2_v1alpha1_KataSet(ref),
		"mydomain.com/mygroup2/kata-operator/pkg/apis/mygroup2/v1alpha1.KataSetSpec":   schema_pkg_apis_mygroup2_v1alpha1_KataSetSpec(ref),
		"mydomain.com/mygroup2/kata-operator/pkg/apis/mygroup2/v1alpha1.KataSetStatus": schema_pkg_apis_mygroup2_v1alpha1_KataSetStatus(ref),
	}
}

func schema_pkg_apis_mygroup2_v1alpha1_KataSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KataSet is the Schema for the katasets API",
				Type:        []string{"object"},
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
							Ref: ref("mydomain.com/mygroup2/kata-operator/pkg/apis/mygroup2/v1alpha1.KataSetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("mydomain.com/mygroup2/kata-operator/pkg/apis/mygroup2/v1alpha1.KataSetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "mydomain.com/mygroup2/kata-operator/pkg/apis/mygroup2/v1alpha1.KataSetSpec", "mydomain.com/mygroup2/kata-operator/pkg/apis/mygroup2/v1alpha1.KataSetStatus"},
	}
}

func schema_pkg_apis_mygroup2_v1alpha1_KataSetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KataSetSpec defines the desired state of KataSet",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_mygroup2_v1alpha1_KataSetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KataSetStatus defines the observed state of KataSet",
				Type:        []string{"object"},
			},
		},
	}
}
