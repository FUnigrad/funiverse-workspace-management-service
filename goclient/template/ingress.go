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

func NewIngressTemplate(namespace string, domain string) IngressTemplate {
	return IngressTemplate{
		WorkspaceAdminWebApp: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "k8s.nginx.org/v1",
				"kind":       "VirtualServer",
				"metadata": map[string]interface{}{
					"name":      fmt.Sprintf("%s-admin-ingress", namespace),
					"namespace": "frontend",
				},
				"spec": map[string]interface{}{
					"host": fmt.Sprintf("admin.%s", domain),
					// "tls": map[string]interface{}{
					// 	"secret": fmt.Sprintf("admin-%s-ssl", strings.Join(strings.Split(domain, "."), "-")),
					// 	"cert-manager": map[string]interface{}{
					// 		"cluster-issuer": "letsencrypt-prod",
					// 	},
					// },
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
					"name":      fmt.Sprintf("%s-workspace-ingress", namespace),
					"namespace": "frontend",
				},
				"spec": map[string]interface{}{
					"host": domain,
					// "tls": map[string]interface{}{
					// 	"secret": fmt.Sprintf("%s-ssl", strings.Join(strings.Split(domain, "."), "-")),
					// 	"cert-manager": map[string]interface{}{
					// 		"cluster-issuer": "letsencrypt-prod",
					// 	},
					// },
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
					"name":      fmt.Sprintf("%s-app-service-ingress", namespace),
					"namespace": namespace,
				},
				"spec": map[string]interface{}{
					"host": fmt.Sprintf("api.%s", domain),
					// "tls": map[string]interface{}{
					// 	"secret": fmt.Sprintf("api-%s-ssl", strings.Join(strings.Split(domain, "."), "-")),
					// 	"cert-manager": map[string]interface{}{
					// 		"cluster-issuer": "letsencrypt-prod",
					// 	},
					// },
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
