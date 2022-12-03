package incident

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

func GetIncidentData() []IncidentData {
	var result []IncidentData
	resp, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil
	}
	return result
}

func GetResultSIncident() []IncidentData {
	var result []IncidentData
	resSupport := GetIncidentData()
	var activeStatus, closedStatus []IncidentData
	for _, val := range resSupport {
		if val.Status == "active" {
			activeStatus = append(activeStatus, val)
		} else {
			closedStatus = append(closedStatus, val)
		}
	}
	result = append(result, activeStatus...)
	result = append(result, closedStatus...)
	return result
}
