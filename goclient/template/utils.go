package template

import "k8s.io/apimachinery/pkg/runtime/schema"

func CreateDeploymentResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}
}

func CreateIngressResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "k8s.nginx.org",
		Version:  "v1",
		Resource: "virtualservers",
	}
}

func CreateServiceResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "services",
	}
}
