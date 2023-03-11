package template

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type VolumeConfig struct {
	Name    string
	Storage string
}

type PersitentVolumeTemplate struct {
	PvRes    schema.GroupVersionResource
	PvSchema *unstructured.Unstructured
}

type PersitentVolumeClaimTemplate struct {
	PvcRes    schema.GroupVersionResource
	PvcSchema *unstructured.Unstructured
}

func NewPersitentVolumeTemplate(config VolumeConfig) PersitentVolumeTemplate {
	return PersitentVolumeTemplate{
		PvRes: schema.GroupVersionResource{
			Group:    "",
			Version:  "v1",
			Resource: "persistentvolumes",
		},
		PvSchema: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "v1",
				"kind":       "PersistentVolume",
				"metadata": map[string]interface{}{
					"name": fmt.Sprintf("pv-for-%s", config.Name),
					"labels": map[string]interface{}{
						"type": "local",
					},
				},
				"spec": map[string]interface{}{
					"storageClassName": "hostpath",
					"capacity": map[string]interface{}{
						"storage": fmt.Sprintf("%sGi", config.Storage),
					},
					"accessModes": []string{
						"ReadWriteOnce",
					},
					"hostPath": map[string]interface{}{
						"path": fmt.Sprintf("/mnt/data/%s", config.Name),
					},
				},
			},
		},
	}
}

func NewPersitentVolumeClaimTemplate(config VolumeConfig) PersitentVolumeClaimTemplate {
	return PersitentVolumeClaimTemplate{
		PvcRes: schema.GroupVersionResource{
			Group:    "",
			Version:  "v1",
			Resource: "persistentvolumeclaims",
		},
		PvcSchema: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "v1",
				"kind":       "PersistentVolumeClaim",
				"metadata": map[string]interface{}{
					"name": "pv-claim",
				},
				"spec": map[string]interface{}{
					"storageClassName": "hostpath",
					"accessModes": []string{
						"ReadWriteOnce",
					},
					"resources": map[string]interface{}{
						"requests": map[string]interface{}{
							"storage": fmt.Sprintf("%sGi", config.Storage),
						},
					},
				},
			},
		},
	}
}
