package models


// A B S T R  A C T I O N---hiding
type Person struct {  
   //this properties will not be accesible in the main file becaus they start with small letters  
	name string            
	age int  
	
}      
// Through Constructor we are able assign values to new struct(Person)--abstraction     

func (p*Person) Constructor(name string, age int)  {    
	p.name = name   
	p.age  = age 
	   
}          

// GETTERS AND SETTERS       
// GET
func (p*Person) GetName()  string{
	  return p.name   
}         

// SETTER   

func (p*Person) SetName(name string)  {   
         
	  p.name = name   
	
}

func (p*Person)  GetAge()  int{    
	
	return  p.age    

}    

func (p*Person) SetAge(age int)  {
	
	p.age = age
}


