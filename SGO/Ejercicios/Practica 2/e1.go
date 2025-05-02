package main

import (
	"fmt";
)


func analizarTemperaturasA(n int) {
	var j float64 = float64(n)
	a := []float64{0,0,0}
	c := []float64{0,0,0}
	var k float64
	for range n {	
		fmt.Scan(&k)
		switch {
		case k > 37.5: 
			{
				a[2] += k 
				c[2] += 1
			}
		case k >= 36.0: 
			{
				a[1] += k
				c[1] += 1
			}
		default:
			{
				a[0] += k	
				c[0] += 1
			}
		}
	}
	for i := range len(a) {
		w := ((c[i]/j)*100) 
		fmt.Printf("El porcentaje de pacientes del grupo %d es %.2f%% \n", i, w)
	}
}

func main() {
	analizarTemperaturasA(10)
	//analizarTemperaturasB(10)
}
