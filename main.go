package main

import (
	"log"

	"github.com/FUnigrad/funiverse-workspace-service/config"
	"github.com/FUnigrad/funiverse-workspace-service/handler"
)

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Cannot load config: ", err)
	}

	server := handler.NewServer(config)

	err = server.Start()
	if err != nil {
		log.Fatalln("Cannot start server:", err)
	}
}

// func main() {

// 	workspace := model.WorkspaceDTO{
// 		Name:         "abc",
// 		Code:         "fpt1",
// 		Domain:       "fpt1.funiverse.world",
// 		PersonalMail: "fpt_12@gmail.com",
// 		EduMail:      "fpt_123@edu.com",
// 	}

// 	request_body, _ := json.Marshal(workspace)
// 	req, _ := http.NewRequest("POST", "http://authen.system.funiverse.world/workspace", bytes.NewReader(request_body))

// 	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJyb2xlIjoiU1lTVEVNX0FETUlOIiwidXNlcm5hbWUiOiJmdW5pZ3JhZDIwMjNAZ21haWwuY29tIiwic3ViIjoiZnVuaWdyYWQyMDIzQGdtYWlsLmNvbSIsImlhdCI6MTY4MDMzODM2NSwiZXhwIjoxNjgwNTExMTY1fQ.ahJ7cM7yDD-VgC567rRqjUGRevGBVwrnieA8zJF1UYc")
// 	req.Header.Set("Content-Type", "application/json")
// 	log.Println(req.Header)
// 	client := http.Client{}

// 	resp, _ := client.Do(req)
// 	body, _ := io.ReadAll(resp.Body)
// 	fmt.Println(string(body))

// }
