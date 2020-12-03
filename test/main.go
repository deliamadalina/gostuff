package main

import (
	"encoding/json"
	"fmt"
)

type Worker struct {
	Viewer                     map[string]interface{} `json:"viewer"`
	Accounts                   map[string]interface{} `json:"accounts"`
	WorkeersInvocationAdaptive map[string]interface{} `json:"workeersInvocationAdaptive"`
	Sum                        []string               `json:"sum"`
	Requests                   float64                `json:"requests"`
	Subrequests                float64                `json:"subrequests"`
	Errors                     float64                `json:"errors"`
	Dimensions                 []string               `json:"dimensions"`
	ScriptName                 string                 `json:"scriptname"`
}

var WorkerAccounts struct {
	Accounts map[string]interface{} `json:"accounts"`
}
var workerdata string = `
{"Viewer":
  {"Accounts":
    [{"WorkeersInvocationAdaptive":
      [{
        "Sum":
          {"Requests":14780,"Subrequests":0,"Errors":0},
        "Dimensions":{"ScriptName":"iedc_worker_prod"}
        },
       {
        "Sum":
          {"Requests":6080969,"Subrequests":6080495,"Errors":0},
        "Dimensions":{"ScriptName":"mobile-prod"}},
       {
        "Sum":{"Requests":256,"Subrequests":258,"Errors":0},
        "Dimensions":{"ScriptName":"root_commerce_worker_int"}},
       {
        "Sum":{"Requests":153,"Subrequests":153,"Errors":0},
        "Dimensions":{"ScriptName":"root_commerce_worker_stage"}},
       {
        "Sum":{"Requests":129703,"Subrequests":129543,"Errors":4},
        "Dimensions":{"ScriptName":"root_commerce_worker_prod"}},
       {
        "Sum":{"Requests":2,"Subrequests":2,"Errors":0},
        "Dimensions":{"ScriptName":"desktop-static"}}]}]}}
`

func main() {
	var Worker Worker
	err := json.Unmarshal([]byte(workerdata), &Worker)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nViewer: ", Worker.Viewer)

	fmt.Println(Worker.Viewer)

}
