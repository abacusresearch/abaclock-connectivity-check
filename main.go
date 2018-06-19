package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

func check(response http.ResponseWriter, request *http.Request) {
    query := request.URL.Query()

    resultMap := map[string]string {
        "deviceId": query["deviceId"][0],
        "deviceLabel": query["deviceLabel"][0],
    }

    resultString, err := json.Marshal(resultMap)

    if err != nil {
        panic(err)
    }

    fmt.Println(string(resultString))

    response.WriteHeader(204)
}

func main() {
    http.HandleFunc("/v1/check", check)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
