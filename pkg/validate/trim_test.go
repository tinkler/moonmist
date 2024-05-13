package validate

import "testing"

// Example struct
type Person struct {
	Name    string
	Age     int
	Address string
}

func TestTrimSpacesStruct(t *testing.T) {
	person := Person{Name: "   John  ", Age: 30, Address: "   123 Main St   "}
	// Before trimming spaces
	println("Before trimming spaces:")
	println("Name:", person.Name)
	println("Age:", person.Age)
	println("Address:", person.Address)
	// Trim spaces from string fields
	TrimStructSpaces(&person)
	// After trimming spaces
	if person.Name != "John" {
		println("Name:", person.Name)
		t.Fail()
	}
	if person.Address != "123 Main St" {
		println("Address:", person.Address)
		t.Fail()
	}
}
