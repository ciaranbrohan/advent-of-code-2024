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
  var difference int
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
  
  difference = 0 
  fmt.Println(difference)
  fmt.Println(odds[0])
  fmt.Println(even[0])


  for i, value := range odds {
    
    diff := value - even[i]
    if diff < 0 {
      diff = -diff
    }
    difference = difference + diff
  } 

  fmt.Println(difference)
}
