package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Payload struct {
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Types string `json:"type"`
	Text  Text   `json:"text,omitempty"`
}

type Text struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Verbatim bool   `json:"verbatim"`
}

func SendSlackAlert(msg string, functionName string, alertUrl string) {
	var AlertFunctionTitle string
	if functionName != "" {
		AlertFunctionTitle = "*[" + functionName + "]*"
	}

	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	jsonmap := Payload{
		Blocks: []Block{
			{
				Types: "section",
				Text: Text{
					Type: "mrkdwn",
					Text: AlertFunctionTitle,
				},
			},
			{
				Types: "section",
				Text: Text{
					Type: "mrkdwn",
					Text: msg,
				},
			},
		},
	}
	jsonBody, err := json.Marshal(jsonmap)
	if err != nil {
		fmt.Println("", "error marshalling json map :", err)
		return
	}
	resp, err := client.Post(alertUrl, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("", "error doing a post call :", err)
		return
	}

	if resp.StatusCode >= 400 {
		fmt.Println("", "error getting status 200 :", resp.Status)
	}

	resp.Body.Close()
}
