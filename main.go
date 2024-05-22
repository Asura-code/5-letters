package main

import (
    "fmt"
    "os"
    "strings"
    "slices"
    "math/rand"
)

func trust(a []string, b string) bool {
  found := slices.Contains(a, b)
  if found == true{
    return true
  }
  return false

}
var mask = [5]string{"*", "*", "*", "*", "*"}
func sravnenie(a string, b string)string{
  
  for i := 0; i < len([]rune(a)); i++ {
    if a[i] == b[i]{
      // mask = strings.Replace(mask, string(mask[i]), string(a[i]), 1)
      mask[i] = string(a[i])
    }
  }
  var mask2 string

  for i := 0; i<5; i++{
    mask2 += string(mask[i])
  }
  

  return mask2
} 
var ans string
func if_in(a string, b string) string{
  for i := 0; i < 5; i++{
    if strings.ContainsAny(a,string(b[i])){
        if strings.ContainsAny(ans,string(b[i]))==false{
        ans += string(b[i])
      }
    }
  }
  return ans
}

func main() {

    b, err := os.ReadFile("D:/golangprak/words.txt")
    if err != nil {
      fmt.Println(err)
      return
   }

   words := strings.Fields(string(b))
   randomIndex := rand.Intn(len(words))
   answer := words[randomIndex]
   // fmt.Println(answer)

  var input string
   q := true
    // fmt.Println(trust(words,"there"))
    for i := 0; i < 5 ; i++{
      


    fmt.Println("Введите слово:")
    fmt.Scanln(&input)
    if trust(words, input){
      if input == answer{
        q = false
      fmt.Println("Правильно")
      break
    }else{
      fmt.Println("Слово содержит эти буквы:", if_in(input,answer)," | ", "Маска:", sravnenie(input,answer))
    }
    }else{
      fmt.Println("такого слова нет, попробуйте снова")
      i -= 1
    }
   }
   if q{
   fmt.Println("Вы проиграли. Ответ:", answer)
}
}