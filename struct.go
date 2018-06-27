package main

//parameters specification for fizzbuzz suite generation  
type FizzBuzz struct{
	Res []string `json:"-"`
	Final string  `json:"-"`
	string1 string `json:"string1"`
	string2 string `json:"string2"`
	int1 int `json:"int1"`
	int2 int `json:"int2"`
	limit int `json:"limit"`
}
