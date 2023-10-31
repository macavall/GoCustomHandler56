package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ReturnValue struct {
	Data string
}
type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

type InvokeResponseStringReturnValue struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue string
}

type InvokeRequest struct {
	Data     map[string]interface{}
	Metadata map[string]interface{}
}

func httpTriggerHandlerStringReturnValue(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Year())
	ua := r.Header.Get("User-Agent")
	fmt.Printf("user agent is: %s \n", ua)
	invocationid := r.Header.Get("X-Azure-Functions-InvocationId")
	fmt.Printf("invocationid is: %s \n", invocationid)

	outputs := make(map[string]interface{})
	outputs["output"] = "Mark Taylor"
	outputs["output2"] = map[string]interface{}{
		"home":   "123-466-799",
		"office": "564-987-654",
	}
	headers := make(map[string]interface{})
	headers["header1"] = "header1Val"
	headers["header2"] = "header2Val"

	res := make(map[string]interface{})
	res["statusCode"] = "201"
	res["body"] = "my world"
	res["headers"] = headers
	outputs["res"] = res
	invokeResponse := InvokeResponseStringReturnValue{outputs, []string{"test log1", "test log2"}, "Hello,World"}

	js, err := json.Marshal(invokeResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if exists {
		fmt.Println("FUNCTIONS_CUSTOMHANDLER_PORT: " + customHandlerPort)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/HttpTriggerStringReturnValue", httpTriggerHandlerStringReturnValue)
	fmt.Println("Go server Listening...on FUNCTIONS_CUSTOMHANDLER_PORT:", customHandlerPort)
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
}
