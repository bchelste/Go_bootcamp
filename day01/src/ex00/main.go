package main

import (

	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"flag"
)

type RecipesJSON struct {
	Data		Recipes			
}

type RecipesXML struct {
	Data		Recipes	
}

type Recipes struct {
	Cakes		[]Cake			`json:"cake" xml:"cake"`
}

type Cake struct {
	Name		string			`json:"name" xml:"name"`
	Time		string			`json:"time" xml:"stovetime"`
	Ingredients	[]Ingridients	`json:"ingredients" xml:"ingredients>item"`
}

type Ingridients struct {
	Name		string			`json:"ingredient_name" xml:"itemname"`
	Count		string			`json:"ingredient_count" xml:"itemcount"`
	Unit		string			`json:"ingredient_unit,omitempty" xml:"itemunit"`
}

type DBReader interface {
	Reading(filename string)
	Output()
}

func (j *RecipesJSON) Reading(filename string) {

	f, err := os.Open(filename)
		if (err != nil) {
			fmt.Println("something wrong with opening the file")
			os.Exit(1)
		}
		defer f.Close()
		dec := json.NewDecoder(f)
		dec.Decode(&j.Data)
}

func (x *RecipesXML) Reading(filename string) {

	f, err := os.Open(filename)
		if (err != nil) {
			fmt.Println("something wrong with opening the file")
			os.Exit(1)
		}
		defer f.Close()
		dec := xml.NewDecoder(f)
		dec.Decode(&x.Data)
		
}

func (j *RecipesJSON) Output() {
	output, err := xml.MarshalIndent(j.Data, "", "    ")
		if (err != nil) {
			fmt.Println("something wroong with json output!")
		}
		fmt.Println(string(output))
}

func (x *RecipesXML) Output() {
	output, err := json.MarshalIndent(x.Data, "", " ")
		if (err != nil) {
			fmt.Println("something wroong with xml output!")
		}
		fmt.Println(string(output))
}

func checkTheFormat(fileName string) (string, bool) {

	extension := filepath.Ext(fileName)
	if ((extension == ".json") || (extension == ".xml")) {
		return (extension), (true);
	}

	return (extension), (false);

}

func main() {

	var fileName *string
	fileName = flag.String("f", "", "filename of arg")
	flag.Parse()
	
	extention, result := checkTheFormat(*fileName)
	if (!result) {
		fmt.Println("Error! The wrong file format!")
		return 
	}

	var tmp DBReader

	if (extention == ".json") {
		tmp = new(RecipesJSON)
	} else if (extention == ".xml") {
		tmp = new(RecipesXML)
	}
	
	tmp.Reading(*fileName)
	tmp.Output()

}
