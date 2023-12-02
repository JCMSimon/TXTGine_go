package main

import (
	"fmt"
)

func main() { // has to be called main
	var name = "Simon"
	const amIStupid bool = true // supplying a type is optional if a value is given

	var stupidTime uint8 // uint meaning its only positive and 0-255
	var todaysDay string

	fmt.Println("Hello", name, "!")
	fmt.Println("Are you stupid?", amIStupid, "!")
	fmt.Printf("Your name is %s and the hypothesis that you are stupid is: %t\n", name, amIStupid)

	//  get input here (any string)
	todaysDay = "monday"

	// get time input here (1-12)
	stupidTime = 8

	switch todaysDay {
	case "monday":
		fmt.Printf("It seems like today is %s (%d am/pm)... ew\n", todaysDay, stupidTime)
	case "tuesday":
		fmt.Printf("It seems like today is %s (%d am/pm)... its getting better\n", todaysDay, stupidTime)
	}

	stupidTime -= 50
	fmt.Printf("Now its (%d am/pm)... ew", stupidTime) // overflowing here. no error.. interesting!
}

// cont here https://youtu.be/yyUHQIec83I?t=2722
