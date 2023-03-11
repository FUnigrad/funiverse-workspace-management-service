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

func (client *GoClient) CreateWorkspace(workspace model.Workspace) (err error) {
	err = client.CreateNamespace(workspace.Code)

	if err != nil {
		return
	}

	err = client.CreateConfigMap(workspace.Code)
	if err != nil {
		return
	}

	volumeConfig := template.VolumeConfig{
		Name:    workspace.Code,
		Storage: "1",
	}

	err = client.CreateVolume(volumeConfig)

	if err != nil {
		return
	}
	err = client.CreateMySql(workspace.Code)
	if err != nil {
		return
	}
	err = client.CreateAppService(workspace.Code)
	if err != nil {
		return
	}

	err = client.CreateIngress(workspace.Code)
	if err != nil {
		return
	}
	return err
}

func (client *GoClient) CreateNamespace(name string) (err error) {
	config := template.NamespaceConfig{
		Name: name,
	}

	_, err = client.Client.Resource(template.NewNameSpaceResource()).Create(
		context.TODO(),
		template.NewNamespaceTemplate(config),
		metav1.CreateOptions{},
	)
	return
}

func (client *GoClient) CreateConfigMap(namespace string) (err error) {

	config := template.ConfigMapConfig{
		Name: namespace,
	}

	configmapTemplate := template.NewConfigMapTemplate(config)

	_, err = client.Client.Resource(configmapTemplate.ConfigMapRes).Namespace(namespace).Create(
		context.TODO(),
		configmapTemplate.ConfigMapSchema,
		metav1.CreateOptions{},
	)
	return
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

func (client *GoClient) CreateAppService(namespace string) error {
	appServiceTemplate := template.NewAppServiceTemplate()

	_, err := client.Client.Resource(template.CreateDeploymentResource()).Namespace(namespace).Create(
		context.TODO(),
		appServiceTemplate.Deploy,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}
	_, err = client.Client.Resource(template.CreateServiceResource()).Namespace(namespace).Create(
		context.TODO(),
		appServiceTemplate.Service,
		metav1.CreateOptions{},
	)

	return err
}

func (client *GoClient) CreateIngress(namespace string) error {

	ingressTemplate := template.NewIngressTemplate(namespace)

	_, err := client.Client.Resource(template.CreateIngressResource()).Namespace(namespace).Create(
		context.TODO(),
		ingressTemplate.AppService,
		metav1.CreateOptions{},
	)

	if err != nil {
		return err
	}
	_, err = client.Client.Resource(template.CreateIngressResource()).Namespace("frontend").Create(
		context.TODO(),
		ingressTemplate.WorkspaceWebApp,
		metav1.CreateOptions{},
	)

	if err != nil {
		return err
	}

	_, err = client.Client.Resource(template.CreateIngressResource()).Namespace("frontend").Create(
		context.TODO(),
		ingressTemplate.WorkspaceAdminWebApp,
		metav1.CreateOptions{},
	)

	return err

}

func (client *GoClient) DeleteWorkspace(workspace model.Workspace) (err error) {

	namespace := workspace.Code

	deletePolicy := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}

	//Delete All Namespace resource
	err = client.Client.Resource(
		template.NewNameSpaceResource(),
	).Delete(
		context.TODO(),
		namespace,
		deleteOptions,
	)
	if err != nil {
		return
	}
	//Delete 2 Frontend Ingress
	err = client.Client.Resource(
		template.CreateIngressResource(),
	).Namespace("frontend").Delete(
		context.TODO(),
		fmt.Sprintf("%s-workspace-ingress", namespace),
		deleteOptions,
	)

	if err != nil {
		return
	}

	err = client.Client.Resource(
		template.CreateIngressResource(),
	).Namespace("frontend").Delete(
		context.TODO(),
		fmt.Sprintf("%s-admin-ingress", namespace),
		deleteOptions,
	)

	if err != nil {
		return
	}

	err = client.Client.Resource(
		template.NewPersitentVolumeResource(),
	).Delete(
		context.TODO(),
		fmt.Sprintf("pv-for-%s", namespace),
		deleteOptions,
	)
	return
}
