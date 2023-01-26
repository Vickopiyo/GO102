package figures

import "math"  

type Circle  struct {  

	 Radius  float64    
      
}      

func (cr*Circle) area()  float64{             
	  
	return  math.Pi*(cr.Radius * cr.Radius )     
	   
}  
           
func (cr*Circle) perimeter()  float64{      
          
	return 2 * math.Pi * cr.Radius 
	
}         