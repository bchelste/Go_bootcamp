package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func checkFlags(linesFlag *bool, charactersFlag *bool, wordsFlag *bool) bool {
	
	if (!*linesFlag && !*charactersFlag && !*wordsFlag) {
		*wordsFlag = true
		return true
	} else if ((!*linesFlag && !*charactersFlag && *wordsFlag) ||
		(!*linesFlag && *charactersFlag && !*wordsFlag) ||
		(*linesFlag && !*charactersFlag && !*wordsFlag)) {
			return true
	}
	return false
}

func countLines(fileName string, characters []byte, linesFlag bool) {
	if (linesFlag == false) {
		return
	}
	newLine := strings.Count(string(characters), "\n")
	fmt.Printf("%d\t%s\n", newLine, fileName)
}

func countCharacters(fileName string, characters []byte, charactersFlag bool) {
	if (charactersFlag == false) {
		return
	}
	fmt.Printf("%d\t%s\n", len(characters), fileName)
}

func countWords(fileName string, characters []byte, wordsFlag bool) {

	if (wordsFlag == false) {
		return
	}
	
	var cntr int = 0

	scanner := bufio.NewScanner(strings.NewReader(string(characters)))
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	scanner.Split(bufio.ScanWords)
	// scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		cntr++
	}
	
	fmt.Printf("%d\t%s\n", cntr, fileName)
}

// func countWords(fileName string, characters []byte, wordsFlag bool) {

// 	if (wordsFlag == false) {
// 		return
// 	}

// 	var cntr int = 0
// 	slovo := 0;
// 	for i := 0; i < len(characters); i++ {
// 		if ((characters[i] != ' ' ) && slovo == 0) {
//      		 slovo = 1;
//      		 cntr++;
//     	} else if (characters[i] == ' ') {
// 			slovo = 0;
// 		}

// 	}

// 	newLine := strings.Count(string(characters), "\n")
// 	if (newLine == 1) {
// 		newLine--
// 	} else if (newLine > 1) {
// 		newLine -= 2
// 	}
	
// 	fmt.Printf("%d\t%s\n", cntr + newLine, fileName)
// }

func main() {

	var linesFlag bool
	var charactersFlag bool
	var wordsFlag bool

	flag.BoolVar(&linesFlag, "l", false, "to count lines")
	flag.BoolVar(&charactersFlag, "m", false, "to count characters")
	flag.BoolVar(&wordsFlag, "w", false, "to count words")

	flag.Parse()

	flagErr := checkFlags(&linesFlag, &charactersFlag, &wordsFlag)
	if (flagErr == false) {
		log.Fatalln("wrong combination of flag! only ONE FLAG should be put")
		return 
	}

	fileStorage := flag.Args()
	if (len(fileStorage) < 1) {
		log.Fatalln("FILENAME should be put as an arg")
	}

	// fmt.Println(linesFlag)
	// fmt.Println(charactersFlag)
	// fmt.Println(wordsFlag)


	var wg sync.WaitGroup

	for _, fileName := range fileStorage {
		wg.Add(1)

		go func(fileToCount string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			characters, err := os.ReadFile(fileToCount)
			if (err != nil) {
				log.Fatalf("something wrong with reading \"%s\"", fileToCount)
			}

			countLines(fileToCount, characters, linesFlag)
			countCharacters(fileToCount, characters, charactersFlag)
			countWords(fileToCount, characters, wordsFlag)


		}(fileName)
	}
	// Wait for all to complete.
	wg.Wait()

}