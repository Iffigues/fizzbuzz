package main

import(
	"os"
	"fmt"
	"strconv"
)

/*
Print help
*/
func help(a []string){
	println("--socket <int> : change port number\n")
	println("--string1 <string> : change string1\n")
	println("--string2 <string> : change string2\n")
	println("--int1 <int> : change int1\n")
	println("--int2 <int> : change int2\n")
	println("--limit <int> : change limit\n")
	println("--command: open bash command\n")
	os.Exit(0)
}


/*
change Socket of server
*/
func changeSocket(a []string){
	if len(a) > 1  && a[1] != "" {
		_,err := strconv.Atoi(a[1])
		if err != nil{
			fmt.Println("socket must be decimal number between 1 and 9999\n")
			return
		}
		Socket = ":"+a[1]
		fmt.Println("Server listening into ",a[1]," socket")
	}
}

/*
Change string1 of simple fizzbuzz
*/
func changeString1(a []string){
	if len(a) > 1{
		f.string1 = a[1]
		fmt.Println("String1 change to: ",a[1])
	}
}

/*
Change string2 of simple fizzbuzz
*/
func changeString2(a []string){
	if len(a) > 1{
		f.string2 = a[1]
		fmt.Println("String2 change to: ",a[1])
	}
}

/*
Change int1 of simple fizzbuzz
*/
func changeInt1(a []string){
	if len(a) > 1{
		i,err := strconv.Atoi(a[1])
		if err == nil && i != 0{
			f.int1 = i
			fmt.Println("int1 change to: ",a[1])
		}
	}
}


/*
Change int2 of simple fizzbuzz
*/
func changeInt2(a []string){
	if len(a) > 1{
		i,err := strconv.Atoi(a[1])
		if err == nil && i != 0{
			f.int1 = i
			fmt.Println("int1 change to: ",a[1])
		}
	}
}

/*
Change limit of simple fizzbuzz 
*/
func changeLimit(a []string){
	if len(a) > 1{
		i,err := strconv.Atoi(a[1])
		if err == nil && i > 0{
			f.limit = i
			fmt.Println("limit change to: ",a[1])
		}
	}
}
