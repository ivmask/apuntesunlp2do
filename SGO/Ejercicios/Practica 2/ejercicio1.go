package main

import (
  "fmt";
  "strings";
  "unicode";
  "bufio";

)

func leerTemps(temps map[string][]float64) map[string][]float64 {
  acum := 0
  index1 := 0
  index2 := 0
  index3 := 0

  for i := range 10 {
    var float64 s = fmt.Scan(&s)
    if (s > 37.5) {
      temps[alta][index1] = s 
      ++index1
    } else if (s >= 36) {
      temps[normal][index2] = s 
      ++index2
    } else {
      temps[baja][index3] = s
      ++index3
    }
  }
}

func main() {
  var temps = make(map[string][]float64) {
    "alta": make([]float64, 10)
    "normal": make([]float64, 10)
    "baja": make([]float64, 10)
  
  reader := bufio.NewReader("\n")

}
  
}
