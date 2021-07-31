package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var metricsPort = "9909"
var metricsHost = "0.0.0.0"
var metricsEndpoint = "/metrics"
var kamStatusFile = "KaM_Remake_Server_Status.xml"

type kamMetrics struct {
	XMLName     xml.Name `xml:"server"`
	RoomCount   string   `xml:"roomcount"`
	ClientCount string   `xml:"clientcount"`
	PlayerCount string   `xml:"playercount"`
	//BytesSent     string   `xml:"bytessent"`
	//BytesReceived string   `xml:"bytesreceived"`
}

func main() {
	metricsPortEnv, present := os.LookupEnv("MetricsPort")
	if !present {
		fmt.Println("MetricsPort variable is not present. Using default: " + metricsPort)
	} else {
		metricsPort = metricsPortEnv
		fmt.Println("MetricsPort variable is present. Overriding with: " + metricsPort)
	}

	metricsHostEnv, present := os.LookupEnv("MetricsHost")
	if !present {
		fmt.Println("MetricsHost variable is not present. Using default: " + metricsHost)
	} else {
		metricsHost = metricsHostEnv
		fmt.Println("MetricsHost variable is present. Overriding with: " + metricsHost)
	}

	fmt.Println("Scrape from: http://" + metricsHost + ":" + metricsPort + metricsEndpoint)

	http.HandleFunc(metricsEndpoint, metrics)
	http.ListenAndServe(metricsHost+":"+metricsPort, nil)
}

func metrics(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "# kam-server-metrics-block #\n")

	xmlFile, err := os.Open(kamStatusFile)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var kamData kamMetrics
	xml.Unmarshal(byteValue, &kamData)

	fmt.Fprintf(w, "kam_roomcount "+kamData.RoomCount+"\n")
	fmt.Fprintf(w, "kam_clientcount "+kamData.ClientCount+"\n")
	fmt.Fprintf(w, "kam_playercount "+kamData.PlayerCount+"\n")
	//fmt.Fprintf(w, "kam_bytessent "+kamData.BytesSent+"\n")
	//fmt.Fprintf(w, "kam_bytesreceived "+kamData.BytesReceived+"\n\n")
}
