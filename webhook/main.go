package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/alertmanager/template"
)

type Message struct {
	Log        string     `json:"log"`
	Time       time.Time  `json:"time"`
	Kubernetes Kubernetes `json:"kubernetes"`
}

type Kubernetes struct {
	PodName        string `json:"pod_name"`
	NamespaceName  string `json:"namespace_name"`
	ContainerName  string `json:"container_name"`
	DockerID       string `json:"docker_id"`
	ContainerImage string `json:"container_image"`
}

//type sendMsg struct {
//	Receiver string   `json:"receiver"`
//	Status   string   `json:"status"`
//	Alert    Alerts   `json:"alerts"`
//}
//type Alerts struct {
//	Status       string      `json:"status"`
//	Labels       template.KV `json:"labels"`
//	Annotations  template.KV `json:"annotations"`
//	StartsAt     time.Time   `json:"startsAt"`
//	EndsAt       time.Time   `json:"endsAt"`
//	GeneratorURL string      `json:"generatorURL"`
//	Fingerprint  string      `json:"fingerprint"`
//}

//type Labels struct {
//	Alertname  string `json:"alertname"`
//	Container  string `json:"container"`
//	Namespace  string `json:"namespace"`
//	Pod        string `json:"pod"`
//	Prometheus string `json:"prometheus"`
//	Severity   string `json:" severity"`
//}

//type Annotations struct {
//	Message     string `json:"message"`
//	Runbook_url string `json:"runbook_url"`
//}

// Sample HTTP receiver for this demo
func main() {

	h := func(w http.ResponseWriter, req *http.Request) {
		b, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			log.Printf(err.Error())
			return
		}

		var msg []Message
		//var sendmsg []sendMsg
		//_ = sendmsg
		err = json.Unmarshal(b, &msg)
		if err != nil {
			log.Printf(err.Error())
			return
		}

		for _, item := range msg {
			var send template.Data
			send.Receiver = "Default"
			send.Status = "firing"
			var alert template.Alert
			alert.Status = "firing"
			alert.Labels = template.KV{
				"alertname":      "logging-alert",
				"container":      item.Kubernetes.ContainerName,
				"containerimgae": item.Kubernetes.ContainerImage,
				"namespace":      item.Kubernetes.NamespaceName,
				"pod":            item.Kubernetes.PodName,
				"label":          "keyword",
				"severity":       "normal",
			}
			alert.Annotations = template.KV{
				"message": item.Log,
			}

			alert.StartsAt = item.Time
			alert.Fingerprint = "83fb3d34d52108b0"
			alert.EndsAt, _ = time.Parse("2006-01-02", "0001-01-01T00:00:00Z")
			send.Alerts = append(send.Alerts, alert)
			sendmsg(send)
			log.Printf("log=%s, kubernetes=%s, time=%s\n", item.Log, item.Kubernetes, item.Time)
			fmt.Printf("log=%s, kubernetes=%s, time=%s\n", item.Log, item.Kubernetes, item.Time)
		}
	}
	http.HandleFunc("/", h)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
func sendmsg(msg template.Data) {

	//sendjson, errs := json.Marshal(msg) //转换成JSON返回的是byte[]
	//if errs != nil {
	//	fmt.Println(errs.Error())
	//}
	//sendstr :=string(sendjson)
	//url:= "http://10.233.50.126:19093/api/v2/alerts"
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(msg)
	fmt.Printf("msg-2.0-------=%s, resquestbody----=%s", msg, requestBody)
	url := "http://notification-manager-svc.kubesphere-monitoring-system.svc:19093/api/v2/alerts"
	req, err := http.NewRequest("POST", url, requestBody)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)

	fmt.Println("response Headers:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))

}
