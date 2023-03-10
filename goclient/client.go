package goclient

import (
	"path/filepath"

	"github.com/FUnigrad/funiverse-workspace-service/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type GoClient struct {
	ClientSet *kubernetes.Clientset
}

func NewClient(config config.Config) (GoClient, error) {
	var clientset *kubernetes.Clientset
	var err error
	if config.Enviroment == "local" {
		clientset, err = NewOutClusterClient()

	} else if config.Enviroment == "prod" {
		clientset, err = NewInClusterClient()
	} else {
		return GoClient{
			ClientSet: nil,
		}, clientcmd.ErrEmptyConfig
	}

	if err != nil {
		return GoClient{
			ClientSet: nil,
		}, err
	}

	return GoClient{
		ClientSet: clientset,
	}, nil
}

func NewOutClusterClient() (*kubernetes.Clientset, error) {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset, err
}

func NewInClusterClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, err
}
