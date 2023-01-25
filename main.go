package main

import (
	// "fmt"
	"packages/figures"
)

// import "packages/messages"
    
func main() {                   
	// messages.Hello()          
	// messages.Output()      

	 circle1 := figures.Circle{Radius: 23 }          
     figures.Measurements(&circle1)     
	

	 square1 := figures.Square{Side: 21}  
	  figures.Measurements(&square1)    


}                             