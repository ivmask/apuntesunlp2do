package main

import("fmt")

func main() {
  var c string
  for (c != "Z") {
    fmt.Println("Ingresa punto cardinal de donde viene el viento, (N, S, E, O), o Z para terminar: ")
    fmt.Scan(&c)
    switch c {
      case "N": fmt.Println("El viento se dirige hacia el sur")
      case "S": fmt.Println("El viento se dirige hacia el norte")
      case "E": fmt.Println("El viento se dirige hacia el oeste")
      case "O": fmt.Println("El viento se dirige hacia el este")
      default: fmt.Println("Eso no es un punto cardinal")
    }
  }
}
