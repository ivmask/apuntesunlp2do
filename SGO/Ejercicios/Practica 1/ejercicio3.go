package main 

import("fmt")

func main () {
  var zz int = 5 // Mala declaracion 
  x := 10;
  var z int = x;
  const n = 5001 // Incorrecta declaracion de constante
  const c int = 5001 // Same
  var e float32 = 6
  var f float32 = e

  fmt.Println(zz, n, c, e, f);
}
