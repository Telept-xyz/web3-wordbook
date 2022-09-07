package validator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error, msg string) {
	if e != nil {
		fmt.Println(string(msg))
		panic(e)
	}
}

func merge(base map[string]interface{}, incremental map[string]interface{}) map[string]interface{} {

	for k, v := range incremental {
		if _, ok := base[k]; ok {
			//there is duplicated key
			check(fmt.Errorf("%s is duplicated"), "Attention")
		}

		base[k] = v
	}

	return base
}

// @title web3-wordbook validator
// @version 1.0
// @description verify the word uniqueness
func main() {

	// verify the word list to make sure no duplicate word
	var wordIndex map[string]interface{}
	dat, err := os.ReadFile("./words.json")
	check(err, "Fail to load words.json")
	fmt.Println(string(dat))
	err = json.Unmarshal([]byte(dat), &wordIndex)
	check(err, "Failed to unmarshal json. The format must be wrong or there is duplicated word")

	fmt.Println("word index check passed")

	wordBooks := make(map[string]interface{})

	files, err := ioutil.ReadDir("./books")
	check(err, "Fail to load books")

	for _, f := range files {
		var book map[string]interface{}
		path := filepath.Join("./books", f.Name())
		dat, err := ioutil.ReadFile(path)
		check(err, "Fail to load "+path)
		fmt.Println(string(dat))
		err = json.Unmarshal([]byte(dat), &book)
		check(err, "Failed to unmarshal "+path+". The format must be wrong or there is duplicated word")

		merge(wordBooks, book)
		fmt.Printf("%s checked \n", f.Name())
	}

	fmt.Println("You change is safe! Please submit your PR")
}
