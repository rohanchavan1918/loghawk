package main

// func sendSlackAlert(msg string) {
// 	var AlertFunctionTitle, URL string
// 	if lambdacontext.FunctionName != "" {
// 		AlertFunctionTitle = "*[" + lambdacontext.FunctionName + "]*"
// 		URL = Env.SlackAPI
// 	} else {
// 		AlertFunctionTitle = fmt.Sprintf("*[SAS-MANAGER-%v]*", Config.IDC)
// 		URL = Config.SlackAPI
// 	}

// 	client := http.Client{
// 		Timeout: time.Duration(5 * time.Second),
// 	}

// 	jsonmap := Payload{
// 		Blocks: []Block{
// 			{
// 				Types: "section",
// 				Text: Text{
// 					Type: "mrkdwn",
// 					Text: AlertFunctionTitle,
// 				},
// 			},
// 			{
// 				Types: "section",
// 				Text: Text{
// 					Type: "mrkdwn",
// 					Text: msg,
// 				},
// 			},
// 		},
// 	}
// 	jsonBody, err := json.Marshal(jsonmap)
// 	if err != nil {
// 		DebugLog("", "error marshalling json map %+v", err)
// 		return
// 	}
// 	resp, err := client.Post(URL, "application/json", bytes.NewBuffer(jsonBody))
// 	if err != nil {
// 		DebugLog("", "error doing a post call %+v", err)
// 		return
// 	}

// 	if resp.StatusCode >= 400 {
// 		DebugLog("", "error getting status 200 %+v", resp.Status)
// 	}

// 	resp.Body.Close()
// }
