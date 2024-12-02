package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
)

func checkLevels (line []int, increasing bool) (isSafe bool) {

  prevNumber := line[0]

  for _, level := range line[1:] {
  
  	var diff int

		if increasing {
			diff = level - prevNumber
		} else {
			diff = prevNumber - level
		}

    if diff < 1 || diff > 3 {
      return false
    }

    prevNumber = level
  }

  return true
}


func reportSafe (line []int) (isSafe bool) {

  if checkLevels(line, true) || checkLevels(line, false) {
    return true
  } 

  for i := 0; i < len(line); i++ {
  
    tempLine := make([]int, 0, len(line)-1)
    tempLine = append(tempLine, line[:i]...)
    tempLine = append(tempLine, line[i+1:]...)

    if checkLevels(tempLine, true) || checkLevels(tempLine, false) {
      return true
    }
  }
  return false
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
