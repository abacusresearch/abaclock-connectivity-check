package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

func check(response http.ResponseWriter, request *http.Request) {
    query := request.URL.Query()

    attributes := []string {
        "accountId",
        "deviceId",
        "deviceLabel",
    }

    resultMap := map[string]string {}

    for _, attribute := range attributes {
        if query[attribute] != nil {
            resultMap[attribute] = query[attribute][0]
        }
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
