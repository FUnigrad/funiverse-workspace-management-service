package template

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type NamespaceConfig struct {
	Name string `json:"name"`
}

type NamespaceTemplate struct {
	NamespaceResult schema.GroupVersionKind
	Namespace       unstructured.Unstructured
}

func (namespace *NamespaceTemplate) New(config NamespaceConfig) {
	namespace.NamespaceResult = schema.GroupVersionKind{
		Group:   "core",
		Version: "v1",
		Kind:    "Namespace",
	}
	namespace.Namespace = unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"Kind":       "Namespace",
			"metadata": map[string]interface{}{
				"name": config.Name,
			},
		},
	}
}
