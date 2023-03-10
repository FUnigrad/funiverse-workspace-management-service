package goclient

import (
	"context"

	"github.com/FUnigrad/funiverse-workspace-service/model"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (client *GoClient) GetPodsName() ([]v1.Pod, error) {

	clientset := client.ClientSet

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	podList := pods.Items

	return podList, nil
}

func (client *GoClient) CreateWorkspace(workspace model.Workspace) error {
	_, err := client.CreateNamespace(workspace.Name)

	if err != nil {
		return err
	}

	return nil
}

func (client *GoClient) CreateNamespace(name string) (*v1.Namespace, error) {
	clientset := client.ClientSet
	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	namespace, err := clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	return namespace, err
}
