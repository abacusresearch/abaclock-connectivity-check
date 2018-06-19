package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

func check(response http.ResponseWriter, request *http.Request) {
    for _, device := range request.URL.Query()["device"] {
        attributes, err := json.Marshal(map[string]string {"device": device})

        if err != nil {
            panic(err)
        }

        fmt.Println(string(attributes))
    }

    response.WriteHeader(204)
}

func main() {
    http.HandleFunc("/v1/check", check)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
