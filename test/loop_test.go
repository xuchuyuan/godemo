package test

import (
	"demo/cases"
	"encoding/json"
	"testing"
)

func BenchmarkForStruct(b *testing.B) {
	person := cases.InitPerson()
	count := len(person)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cases.ForStruct(person, count)
	}
}

func BenchmarkForRangeStruct(b *testing.B) {
	person := cases.InitPerson()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cases.ForRangeStruct(person)
	}
}
func BenchmarkJsonToStruct(b *testing.B) {
	var (
		person       = cases.InitPerson()
		againPersons []cases.AgainPerson
	)
	data, err := json.Marshal(person)
	if err != nil {
		b.Fatalf("json.Marshal err: %v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cases.JsonToStruct(data, againPersons)
	}
}

func BenchmarkJsonIteratorToStruct(b *testing.B) {
	var (
		person       = cases.InitPerson()
		againPersons []cases.AgainPerson
	)
	data, err := json.Marshal(person)
	if err != nil {
		b.Fatalf("json.Marshal err: %v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cases.JsonIteratorToStruct(data, againPersons)
	}
}
