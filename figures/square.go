package figures   

type Square  struct {   

	Side float64    
}       

func (sq*Square) area () float64 {   

	return sq.Side* sq.Side  
	     
}   

func (sq*Square) perimeter()  float64{  

	return sq.Side * 4    

}     

