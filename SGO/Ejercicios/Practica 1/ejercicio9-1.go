package main

import (
  "fmt";
  "strings";
  "bufio";
  "unicode";
  "os"
)

func main() {
  fmt.Println("Ingresa una frase para reemplazar las ocurrencias de 'jueves' por 'martes': ")
  reader := bufio.NewReader(os.Stdin)
  s, _ := reader.ReadString('\n')


  auxSlice := strings.Fields(s)
  auxSliceLowercase := strings.Fields(strings.ToLower(s))
  
  for i := range len(auxSliceLowercase) {
//    if (strings.Contains(auxSliceLowercase[i], "jueves")) {
    if (auxSliceLowercase[i] == "jueves") {
      auxWord := auxSlice[i]
      auxMartes := "martes"
      var auxRuneArrayAuxWord = []rune(auxWord)
      var auxRuneArrayMartes = []rune(auxMartes)
      for j := range len(auxRuneArrayAuxWord) {
        if (unicode.IsUpper(auxRuneArrayAuxWord[j])) {
          auxRuneArrayMartes[j] = unicode.ToUpper(auxRuneArrayMartes[j])
      auxMartes := string(auxRuneArrayMartes)
      auxSlice[i] = auxMartes
        }
      }
    }
  }
  retorno := strings.Join(auxSlice, " ")
  fmt.Println(retorno)
}
