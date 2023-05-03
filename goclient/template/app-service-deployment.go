package template

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type AppServiceTemplate struct {
	Deploy  *unstructured.Unstructured
	Service *unstructured.Unstructured
}

func NewAppServiceTemplate() AppServiceTemplate {
	return AppServiceTemplate{
		Deploy: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "apps/v1",
				"kind":       "Deployment",
				"metadata": map[string]interface{}{
					"name": "app-service",
				},
				"spec": map[string]interface{}{
					"replicas":             1,
					"revisionHistoryLimit": 0,
					"minReadySeconds":      10,
					"selector": map[string]interface{}{
						"matchLabels": map[string]interface{}{
							"type": "service",
							"name": "app-service",
						},
					},
					"template": map[string]interface{}{
						"metadata": map[string]interface{}{
							"labels": map[string]interface{}{
								"type": "service",
								"name": "app-service",
							},
						},
						"spec": map[string]interface{}{
							"containers": []map[string]interface{}{
								{
									"image": "funiverse/app-service:latest",
									"name":  "application",
									"env": []map[string]interface{}{
										{
											"name": "SPRING_DATASOURCE_URL",
											"valueFrom": map[string]interface{}{
												"configMapKeyRef": map[string]interface{}{
													"name": "config",
													"key":  "SPRING_DATASOURCE_URL",
												},
											},
										},
										{
											"name": "SPRING_DATASOURCE_USERNAME",
											"valueFrom": map[string]interface{}{
												"configMapKeyRef": map[string]interface{}{
													"name": "config",
													"key":  "SPRING_DATASOURCE_USERNAME",
												},
											},
										},
										{
											"name": "SPRING_DATASOURCE_PASSWORD",
											"valueFrom": map[string]interface{}{
												"configMapKeyRef": map[string]interface{}{
													"name": "config",
													"key":  "MYSQL_ROOT_PASSWORD",
												},
											},
										},
										{
											"name": "SPRING_JPA_HIBERNATE_DDL_AUTO",
											"valueFrom": map[string]interface{}{
												"configMapKeyRef": map[string]interface{}{
													"name": "config",
													"key":  "SPRING_JPA_HIBERNATE_DDL_AUTO",
												},
											},
										},
									},
									"ports": []map[string]interface{}{
										{
											"containerPort": 8080,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		Service: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "v1",
				"kind":       "Service",
				"metadata": map[string]interface{}{
					"name": "app-service",
				},
				"spec": map[string]interface{}{
					"ports": []map[string]interface{}{
						{
							"port":       8080,
							"targetPort": 8080,
							"name":       "http",
						},
					},
					"selector": map[string]interface{}{
						"type": "service",
						"name": "app-service",
					},
				},
			},
		},
	}
}
