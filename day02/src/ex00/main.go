package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func checkFlags(symlinksFlag *bool, dirFlag *bool, fileFlag *bool, extentionFlag *string) {
	flag.Visit(func(f *flag.Flag) {
		if (f.Name == "sl") {
			*symlinksFlag = true
		} else if (f.Name == "d") {
			*dirFlag = true
		} else if (f.Name == "f") {
			*fileFlag = true
		}
	})

	if (*symlinksFlag == false) && (*dirFlag == false) && (*fileFlag == false) {
		*symlinksFlag = true
		*dirFlag = true
		*fileFlag = true
		if (*extentionFlag != "") {
			*extentionFlag = ""
		}
	}

	if (*fileFlag == false) && (*extentionFlag != "") {
		*extentionFlag = ""
	}
}

func isHidden(fileToCheck string) (bool) {
	if (fileToCheck[0] == '/') {
		return (fileToCheck[1] == '.')
	}
	return (fileToCheck[0] == '.')
}

func main() {

	var symlinksFlag bool
	var dirFlag bool
	var fileFlag bool
	var extentionFlag string

	flag.BoolVar(&symlinksFlag, "sl", false, "to print only symlinks")
	flag.BoolVar(&dirFlag, "d", false, "to print only directories")
	flag.BoolVar(&fileFlag, "f", false, "to print only files")
	flag.StringVar(&extentionFlag, "ext", "", "only when -f is specified, to print files with certain extentions")
	flag.Parse()
	checkFlags(&symlinksFlag, &dirFlag, &fileFlag, &extentionFlag)
	var pathToFind = flag.Arg(0)

	// fmt.Println(pathToFind)
	// fmt.Println(symlinksFlag)
	// fmt.Println(dirFlag)
	// fmt.Println(fileFlag)
	// fmt.Println(extentionFlag)

	pathFile, err := os.Open(pathToFind)
	if (err != nil) {
		fmt.Println("something wrong with the path")
	}
	if (pathToFind == "") {
		fmt.Println("empty path")
	}
	defer pathFile.Close()
	err = pathFile.Close()
	if (err != nil) {
		log.Fatalln("Error:", err)
		return 
	}

	tmp := filepath.WalkDir(pathToFind, func(path string, entry fs.DirEntry, err error) error {
		if (err != nil) {
			return (nil)
		}
		nextPath := strings.TrimPrefix(path, pathToFind)
		if (nextPath != "") && (!isHidden(nextPath)) {

			itemInfo, errInfo := entry.Info()

			if (errInfo != nil) {
				return nil
			}

			if symlinksFlag && itemInfo.Mode()&fs.ModeSymlink != 0 {
				if link, err := filepath.EvalSymlinks(path); err != nil {
				  fmt.Printf("%s -> [broken]\n", path)
				} else {
				  fmt.Printf("%s -> %s\n", path, link)
				}
				return nil
			}

			if fileFlag && itemInfo.Mode().IsRegular() {
				if extentionFlag != "" {
				  if matchedExt, err := filepath.Match(
					fmt.Sprintf("*.%s", extentionFlag), entry.Name(),
				  ); !matchedExt || err != nil {
					return nil
				  }
				}
				fmt.Println(path)
				return nil
			}

			if dirFlag && entry.IsDir() {
				fmt.Println(path)
				return nil
			}
			  
		}
		return nil
	})

	if (tmp != nil) {
		fmt.Printf("error walking: %v\n", err)
		return
	}

}
