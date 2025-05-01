package main

import("fmt")

func division(int k1, int k2) int {
  if (k1 > k2 ) {
    return k1/k2;
  } else {
    return k2/k1;
  }
}

func main() {
  var int n1
  var int n2
  n1 = fmt.Scan(&n1)
  n2 = fmt.Scan(&n2)
  
  division()
}
