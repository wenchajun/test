package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Log        string     `json:"log"`
	Time       string     `json:"time"`
	Kubernetes Kubernetes `json:"kubernetes"`
}

type Kubernetes struct {
	PodName        string `json:"pod_name"`
	NamespaceName  string `json:"namespace_name"`
	ContainerName  string `json:"container_name"`
	DockerID       string `json:"docker_id"`
	ContainerImage string `json:"container_image"`
}

type sendMsg struct {
	Receiver string `json:"receiver"`
	Status   string `json:"status"`
	Alert    []Alerts  `json:"alert"`
}
type Alerts struct {
	Status       string      `json:"status"`
	Labels       Labels      `json:"labels"`
	Annotations  Annotations `json:"annotations"`
	StartsAt     string      `json:"startsAt"`
	EndsAt       string      `json:"endsAt"`
	GeneratorURL string      `json:"generatorURL"`
	Fingerprint  string      `json:"fingerprint"`
}
type Labels struct {
	Alertname  string `json:"alertname"`
	Container  string `json:"container"`
	Namespace  string `json:"namespace"`
	Pod        string `json:"pod"`
	Prometheus string `json:"prometheus"`
	Severity   string `json:" severity"`
}

type Annotations struct {
	Message     string `json:"message"`
	Runbook_url string `json:"runbook_url"`
}

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
			var send sendMsg
			send.Receiver = "Default"
			send.Status = "firing"
            var alert Alerts
			alert.Status = "firing"
            alert.Labels.Alertname="Podinfo"
            alert.Labels.Container=item.Kubernetes.ContainerName
			alert.Labels.Namespace = item.Kubernetes.NamespaceName
			alert.Labels.Pod = item.Kubernetes.PodName
			alert.Labels.Prometheus = "kubesphere-monitoring-system/k8s"
			alert.Labels.Severity = "warning"
			alert.Annotations.Message = item.Log
			alert.Annotations.Runbook_url = "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-cputhrottlinghigh"
			alert.StartsAt = item.Time
			alert.GeneratorURL = "http://prometheus-k8s-0:9090/graph?g0.expr=sum+by%28container%2C+pod%2C+namespace%29+%28increase%28container_cpu_cfs_throttled_periods_total%7Bcontainer%21%3D%22%22%7D%5B5m%5D%29%29+%2F+sum+by%28container%2C+pod%2C+namespace%29+%28increase%28container_cpu_cfs_periods_total%5B5m%5D%29%29+%3E+%2825+%2F+100%29\u0026g0.tab=1"
			alert.Fingerprint = "83fb3d34d52108b0"
			alert.EndsAt="0001-01-01T00:00:00Z"
			send.Alert=append(send.Alert, alert)
            sendmsg(send)
			log.Printf("log=%s, stream=%s, time=%s\n", item.Log, item.Kubernetes, item.Time)
			fmt.Printf("log=%s, stream=%s, time=%s\n", item.Log, item.Kubernetes, item.Time)
		}
	}
	http.HandleFunc("/", h)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
func sendmsg(msg sendMsg)  {

	//sendjson, errs := json.Marshal(msg) //转换成JSON返回的是byte[]
	//if errs != nil {
	//	fmt.Println(errs.Error())
	//}
	//sendstr :=string(sendjson)
	//url:= "http://10.233.50.126:19093/api/v2/alerts"
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(msg)
	fmt.Printf("msg--------=%s, resquestbody----=%s", msg,requestBody )
	url:= "http://10.233.50.126:19093/api/v2/alerts"
	req , err := http.NewRequest("POST",url,requestBody)

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