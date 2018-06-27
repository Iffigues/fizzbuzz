package main

import(
	"net/http"
	"strings"
	"strconv"
)

type Res struct{
	A string
}

/*
send the simple fizzbuzz 
*/
func first(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET"{
		sendRes(202,Example,w)
	}
}

/*
get fizzbuzz url parameters
*/
func getPar(a string)(tt *FizzBuzz,err error){
	var t FizzBuzz
	var i int
	par := strings.Split(a,"/")[1:]
	t.string1 = par[0]
	t.string2 = par[1]
	i, err = strconv.Atoi(par[2])
	if err != nil{
		return nil,err
	}
	t.int1 = i
	i, err = strconv.Atoi(par[3])
	if err != nil {
		return nil,err
	}
	t.int2 = i
	i, err = strconv.Atoi(par[4])
	if err != nil{
		return nil,err
	}
	t.limit = i
	return &t,nil
}

/*
Send FizzBuzz with specific parameters
*/
func two(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET"{
		f, err := getPar(r.URL.Path)
		if err != nil {
			sendRes(405,err.Error(),w)
			return
		}
		err = f.Error()
		if err != nil {
			sendRes(405,err.Error(),w)
			return
		}
		mem := f.getCache()
		if memoire, ok := caches[mem];ok{
			sendJson(memoire.Res,w)
			return
		}
		f.MakeArray()
		caches[mem] = f
		sendJson(f.Res,w)
		return
	}
	sendRes(404,"Method Not allowed",w)
}
