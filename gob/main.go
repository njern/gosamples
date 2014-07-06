package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
)

const filename = "Gob.gob"

type Person struct {
	Name string
	Age  int32
}

func gobToBuffer(data interface{}) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func main() {
	// Let's create a few people
	var people []Person
	people = append(people, Person{Name: "Rob", Age: 57})
	people = append(people, Person{Name: "Ken", Age: 71})
	people = append(people, Person{Name: "Robert", Age: 43})

	// Gob the data
	buffer, err := gobToBuffer(people)
	if err != nil {
		log.Fatalf("Failed to encode data: %v", err)
	}

	// Write the data to a file, with 600 perms (r/w)
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		log.Fatalf("Failed writing data to disk. Error was '%s'", err)
	}

	// Now we can read it back in and put it in a buffer
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	buffer = bytes.NewBuffer(file)

	// Then just put it in a decoder
	decoder := gob.NewDecoder(buffer)

	var samePeople []Person
	err = decoder.Decode(&samePeople)
	if err != nil {
		log.Fatalf("Failed to decode Person slice: %v", err)
	}

	log.Printf("We started out with a Person slice that looks like this: %#v", people)
	log.Printf("After encoding it with gob and decoding it again, it looks like this %#v", samePeople)
}
