package template

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewNameSpaceResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "namespaces",
	}
}

type NamespaceConfig struct {
	Name string `json:"name"`
}

func NewNamespaceTemplate(config NamespaceConfig) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"Kind":       "Namespace",
			"metadata": map[string]interface{}{
				"name": config.Name,
			},
		},
	}
}
