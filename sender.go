package main

import(
	"net/http"
	"encoding/json"
	"fmt"
)

/*
Send Json response to client
*/
func sendJson(ar interface{},w http.ResponseWriter)(err error){
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(ar);
	if err != nil{
		fmt.Println(err)
	}
	return
}

/*
send response to client
*/
func sendRes(aa int,res string,w http.ResponseWriter){
	w.WriteHeader(aa)
	w.Write([]byte(res));
}
