package main

import (
	"fmt"
	"strings"
	"valyntyler.com/aoc-2024-day-05/input"
)

func main()  {
  var s string = string(input.Bytes)
  var result string = strings.Trim(s, " \n")
  var lines []string = strings.Split(result, "\n")

  for _, line := range lines {
    fmt.Println(line)
  }
}
