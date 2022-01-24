package main

import "testing"

var (
	multiLine = `{
    "firstName": "John",
    "lastName": "Doe"
}`
	singleLine  = `{"firstName":"John","lastName":"Doe"}`
	stringified = `"{\"firstName\":\"John\",\"lastName\":\"Doe\"}"`
)

func TestCompact(t *testing.T) {
	out, err := Compact(multiLine)
	if err != nil {
		t.Error(err)
	}
	if out != singleLine {
		t.Errorf("Expected: %s; Got: %s", singleLine, out)
	}

	out, err = Compact(stringified)
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

	out, err = Prettify(stringified)
	if err != nil {
		t.Error(err)
	}
	if out != multiLine {
		t.Errorf("Expected: %s; Got: %s", multiLine, out)
	}
}

func TestStringify(t *testing.T) {
	out, err := Stringify(singleLine)
	if err != nil {
		t.Error(err)
	}
	if out != stringified {
		t.Errorf("Expected: %s; Got: %s", stringified, out)
	}
}

//func TestPrettify_Stringify(t *testing.T) {
//	out, err := Prettify(singleLine)
//	if err != nil {
//		t.Error(err)
//	}
//	if out != multiLine {
//		t.Errorf("Expected: %s; Got: %s", multiLine, out)
//	}
//
//	out, err = Stringify(out)
//	if err != nil {
//		t.Error(err)
//	}
//	if out != singleLine {
//		t.Errorf("Expected: %s; Got: %s", multiLine, out)
//	}
//}
