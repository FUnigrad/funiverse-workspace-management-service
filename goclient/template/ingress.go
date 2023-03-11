package template

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type IngressTemplate struct {
	AppService           *unstructured.Unstructured
	WorkspaceWebApp      *unstructured.Unstructured
	WorkspaceAdminWebApp *unstructured.Unstructured
}

func NewIngressTemplate(name string) IngressTemplate {
	return IngressTemplate{
		WorkspaceAdminWebApp: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "k8s.nginx.org/v1",
				"kind":       "VirtualServer",
				"metadata": map[string]interface{}{
					"name":      fmt.Sprintf("%s-admin-ingress", name),
					"namespace": "frontend",
				},
				"spec": map[string]interface{}{
					"host": fmt.Sprintf("admin.%s.funiverse.world", name),
					"upstreams": []map[string]interface{}{
						{
							"name":    "workspace-admin-web-app",
							"service": "workspace-admin-web-app",
							"port":    80,
						},
					},
					"routes": []map[string]interface{}{
						{
							"path": "/",
							"action": map[string]interface{}{
								"pass": "workspace-admin-web-app",
							},
						},
					},
				},
			},
		},
		WorkspaceWebApp: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "k8s.nginx.org/v1",
				"kind":       "VirtualServer",
				"metadata": map[string]interface{}{
					"name":      fmt.Sprintf("%s-workspace-ingress", name),
					"namespace": "frontend",
				},
				"spec": map[string]interface{}{
					"host": fmt.Sprintf("%s.funiverse.world", name),
					"upstreams": []map[string]interface{}{
						{
							"name":    "workspace-web-app",
							"service": "workspace-web-app",
							"port":    80,
						},
					},
					"routes": []map[string]interface{}{
						{
							"path": "/",
							"action": map[string]interface{}{
								"pass": "workspace-web-app",
							},
						},
					},
				},
			},
		},
		AppService: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "k8s.nginx.org/v1",
				"kind":       "VirtualServer",
				"metadata": map[string]interface{}{
					"name":      fmt.Sprintf("%s-app-service-ingress", name),
					"namespace": name,
				},
				"spec": map[string]interface{}{
					"host": fmt.Sprintf("api.%s.funiverse.world", name),
					"upstreams": []map[string]interface{}{
						{
							"name":    "app-service",
							"service": "app-service",
							"port":    8080,
						},
					},
					"routes": []map[string]interface{}{
						{
							"path": "/",
							"action": map[string]interface{}{
								"pass": "app-service",
							},
						},
					},
				},
			},
		},
	}
}
