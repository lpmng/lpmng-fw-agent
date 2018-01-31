package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)


func echoHandler(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
          http.Error(w, "Invalid request method.", 405)
          return
        }

        var event map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&event)

        if err != nil {
		http.Error(w, "Invalid json", 400)
                return
	}

	fmt.Println(event)
}



func main(){
	http.HandleFunc("/event/session", echoHandler)

	http.ListenAndServe("0.0.0.0:8000", nil)
}


