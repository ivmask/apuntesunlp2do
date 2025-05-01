package main

import("fmt");

func main() {
 
  const tope = 250;
  const tope2 = 2;
  var s int = 0;

  for i := 2; i <= tope; {
    s = s + i;
    i = i + 2;
  }
  
  fmt.Println(s)

  s=0;
  for i := 250; i >= tope2; {
    s = s + i;
    i = i-2;
  }

  fmt.Println(s)
}
