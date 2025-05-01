package main

import (
  "fmt"
  "bufio"
  "strings"
_  "unicode"
  "os"
)

func invertirImpares (s string) string {
  auxSlice := strings.Fields(s)
  
  for i := range len(auxSlice) {
    if (i == 0) || ((i % 2) == 0) {
      var auxRuneSliceWord = []rune(auxSlice[i])
      for j, k := 0, len(auxRuneSliceWord)-1; j<k; j, k = j+1, k-1 {
        auxRuneSliceWord[j], auxRuneSliceWord[k] = auxRuneSliceWord[k], auxRuneSliceWord[j]
        var auxWord = string(auxRuneSliceWord)
        auxSlice[i] = auxWord
      } 
    }
  }

  retorno := strings.Join(auxSlice, " ") 
  return retorno
}

func main () {
  reader := bufio.NewReader(os.Stdin)
  s, _ := reader.ReadString('\n')

  invertido := invertirImpares(s)
  fmt.Println(invertido)
}
