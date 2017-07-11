package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"log"
	"os"
	"strconv"
)

type EvalMatches struct {
	Metric string       `json:"metric"`
	Tags   interface{}  `json:"tags"`
	Value  float64        `json:"value"`
}

type GrafanaCallBack struct {
	Title       string `json:title`
	RuleId      int32  `json:ruleId`
	RuleName    string `json:ruleName`
	RuleUrl     string `json:ruleUrl`
	State       string `json:"state"`
	ImageUrl    string `json:"imageUrl"`
	Message     string `json:"message"`
	EvalMatches []EvalMatches `json:"evalMatches"`
}

var serviceKey string
var serverId int64
var serverUrl string
var advertisedUrl string

func init() {
	serviceKey = os.Getenv("SERVICEKEY")
	serverId, _ = strconv.ParseInt(os.Getenv("SERVERID"), 10, 64)
	serverUrl = os.Getenv("SERVERURL")
	advertisedUrl = os.Getenv("ADVERTISEDURL")

	log.Println("serviceKey: " + serviceKey)
	log.Println("serverId:  " + strconv.FormatInt(serverId, 10))
	log.Println("serverUrl:  " + serverUrl)
	log.Println("advertisedUrl:  " + advertisedUrl)
}

func main() {

	http.HandleFunc("/", payload)
	if err := http.ListenAndServe(advertisedUrl, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func payload(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	body_str := string(body)
	log.Println(body_str)

	var grafana GrafanaCallBack
	err := json.Unmarshal([]byte(body_str), &grafana)
	if err == nil {
		if strings.EqualFold("alerting", grafana.State) {
			alert := new(Alert)
			alert.Init(serviceKey, serverId)
			alert.Alert(grafana.Message)
		}
	} else {
		log.Println(err)
	}
}
