package main

import (
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"log"
)

type Alert struct {
	ServiceKey  string      `json:"-"`
	ServiceId   int64       `json:"service_id"`
	Description string      `json:"description"`
	EventType   string      `json:"event_type"`
}

func (alert *Alert) Init(serviceKey string, serverId int64) {
	alert.EventType = "trigger"
	alert.ServiceKey = serviceKey
	alert.ServiceId = serverId
}

func (alert *Alert) Alert(description string) {
	if description == "" {
		description = "nil message"
	}
	alert.Description = description
	byteJson, _ := json.Marshal(alert)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", serverUrl, bytes.NewReader(byteJson))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("servicekey", string(alert.ServiceKey))

	res, _ := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err == nil {
		log.Println("Request Baidu :  " + string(body))
	} else {
		log.Println(err)
	}
}
