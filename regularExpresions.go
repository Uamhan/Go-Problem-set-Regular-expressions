package main

import (
	"fmt"
	"math/rand"
	"time"
)

func elizaResponses(s string) (string) {

	rand.Seed(time.Now().UnixNano())
	//input := s
	random := rand.Intn(3)
	var response string

	if random == 0{
		response = "I’m not sure what you’re trying to say. Could you explain it to me?"
	}else if random == 1{
		response = "How does that make you feel?"
	}else{
		response = "Why do you say that?"
	}

	return response

} 

func main() {
	
	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(elizaResponses("People say I look like both my mother and father."))

	fmt.Println("Father was a teacher.")
	fmt.Println(elizaResponses("Father was a teacher."))
	
	fmt.Println("I was my father’s favourite.")
	fmt.Println(elizaResponses("I was my father’s favourite."))
	
	fmt.Println("I’m looking forward to the weekend.")
	fmt.Println(elizaResponses("I’m looking forward to the weekend."))
	
	fmt.Println("My grandfather was French!")
	fmt.Println(elizaResponses("My grandfather was French!"))
	

}