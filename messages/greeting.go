package messages
  
import "fmt"
// Exported name MUST be in UpperCase  
func Hello() {  

fmt.Println("Hello from main package !")   
	
}        
          
const salamu = "This is a constant!!"            
    //  Cannot be exported because func begins with small letters 
func functionPrivate()  {         
	
   fmt.Println("Hello from Private Function")     	   

}   
      
  func Output()  {       

	   fmt.Println(salamu)
	   functionPrivate()     
  }    

      

    

            
               
          