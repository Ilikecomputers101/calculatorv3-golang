package main
import (
	"userinput"
	"fmt"
        "log"
)

func main() {
	fmt.Println("1=+, 2=-,3=/")
	fmt.Print("enter a number: ")
	number, err := userinput.Userput()
	fmt.Print("enter second number: ")
	number2, err := userinput.Userput()
	fmt.Print("choose operator: ")
	perator, err := userinput.Userput()
	if err != nil {
		log.Fatal(err)
	}
	if perator <= 1 {
		awnser := number + number2
	        fmt.Println(awnser)
        }
	if perator >= 2 {
		awnser := number - number2
		fmt.Println(awnser)
	}
        if perator >= 3 {
		awnser := number / number2
		fmt.Println(awnser)
	}
	

	
}
