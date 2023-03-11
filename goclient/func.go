package goclient

import (
	"context"
	"fmt"

	"github.com/FUnigrad/funiverse-workspace-service/goclient/template"
	"github.com/FUnigrad/funiverse-workspace-service/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type Template = unstructured.Unstructured

func (client *GoClient) CreateWorkspace(workspace model.Workspace) error {
	_, err := client.CreateNamespace(workspace.Code)

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = client.CreateConfigMap(workspace.Code)
	if err != nil {
		fmt.Println(err.Error())
	}

	volumeConfig := template.VolumeConfig{
		Name:    workspace.Code,
		Storage: "1",
	}

	err = client.CreateVolume(volumeConfig)

	if err != nil {
		fmt.Println(err.Error())
	}
	err = client.CreateMySql(workspace.Code)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (client *GoClient) CreateNamespace(name string) (*Template, error) {
	config := template.NamespaceConfig{
		Name: name,
	}
	namespaceTemplate := template.NewNamespaceTemplate(config)

	result, err := client.Client.Resource(namespaceTemplate.NamespaceRes).Create(
		context.TODO(),
		namespaceTemplate.NamespaceSchema,
		metav1.CreateOptions{},
	)
	return result, err
}

func (client *GoClient) CreateConfigMap(namespace string) (*Template, error) {

	config := template.ConfigMapConfig{
		Name: namespace,
	}

	configmapTemplate := template.NewConfigMapTemplate(config)

	result, err := client.Client.Resource(configmapTemplate.ConfigMapRes).Namespace(namespace).Create(
		context.TODO(),
		configmapTemplate.ConfigMapSchema,
		metav1.CreateOptions{},
	)
	return result, err
}

func (client *GoClient) CreateVolume(config template.VolumeConfig) error {

	pvTemplate := template.NewPersitentVolumeTemplate(config)

	_, err := client.Client.Resource(pvTemplate.PvRes).Create(
		context.TODO(),
		pvTemplate.PvSchema,
		metav1.CreateOptions{},
	)

	if err != nil {
		return err
	}

	pvcTemplate := template.NewPersitentVolumeClaimTemplate(config)

	_, err = client.Client.Resource(pvcTemplate.PvcRes).Namespace(config.Name).Create(
		context.TODO(),
		pvcTemplate.PvcSchema,
		metav1.CreateOptions{},
	)

	return err
}

func (client *GoClient) CreateMySql(namespace string) error {
	mySqlTemplate := template.NewMySqlTemplate()

	_, err := client.Client.Resource(template.CreateDeploymentResource()).Namespace(namespace).Create(
		context.TODO(),
		mySqlTemplate.Deploy,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}
	_, err = client.Client.Resource(template.CreateServiceResource()).Namespace(namespace).Create(
		context.TODO(),
		mySqlTemplate.Service,
		metav1.CreateOptions{},
	)

	return err
}
