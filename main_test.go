package main

import "testing"

var (
	multiLine = `{
    "firstName": "John",
    "lastName": "Doe"
}`
	singleLine = `{"firstName":"John","lastName":"Doe"}`
)

func TestStringify(t *testing.T) {
	out, err := Stringify(multiLine)
	if err != nil {
		t.Error(err)
	}
	if out != singleLine {
		t.Errorf("Expected: %s; Got: %s", singleLine, out)
	}
}

func TestPrettify(t *testing.T) {
	out, err := Prettify(singleLine)
	if err != nil {
		t.Error(err)
	}
	if out != multiLine {
		t.Errorf("Expected: %s; Got: %s", multiLine, out)
	}
}

func TestPrettify_Stringify(t *testing.T) {
	out, err := Prettify(singleLine)
	if err != nil {
		t.Error(err)
	}
	if out != multiLine {
		t.Errorf("Expected: %s; Got: %s", multiLine, out)
	}

	out, err = Stringify(out)
	if err != nil {
		t.Error(err)
	}
	if out != singleLine {
		t.Errorf("Expected: %s; Got: %s", multiLine, out)
	}
}
