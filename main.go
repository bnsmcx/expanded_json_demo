package main

import (
	"encoding/json"
	"html/template"
	"log"
	"os"
	"time"
)

func main() {
	rawData, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatalln(err)
	}

	var cse CSE
	err = json.Unmarshal(rawData, &cse)
	if err != nil {
		log.Fatalln(err)
	}

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("demo.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, cse)
	if err != nil {
		log.Fatalln(err)
	}
}

type CSE struct {
	AccountId          string    `json:"accountId"`
	EventFormatVersion string    `json:"eventFormatVersion"`
	EventId            string    `json:"eventId"`
	EventName          string    `json:"eventName"`
	EventTime          time.Time `json:"eventTime"`
	Insight            struct {
		Confidence       float64   `json:"confidence"`
		Created          time.Time `json:"created"`
		Description      string    `json:"description"`
		EntityId         string    `json:"entityId"`
		EntitySensorZone string    `json:"entitySensorZone"`
		EntityType       string    `json:"entityType"`
		EntityValue      string    `json:"entityValue"`
		Id               string    `json:"id"`
		Name             string    `json:"name"`
		ReadableId       string    `json:"readableId"`
		RiskScore        int       `json:"riskScore"`
		Severity         string    `json:"severity"`
		SeverityName     string    `json:"severityName"`
		Signals          []struct {
			ContentType       string        `json:"contentType"`
			Created           time.Time     `json:"created"`
			Description       string        `json:"description"`
			EntityId          string        `json:"entityId"`
			Id                string        `json:"id"`
			Name              string        `json:"name"`
			Prototype         bool          `json:"prototype"`
			RuleId            string        `json:"ruleId"`
			RuleName          string        `json:"ruleName"`
			Severity          string        `json:"severity"`
			Summary           string        `json:"summary"`
			SuppressedReasons []interface{} `json:"suppressedReasons"`
			Tags              []string      `json:"tags"`
			Timestamp         time.Time     `json:"timestamp"`
		} `json:"signals"`
		Source          string    `json:"source"`
		Status          string    `json:"status"`
		Tags            []string  `json:"tags"`
		TimeToDetection float64   `json:"timeToDetection"`
		Timestamp       time.Time `json:"timestamp"`
	} `json:"insight"`
	InsightIdentity struct {
		Id         string `json:"id"`
		ReadableId string `json:"readableId"`
	} `json:"insightIdentity"`
	Subsystem string `json:"subsystem"`
}
