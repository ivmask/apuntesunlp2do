package main

import(
  "fmt"
  "strings"
  "unicode"
  "bufio"
  "os"
)

func reemplazarOcurrencia (frase string) string {
  auxSliceLower := strings.Fields(strings.ToLower(frase))
  auxSlice := strings.Fields(frase)
  for i := range len(auxSliceLower) {
  if strings.Compare(auxSliceLower[i-1], "jueves") {
      auxWord := auxSlice[i-1]
  auxMartes := "martes"
      for j := range len(auxWord) {
        if (strings.IsUpper(auxWord[j-1]) == true) {
          strings.ToUpper(auxMartes[j-1])
        }
      }
      auxSlice[i-1] = auxMartes
    }
  }
  retorno := strings.Join(auxSlice, " ")
  return retorno
}

func main() {
  var string a 
  scanner := bufio.NewScanner(os.Stdin)
  conversion = reemplazarOcurrencia(a)
  fmt.Println(conversion)
}
