package goclient

import (
	"log"
	"path/filepath"

	"github.com/FUnigrad/funiverse-workspace-service/config"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type GoClient struct {
	Client *dynamic.DynamicClient
}

func GetInClusterConfig() (config *rest.Config) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	return
}
func GetOutClusterConfig() (config *rest.Config) {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return
}

func NewClient(config config.Config) (*GoClient, error) {
	var configK8s *rest.Config
	if config.Enviroment == "local" {
		configK8s = GetOutClusterConfig()
	} else if config.Enviroment == "prod" {
		configK8s = GetInClusterConfig()
	} else {
		log.Fatalln("Enviroment configured not correct!")
	}

	client, err := dynamic.NewForConfig(configK8s)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return &GoClient{Client: client}, err
}
