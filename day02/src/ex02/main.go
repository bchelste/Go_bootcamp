package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

// echo -e "go.mod\ngo.mod" | ./myXargs cat

func main() {

	clinput, err := io.ReadAll(os.Stdin)
	if (err != nil) {
		log.Fatalln("something wron with reading stdin")
		return
	}
	
	argsCL := os.Args[2:]

	inputItems := strings.Split(string(clinput), "\n")

	fmt.Println(inputItems)

	for item := range inputItems {
		currentArg := append(argsCL, inputItems[item])
		cmd := exec.Command(os.Args[1], currentArg...)
		output, err := cmd.CombinedOutput()
		if  err != nil {
			return 
		}
		fmt.Printf("%s", output)
	}
}
