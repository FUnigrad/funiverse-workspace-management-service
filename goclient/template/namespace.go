package template

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type NamespaceConfig struct {
	Name string `json:"name"`
}

type NamespaceTemplate struct {
	NamespaceRes    schema.GroupVersionResource
	NamespaceSchema *unstructured.Unstructured
}

func NewNamespaceTemplate(config NamespaceConfig) NamespaceTemplate {
	return NamespaceTemplate{
		NamespaceRes: schema.GroupVersionResource{
			Group:    "",
			Version:  "v1",
			Resource: "namespaces",
		},
		NamespaceSchema: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "v1",
				"Kind":       "Namespace",
				"metadata": map[string]interface{}{
					"name": config.Name,
				},
			},
		},
	}
}
