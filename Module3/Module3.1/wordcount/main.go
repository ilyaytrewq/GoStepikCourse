package main

import (
	"os"
	"fmt"

	"github.com/ilyaytrewq/GoStepikCourse/tree/master/Module3/Module3.1/wordcount/wordcnt"
)

func main() {
	s := os.Args[1]
	fmt.Println(wordcnt.WordCount(s))
}
