package main

import(
	"strings"
	"bufio"
	"fmt"
	"os"
)

/*
Start bash command before start server
*/
func command(a []string){
	var (
		reader = bufio.NewReader(os.Stdin)
		vrai = true
	)	
	commands := map[string]func([]string){
		"1\n" : changeString1,
		"2\n" : changeString2,
		"3\n" : changeInt1,
		"4\n" : changeInt2,
		"5\n" : changeLimit,
		"6\n" : changeSocket,
	}
	for vrai {
		fmt.Print("Choose choice:\n")
		fmt.Print("1) string1\n")
		fmt.Print("2) string2\n")
		fmt.Print("3) int1\n")
		fmt.Print("4) int2\n")
		fmt.Print("5) limit1\n")
		fmt.Print("6) Socket\n")
		fmt.Print("quit) Exit command\n")
		text, _ := reader.ReadString('\n')		
		if text == "quit\n" {
			vrai = false
		}else{
			if com,ok := commands[text];ok{
				changeCommand(text,com)
			} 
		}
	}
	fmt.Println("exit command\n")
}


func changeCommand(text string, fun func([]string)){
	var (
		reader = bufio.NewReader(os.Stdin)
		array []string
		com = strings.Trim(text,"\n")
	)
	array = append(array,com)
	fmt.Print("Enter text: ")
	texte, _ := reader.ReadString('\n')
	array = append(array,strings.Trim(texte,"\n"))
	fun(array)	
}
