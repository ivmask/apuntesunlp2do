// Array = Secuencia indexada de un mismo tipo, con primer indice en cero

//var z [n]tipo;
var x [5]int
x[4] = 100

z := [5]float64{
  12,
  23,
  34,
  45,
  56
} // Inicializacion posicional

arr := [5]int{1:10, 2:20, 3:45} // Inicializacion nombrada

// FOR RANGE = Iteracion por cada elemento de un Array

var total float64 = 0
for i, value := range x {
  total += value
  fmt.Println(i, value)
}

// Arrays bidimensionales

var a [2][2]string

a[0][0] = ...
a[0][1]
a[1][0]
a[1][1]

// SLICE = Porcion de un Array
// Posee las operaciones .len() y .cap() que devuelven largo y capacidad maxima determinada por el arreglo subyacente
a := [6] int {1,2,3,4,5,6}
s := a[2:4]//
