package template

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ConfigMapConfig struct {
	Namespace string
}

func CreateConfigMapResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "configmaps",
	}
}

func CreateConfigMapManifest() *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"Kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name": "config",
			},
			"data": map[string]interface{}{
				"MYSQL_ROOT_PASSWORD":           "root",
				"MYSQL_DATABASE":                "mydb",
				"SPRING_DATASOURCE_URL":         "jdbc:mysql://mysql:3306/mydb",
				"SPRING_DATASOURCE_USERNAME":    "root",
				"SPRING_JPA_HIBERNATE_DDL_AUTO": "update",
			},
		},
	}
}
