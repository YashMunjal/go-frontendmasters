package main

import "fmt"

func main(){
	mySentence := "hello this is a sentence"
	for index,value :=range mySentence{
		if index%2 ==0{
			fmt.Println(string(value))
		}
	}
}