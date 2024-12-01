package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
  "sort"
)

func main () {

  var odds []int
  var even []int
  var similarityScore int
  file,err := os.Open("lists.txt")

  if err != nil {
    fmt.Println("Error opening file:", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan(){
    numbers := strings.Fields(scanner.Text())

    if len(numbers) != 2 {
      continue
    }

    for i, numStr := range numbers {
        
      num, err := strconv.Atoi(numStr)

      if err != nil {
        continue
      }


      if i%2 == 0 {
        odds = append(odds, num)
      } else {
        even = append(even, num)
      }

    }


  }

  sort.Ints(odds)
  sort.Ints(even)
  
  for  i := range odds {
    count := 0

    for j := range even {
      if odds[i] == even[j] {
        count++
      }
    }

    if count > 0 {
      similarityScore += odds[i] * count
    }
    
  } 

  fmt.Println(similarityScore)
}
