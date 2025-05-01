package main

import("fmt")

func main() {
  var temps [10]float64
  var arrP [3]int
  var tot float64
  //var pmym float64
  var max, min float64 = 1, 9999

  for i, s := range temps {
    fmt.Scan(&s)
    temps[i] = s
    tot += s
    
    if (s > 37.5) {
      arrP[0]++
    } else if (s >= 36) {
      arrP[1]++
    } else {
      arrP[2]++
    }
    
    if (s > max) {
      max = s
    } else if (s < min) {
      min = s
    }
    
  }

  for i := range arrP {
    fmt.Println("El promedio ", i+1, " es: ")
    j := float64((arrP[i]/10.0)*100)
    fmt.Println("%v%%\n", j) 
  }

}
