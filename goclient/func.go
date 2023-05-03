package goclient

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/FUnigrad/funiverse-workspace-service/goclient/template"
	"github.com/FUnigrad/funiverse-workspace-service/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type Template = unstructured.Unstructured

func (client *GoClient) CreateWorkspace(workspace model.WorkspaceDTO) (err error) {

	namespace := strings.ToLower(workspace.Code)
	domain := strings.ToLower(workspace.Domain)

	err = client.CreateNamespace(namespace)

	if err != nil {
		return
	}

	err = client.CreateConfigMap(namespace)
	if err != nil {
		return
	}

	volumeConfig := template.VolumeConfig{
		Storage:    2,
		AccessMode: "ReadWriteMany",
	}

	err = client.CreateVolume(namespace, volumeConfig)

	if err != nil {
		return
	}
	err = client.CreateMySql(namespace)
	if err != nil {
		return
	}
	err = client.CreateAppService(namespace)
	if err != nil {
		return
	}

	err = client.CreateIngress(namespace, domain)
	if err != nil {
		return
	}

	for {
		resp, _ := http.Get(fmt.Sprintf("http://api.%s/actuator/health", domain))
		if resp == nil {
			continue
		}
		if resp.StatusCode == 200 {
			break
		}
	}

	return err
}

func (client *GoClient) CreateNamespace(name string) (err error) {
	config := template.NamespaceConfig{
		Name: name,
	}

	_, err = client.Client.Resource(template.CreateNameSpaceResource()).Create(
		context.TODO(),
		template.CreateNamespaceManifest(config),
		metav1.CreateOptions{},
	)
	return
}

func (client *GoClient) CreateConfigMap(namespace string) (err error) {

	_, err = client.Client.Resource(template.CreateConfigMapResource()).Namespace(namespace).Create(
		context.TODO(),
		template.CreateConfigMapManifest(),
		metav1.CreateOptions{},
	)
	return
}

func (client *GoClient) CreateVolume(namespace string, config template.VolumeConfig) (err error) {

	_, err = client.Client.Resource(template.CreatePVCResource()).Namespace(namespace).Create(
		context.TODO(),
		template.CreatePVCManifest(config),
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

func (client *GoClient) CreateIngress(namespace string, domain string) error {

	ingressTemplate := template.NewIngressTemplate(namespace, domain)

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

	namespace := strings.ToLower(workspace.Code)

	deletePolicy := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
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
	//Delete All Namespace resource
	err = client.Client.Resource(
		template.CreateNameSpaceResource(),
	).Delete(
		context.TODO(),
		namespace,
		deleteOptions,
	)
	if err != nil {
		return
	}

	return
}
