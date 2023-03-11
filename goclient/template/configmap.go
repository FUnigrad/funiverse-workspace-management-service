package template

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ConfigMapConfig struct {
	Name string
}

type ConfigMapTemplate struct {
	ConfigMapRes    schema.GroupVersionResource
	ConfigMapSchema *unstructured.Unstructured
}

func NewConfigMapTemplate(config ConfigMapConfig) ConfigMapTemplate {
	return ConfigMapTemplate{
		ConfigMapRes: schema.GroupVersionResource{
			Group:    "",
			Version:  "v1",
			Resource: "configmaps",
		},
		ConfigMapSchema: &unstructured.Unstructured{
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
		},
	}
}
