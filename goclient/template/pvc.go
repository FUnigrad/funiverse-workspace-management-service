package template

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type VolumeConfig struct {
	Storage    int
	AccessMode string
}

func CreatePVCResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "persistentvolumeclaims",
	}
}

func CreatePVCManifest(config VolumeConfig) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "PersistentVolumeClaim",
			"metadata": map[string]interface{}{
				"name": "mysql-volume",
			},
			"spec": map[string]interface{}{
				"storageClassName": "longhorn",
				"accessModes": []string{
					config.AccessMode,
				},
				"resources": map[string]interface{}{
					"requests": map[string]interface{}{
						"storage": fmt.Sprintf("%dGi", config.Storage),
					},
				},
			},
		},
	}
}
