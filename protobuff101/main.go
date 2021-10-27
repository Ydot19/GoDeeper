package main

import (
	"encoding/json"
	"fmt"
	"github.com/Ydot19/GoDeeper/protobuff101/model"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	book := &model.Book{
		Id: 1,
		Title: "The road to wigan pier",
		Authors: []*model.Author{
			{Id: 1, Name: "oscar wilde"},
		},
		Category: model.Category_Spiritual,
	}

	data, err := proto.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("book.protobuf", data, 0600); err != nil {
		log.Fatal(err)
	}

	// comparing with json
	data, err = json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("book.json", data, 0600); err != nil {
		log.Fatal(err)
	}

	// decode
	// creates a byte struct
	data, err = ioutil.ReadFile("book.protobuf")
	if err != nil {
		log.Fatal(err)
	}

	bookFromFile := model.Book{}
	if err := proto.Unmarshal(data, &bookFromFile); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Book from protobuf file %+v\n", bookFromFile)
}