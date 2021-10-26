package cases

import (
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

//test: for / range / [json.Marshal/Unmarshal]
type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Avatar string `json:"avatar"`
	Type   string `json:"type"`
}

type AgainPerson struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Avatar string `json:"avatar"`
	Type   string `json:"type"`
}

const MAX = 10000

func InitPerson() []Person {
	var persons []Person
	for i := 0; i < MAX; i++ {
		persons = append(persons, Person{
			Name:   "xuchuyuan",
			Age:    i,
			Avatar: "https://github.com/xuchuyuan",
			Type:   "Person",
		})
	}

	return persons
}

func ForStruct(p []Person, count int) {
	for i := 0; i < count; i++ {
		_, _ = i, p[i]
	}
}

func ForRangeStruct(p []Person) {
	for i, v := range p {
		_, _ = i, v
	}
}

func JsonToStruct(data []byte, againPerson []AgainPerson) ([]AgainPerson, error) {
	err := json.Unmarshal(data, &againPerson)
	return againPerson, err
}

func JsonIteratorToStruct(data []byte, againPerson []AgainPerson) ([]AgainPerson, error) {
	var jsonIter = jsoniter.ConfigCompatibleWithStandardLibrary
	err := jsonIter.Unmarshal(data, &againPerson)
	return againPerson, err
}
