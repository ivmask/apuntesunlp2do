package main

import (
  "fmt";
  "strings";
  "unicode";
  "os";
  "bufio";
)

func buscarYEditar(b string, s string) string {
  auxSlice := strings.Fields(s)
  // Preparo slice de busqueda para la palabra y parametro en minuscula
  auxSliceLowercase := strings.Fields(strings.ToLower(s)) 
  lowercaseParameter := strings.ToLower(b)
  lowercaseParameter = lowercaseParameter[:len(lowercaseParameter)-1]
  // Loopeo
  for i := range len(auxSliceLowercase) {
    if (auxSliceLowercase[i] == lowercaseParameter) {
      // runa alternativa para trabajar mi palabra
      var auxRuneWord = []rune(auxSlice[i])
      for j := range len(auxRuneWord) {
        // Si es mayus, paso a minus, y viceversa
        if (unicode.IsUpper(auxRuneWord[j])) {
          auxRuneWord[j] = unicode.ToLower(auxRuneWord[j])
        } else {
            auxRuneWord[j] = unicode.ToUpper(auxRuneWord[j])
        }
        // Reemplazo el string modificado en mi slice original
        finishedString := string(auxRuneWord)
        auxSlice[i] = finishedString
      }
    }
  }
  // Paso slice a string y retorno
  var retorno string = strings.Join(auxSlice, " ")
  return retorno
}

func main() {
  reader := bufio.NewReader(os.Stdin)

  fmt.Println("Ingresa una frase: ")
  s, _ := reader.ReadString('\n')
  fmt.Println("Ingresa una palabra: ")
  k, _ := reader.ReadString('\n')
  
  r := buscarYEditar(k, s)
  fmt.Println(r)
}
