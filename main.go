package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	xmlFile, err := os.Open("ets.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	var items EtsItems
	err = xml.Unmarshal(b, &items)
	if err != nil {
		fmt.Println("Error unmarshaling xml:", err)
		return
	}

	fmt.Println("items unmarshaled. count is", len(items.Items))
	fmt.Printf("%+v\n", items.Items[0])
}
