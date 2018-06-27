package main

import(
	"net/http"
	"strconv"
	"time"
	"log"
	"os"
	"fmt"
)

var(
	f = new(FizzBuzz)
	h = http.NewServeMux()
	test = NewRoute()
	Example string
	ReadTimeOut,_ = strconv.Atoi(getKey(htp,"ReadTime"))
	WriteTimeOut ,_= strconv.Atoi(getKey(htp,"WriteTime"))
	Socket = getKey(htp,"Socket")
)

func init(){
	test.HandleFunc(1,first)
	test.HandleFunc(5,two)
}


/*
First fizzbuzz specification
*/
func init(){
	f.string1 = "fizz"
	f.string2 = "buzz"
	f.int1 = 3
	f.int2 = 5
	f.limit = 100
	
}

/*
server side configuration and start
*/
func serve(){	
	s := &http.Server{
		Addr:           Socket,
		Handler:    test,
		ReadTimeout:    time.Duration(ReadTimeOut) * time.Second,
		WriteTimeout:   time.Duration(WriteTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("server start to port: "+Socket+"\n")	
	log.Fatal(s.ListenAndServe())
}


/*
 command specifications before starting server
*/
func main(){
	arg := os.Args[1:]
	commands := map[string]func([]string){
		"--command": command,
		"--help": help,
		"--server": changeSocket,
		"--string1": changeString1,
		"--string2": changeString2,
		"--int1": changeInt1,
		"--int2": changeInt2,
		"--limit": changeLimit,
	}
	for i,ok := range arg{
		if lance,ok := commands[ok];ok{
			lance(arg[i:])
		}
	}
	err := f.Error()
	if err != nil{
		log.Fatal(err)
	}
	f.MakeArray()
	Example,err = f.String()
	if err != nil{
		log.Fatal(err)
	}
	serve()
}
