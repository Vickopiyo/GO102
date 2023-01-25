package figures

import "fmt"

type Geometry interface {
	area() float64
	perimeter() float64
}

func Measurements(gr Geometry) {  

    fmt.Println(gr)    
	fmt.Println(gr.area())   
	fmt.Println(gr.perimeter())          

}
