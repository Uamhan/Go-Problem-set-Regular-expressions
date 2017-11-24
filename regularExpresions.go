package main

import (
	"fmt"
	"math/rand"
	"time"
	"regexp"
	"strings"
)

func elizaResponses(s string) (string) {

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(3)
	var response string

	matched, _ := regexp.MatchString("(?i)^.*\\bfather\\b.*$", s)
	
	imMatched, _:=regexp.MatchString("(?i)^\\s*(?:I am|I'm) ([^.!?]*)[.!?\\s]*$", s)

	matched1, _ := regexp.MatchString("(?i)^.*psychiatrist.*$", s)

	matched2, _ := regexp.MatchString("^.*hair.*$", s)

	matched3, _ := regexp.MatchString("^.*holiday$", s)

	if matched == true{
		response = "Why don’t you tell me more about your father?"
	}else if imMatched ==true{
		
		response = "How do you know you are$1?"
		s = strings.Replace(s, "I am","",-1)
		s = strings.Replace(s, ".","",-1)
		s = strings.Replace(s, "?","",-1)
		s = strings.Replace(s, "you're", "I'm", -1)
		s = strings.Replace(s, "your", "my", -1)
		s = strings.Replace(s, "you", "i", 1)
		s = strings.Replace(s, "me", "you", -1)
		response = strings.Replace(response, "$1", s, -1)
	}else if matched1 == true {
		response = "if i was a psychiatrist would that make you feel difrent about me?"
	}else if matched2 == true {
		response = "I'm glad you do"
	}else if matched3 == true {
		response = "What was your favorite holiday?"
	}else if random == 0{
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
	fmt.Println("")
	fmt.Println("Father was a teacher.")
	fmt.Println(elizaResponses("Father was a teacher."))
	fmt.Println("")
	fmt.Println("I was my father’s favourite.")
	fmt.Println(elizaResponses("I was my father’s favourite."))
	fmt.Println("")
	fmt.Println("I’m looking forward to the weekend.")
	fmt.Println(elizaResponses("I’m looking forward to the weekend."))
	fmt.Println("")
	fmt.Println("My grandfather was French!")
	fmt.Println(elizaResponses("My grandfather was French!"))
	fmt.Println("")
	fmt.Println("I am happy.")
	fmt.Println(elizaResponses("I am happy."))
	fmt.Println("")
	fmt.Println("I'm not happy with your responses.")
	fmt.Println(elizaResponses("I am not happy with your responses."))
	fmt.Println("")
	fmt.Println("I'M not sure that you understand the effect that your questions are having on me.")
	fmt.Println(elizaResponses("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println("")
	fmt.Println("i am supposed to just take what you're saying at face value?")
	fmt.Println(elizaResponses("I am supposed to just take what you're saying at face value?"))
	fmt.Println("")
	fmt.Println("Are you a psychiatrist?")
	fmt.Println(elizaResponses("Are you a psychiatrist?"))
	fmt.Println("")
	fmt.Println("I like this chair")
	fmt.Println(elizaResponses("i like this chair"))
	fmt.Println("")
	fmt.Println("I think i need a holiday")
	fmt.Println(elizaResponses("I think i need a holiday"))
	fmt.Println("")
	fmt.Println("My favorite holiday was to Greece")
	fmt.Println(elizaResponses("My favorite holiday was to Greece"))
	

	
	
}