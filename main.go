package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	res := getData()

	for _, data := range res.Data {
		data.Content = getContent(data.Link)
	}

	writeF(res)
}

func writeF(cont DataCont) {
	res, err := json.Marshal(cont)

	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile("./data/newdata.json", res, 0777)
}

func getContent(path string) string {
	output, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln(err)
	}

	return string(output)
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
