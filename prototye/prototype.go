package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

type Address struct {
	Street   string
	Province string
	Country  string
}

func (s *Address) DeepCopy() *Address {
	q := *s
	return &q
}

func (s *Address) DeepCopyBySerialization() *Address {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(s)

	d := gob.NewDecoder(&b)
	result := Address{}
	d.Decode(&result)
	return &result
}

type Person struct {
	Name    string
	Age     int
	Birth   string
	Address *Address
}

func (s *Person) DeepCopy() *Person {
	q := *s
	q.Address = s.Address.DeepCopy()
	return &q
}
func (s *Person) DeepCopyBySerialization() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(s)

	d := gob.NewDecoder(&b)
	result := Person{}
	d.Decode(&result)
	return &result
}

func main() {
	person := &Person{
		Name:  "Hiro",
		Age:   22,
		Birth: "15/11/2002",
		Address: &Address{
			Street:   "Thach Khoi",
			Province: "Hai Duong",
			Country:  "Viet Nam",
		},
	}

	if v, _ := json.Marshal(person.DeepCopyBySerialization()); person != person.DeepCopyBySerialization() {
		fmt.Println("Copy Success", string(v))
	} else {
		fmt.Println("False copy")
	}
}
