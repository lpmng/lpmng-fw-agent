package main

import(
    "fmt"
    "net/http"
    "encoding/json"
)

type Session struct {
    Mac      string `json:"max"`
    Ip4      string `json:"ip4"`
    Internet bool   `json:"internet"`
    User     int    `json:"user"`
}

type EventSession struct {
    Action  string  `json:"action"`
    Name    string  `json:"name"`
    Session Session `json:"param"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Invalid request method.", 405)
        return
    }

    var event EventSession
    err := json.NewDecoder(r.Body).Decode(&event)

    if err != nil {
        http.Error(w, err.Error(), 400)
        return
    }

    switch event.Action {
        case "created", "updated" :
            fmt.Println(event.Session.Ip4)
            fmt.Println(event.Session.Internet)
            fmt.Println(event.Action)
        default:
            http.Error(w, "Wrong action type. Supported action are 'created', 'updated'.", 400)
    }
}



func main(){
    http.HandleFunc("/event/session", echoHandler)

    http.ListenAndServe("0.0.0.0:8000", nil)
}


