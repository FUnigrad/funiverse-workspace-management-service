package template

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type MySqlTemplate struct {
	Deploy  *unstructured.Unstructured
	Service *unstructured.Unstructured
}

func NewMySqlTemplate() MySqlTemplate {
	return MySqlTemplate{
		Deploy: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "apps/v1",
				"kind":       "Deployment",
				"metadata": map[string]interface{}{
					"name": "mysql",
				},
				"spec": map[string]interface{}{
					"selector": map[string]interface{}{
						"matchLabels": map[string]interface{}{
							"app": "mysql",
						},
					},
					"template": map[string]interface{}{
						"metadata": map[string]interface{}{
							"labels": map[string]interface{}{
								"app": "mysql",
							},
						},
						"spec": map[string]interface{}{
							"containers": []map[string]interface{}{
								{
									"image": "mysql",
									"name":  "mysql",
									"env": []map[string]interface{}{
										{
											"name": "MYSQL_ROOT_PASSWORD",
											"valueFrom": map[string]interface{}{
												"configMapKeyRef": map[string]interface{}{
													"name": "config",
													"key":  "MYSQL_ROOT_PASSWORD",
												},
											},
										},
										{
											"name": "MYSQL_DATABASE",
											"valueFrom": map[string]interface{}{
												"configMapKeyRef": map[string]interface{}{
													"name": "config",
													"key":  "MYSQL_DATABASE",
												},
											},
										},
									},
									"ports": []map[string]interface{}{
										{
											"containerPort": 3306,
											"name":          "mysql",
										},
									},
									"volumeMounts": []map[string]interface{}{
										{
											"name":      "mysql-persistent-storage",
											"mountPath": "/var/lib/mysql",
										},
									},
								},
							},
							"volumes": []map[string]interface{}{
								{
									"name": "mysql-persistent-storage",
									"persistentVolumeClaim": map[string]interface{}{
										"claimName": "mysql-volume",
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
					"name": "mysql",
				},
				"spec": map[string]interface{}{
					"ports": []map[string]interface{}{
						{
							"port": 3306,
						},
					},
					"selector": map[string]interface{}{
						"app": "mysql",
					},
				},
			},
		},
	}
}
