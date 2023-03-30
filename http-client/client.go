package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/FUnigrad/funiverse-workspace-service/config"
	"github.com/FUnigrad/funiverse-workspace-service/model"
)

type HttpClient struct {
	Hostname string
	Client   *http.Client
}

func NewClient(config config.Config) (*HttpClient, error) {
	var hostname string
	if config.Enviroment == "local" {
		hostname = "authen.system.funiverse.world"
	} else if config.Enviroment == "prod" {
		hostname = "authen-service:8000"
	} else {
		return nil, errors.New("configuration incorrect at env")
	}

	httpClient := HttpClient{
		Hostname: hostname,
		Client:   &http.Client{},
	}
	return &httpClient, nil
}

func (client *HttpClient) GetAllWorkspace(token string) (workspaces []model.Workspace) {

	url := fmt.Sprintf("http://%s/workspace", client.Hostname)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", token)

	resp, err := client.Client.Do(req)

	if err != nil {
		log.Print(err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Print(err.Error())
		return nil
	}

	if err = json.Unmarshal(body, &workspaces); err != nil {
		log.Print(err.Error())
		return nil
	}

	return
}

func (client *HttpClient) GetWorkspaceById(id int, token string) (workspace *model.Workspace) {
	url := fmt.Sprintf("http://%s/workspace/%d", client.Hostname, id)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", token)

	resp, err := client.Client.Do(req)

	if err != nil {
		log.Print(err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Print(err.Error())
		return nil
	}

	if err = json.Unmarshal(body, &workspace); err != nil {
		log.Print(err.Error())
		return nil
	}

	return
}
