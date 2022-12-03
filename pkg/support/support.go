package support

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func GetSupportData() []SupportData {
	var result []SupportData
	resp, err := http.Get("http://127.0.0.1:8383/support")
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

func GetResultSupport() []int {
	var result []int
	var sumTickets, waitTime int
	resSupport := GetSupportData()
	for _, val := range resSupport {
		sumTickets += val.ActiveTickets
	}
	if sumTickets < 9 {
		result = append(result, 1)
	} else if sumTickets >= 9 && sumTickets <= 16 {
		result = append(result, 2)
	} else {
		result = append(result, 3)
	}
	waitTime = sumTickets * (60 / 18)
	result = append(result, waitTime)
	return result
}
