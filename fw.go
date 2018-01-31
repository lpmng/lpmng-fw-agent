package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)


func echoHandler(w http.ResponseWriter, r *http.Request){
	var event map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil{
		panic(err)
	}
	fmt.Println(event)
}



func main(){
	http.HandleFunc("/echo", echoHandler)

	http.ListenAndServe("10.82.0.42:8000", nil)
}


