package main

import (
	// "fmt"
	// "packages/figures"
	// "packages/models"

	"github.com/donvito/hellomod"

	"github.com/Vickopiyo/figures"
)

// import "packages/messages"
    
func main() {                   
	// messages.Hello()          
	// messages.Output()      

	//  circle1 := figures.Circle{Radius: 23 }          
    //  figures.Measurements(&circle1)      
	 
	//  square1 := figures.Square{Side: 21}  
	//   figures.Measurements(&square1)             


	// CANNOT assign fields because they inaccesble(abstraction)
    //  vick := models.Person{ }           
	//  Through  Constructor func , allowed fields can be assigned!    
    //   vick.Constructor("Vick", 24)  

      // Impossible bcz type Person has no fields 
	//   vick.name  = "VIKKKK"

	//    fmt.Println(vick)   
    //    fmt.Println(vick.GetName())     
	//    fmt.Println(vick.GetAge())               
	//    vick.SetName("VICK OPIYO")
	//    fmt.Println(vick)              
    
// BY IMPORTING github/donvito/hellomod---we are able to acces its package(hellomod)and its related functions (sayHello)   
                                  
        hellomod.SayHello()        

// IMPORT FROM vickopiyo/figures  ----figures is the  package  name (repo name)    

     circle1 := figures.Circle{
		Radius: 23,
	 }   

	 figures.Measurements(&circle1)        
                


		

              
}                             