package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	res := getData()

	for _, data := range res.Data {
		fmt.Println(*data)
	}

}

func getData() DataCont {
	res, err := ioutil.ReadFile("./data/data.json")

	if err != nil {
		log.Fatalln(err)
	}

	var container DataCont

	err = json.Unmarshal(res, &container)

	if err != nil {
		log.Fatalln(err)
	}
	return container
}
