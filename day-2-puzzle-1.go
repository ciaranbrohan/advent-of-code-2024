package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
)

func checkNumber (number int, nextNumber int) (isSafe bool) {

  //normalise number to positive


  diff := number - nextNumber

  if(diff < 0) {
    diff = -diff
  }

  if diff > 1 && diff < 3 {
    fmt.Println("safe")
    return true
  } else {
    return false
  }

}


func reportSafe (line []int) (isSafe bool) {

  if checkLevels(line, true) {
    return true
  } 

  return checkLevels(line, false)
  
}
func main () {


  // import txt of codes
  file, err := os.Open("codes.txt")
  if err != nil {
    fmt.Println("error opening file:", err)
    return
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  var lines [][]int

  for scanner.Scan() {
    line := scanner.Text()
    fields := strings.Fields(line)
    
    var numbers []int
    for _, field := range fields {
      
      num, err := strconv.Atoi(field)
      if err != nil {
        fmt.Println("error converting string to int:", err)
        return
      }

      numbers = append(numbers, num)
    }

    lines = append(lines, numbers)
  }



  safeCounter := 0
  // logic to handle codes#

  // safeCounter := 0
	// for _, line := range lines {
	// 	if isReportSafe(parseLine(line), false) {
	// 		safeCounter++
	// 	}
	// }

  for _, numbers := range lines {

    if reportSafe(numbers) {
      safeCounter++
    }

}

fmt.Println(safeCounter)
}
