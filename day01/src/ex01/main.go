package main

import (
	"fmt"
	"flag"
	"path/filepath"
	"os"
	"encoding/json"
	"encoding/xml"
	"reflect"
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
	GetDB() (Recipes)
}

func (j *RecipesJSON) GetDB() (Recipes){
	return (j.Data)
}

func (x *RecipesXML) GetDB() (Recipes){
	return (x.Data)
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

func findCake(cakes []Cake, victim string) (Cake, bool) {
	var result bool
	var position Cake
	for _, item := range cakes {
		if item.Name == victim {
			result = true
			position = item
		}
	}
	return position, result
}

func checkCakes(reason string, first []Cake, second []Cake) {
	for i := 0; i < len(second); i++ {
		_, res := findCake(first, second[i].Name)
		if (!res) {
			fmt.Printf("%s cake \"%s\"\n", reason, second[i].Name)
		}
	}
}

func checkChangedTime(first []Cake, second []Cake) {
	for i := 0; i < len(second); i++ {
		pos, res := findCake(first, second[i].Name)
		if (res) {
			if (pos.Time != second[i].Time) {
				fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", second[i].Name, second[i].Time, pos.Time)
			}
		}
	}
}

func findIgredient(ingredients []Ingridients, victim string) (Ingridients, bool) {
	var result bool
	var position Ingridients
	for _, item := range ingredients {
		if item.Name == victim {
			result = true
			position = item
		}
	}
	return position, result
}

func checkIngredients(reason string, first []Cake, second []Cake) {
	for i := 0; i < len(second); i++ {
		pos, res := findCake(first, second[i].Name)
		if (res) {
			for j := 0; j < len(second[i].Ingredients); j++ {
				_, check := findIgredient(pos.Ingredients, second[i].Ingredients[j].Name)
				if (!check) {
					fmt.Printf("%s ingredient \"%s\" for cake \"%s\"\n", reason, second[i].Ingredients[j].Name, second[i].Name)
				}
			}
		}
	}
}

func checkChangedIngredients(first []Cake, second []Cake) {
	for i := 0; i < len(second); i++ {
		pos, res := findCake(first, second[i].Name)
		if (res) {
			for j := 0; j < len(second[i].Ingredients); j++ {
				ingr, check := findIgredient(pos.Ingredients, second[i].Ingredients[j].Name)
				if (check) {
					if (ingr.Unit != second[i].Ingredients[j].Unit) && (ingr.Unit != "") && (second[i].Ingredients[j].Unit != "")  {
						fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
							ingr.Name, pos.Name, second[i].Ingredients[j].Unit, ingr.Unit)
						continue
					}
					if (ingr.Unit != "") && (second[i].Ingredients[j].Unit == "") {
						fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
							ingr.Unit, ingr.Name, pos.Name)
						continue
					}
					if (ingr.Unit == "") && (second[i].Ingredients[j].Unit != "") {
						fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
							second[i].Ingredients[j].Unit, ingr.Name, pos.Name)
						continue
					}
					if (ingr.Count != second[i].Ingredients[j].Count) {
						fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
							ingr.Name, pos.Name, second[i].Ingredients[j].Count, ingr.Count)
					}
					

				}
			}
		}
	}
}

func compareDB(oldR Recipes, newR Recipes) {

	if (reflect.DeepEqual(oldR, newR)) {
		fmt.Println("EQUAL DB")
		return
	}

	checkCakes("ADDED", oldR.Cakes, newR.Cakes)
	checkCakes("REMOVED", newR.Cakes, oldR.Cakes)

	checkChangedTime(oldR.Cakes, newR.Cakes)

	checkIngredients("ADDED", oldR.Cakes, newR.Cakes)
	checkIngredients("REMOVED", newR.Cakes, oldR.Cakes)

	checkChangedIngredients(oldR.Cakes, newR.Cakes)
}

func main() {
	
	oldFilename := flag.String("old", "", "file with OLD DB")
	newFilename := flag.String("new", "", "file with NEW DB")
	flag.Parse()

	extentionOld, resultOld := checkTheFormat(*oldFilename)
	if (!resultOld) {
		fmt.Println("Error! The wrong file format with OLD DB!")
		return 
	}
	extentionNew, resultNew := checkTheFormat(*newFilename)
	if (!resultNew) {
		fmt.Println("Error! The wrong file format with NEW DB!")
		return 
	}

	var oldBD DBReader
	if (extentionOld == ".json") {
		oldBD= new(RecipesJSON)
	} else if (extentionOld == ".xml") {
		oldBD = new(RecipesXML)
	}
	oldBD.Reading(*oldFilename)

	var newBD DBReader
	if (extentionNew == ".json") {
		newBD = new(RecipesJSON)
	} else if (extentionNew == ".xml") {
		newBD = new(RecipesXML)
	}
	newBD.Reading(*newFilename)

	compareDB(oldBD.GetDB(), newBD.GetDB())

}
