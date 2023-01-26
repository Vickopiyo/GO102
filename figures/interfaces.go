package figures

import "fmt"  

// Interface calls  the two shared properties(functions) of struct 
type Geometry interface {     
	area() float64
	perimeter() float64
}
  // Takes the  interface  and accesses each func in the interface   

func Measurements(gr Geometry) {  

    fmt.Println(gr)           
	fmt.Println(gr.area())   
	fmt.Println(gr.perimeter())    
        
                  
}  


           