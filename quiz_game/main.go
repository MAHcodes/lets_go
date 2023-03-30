package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
  file := flag.String("f", "problems.csv", "specify problems csv file")
  flag.Parse()

  filename := *file

  fileContent, err := os.Open(filename)

  if err != nil {
    log.Fatalln(err)
  }
  defer fileContent.Close()

  lines, err := csv.NewReader(fileContent).ReadAll()
  if err != nil {
    log.Fatalln(err)
  }

  reader := bufio.NewReader(os.Stdin)
  var score int

  for _, line := range lines {
    question, solution := line[0], line[1]
    fmt.Printf("%s: ", question)
    answer, err := reader.ReadString('\n')

    if err != nil {
      log.Fatalln(err)
    }

    if (strings.Trim(answer, "\n") == solution) {
      score++
    }
  }

  fmt.Printf("You scored %d out of %d\n", score, len(lines))
}
