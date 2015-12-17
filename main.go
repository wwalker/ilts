package main // Author wwalker
import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	timeFormat := "2006-01-02T15:04:05.000000"
	printfFormat := "%s - %s\n"
	if len(os.Args) > 1 {
		timeFormat = os.Args[1]
	}
	if len(os.Args) > 2 {
		printfFormat = os.Args[2]
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		time := time.Now()
		fmt.Printf(printfFormat, time.Format(timeFormat), line)
	}
}
