package pather

import (
	"testing"
)

func getTestData() map[string]any {
	return map[string]any{
		"values": []any{
			1, 2, 3, 4, "55",
		},
		"mapValues": map[string]any{
			"a": 1,
			"b": 2,
			"nestedList": []any{
				1, 2, 3,
			},
		},
		"b": 123,
	}
}
func TestGetPath(t *testing.T) {
	testData := getTestData()
	res, err := GetPath(testData, "values.2")
	if err != nil {
		t.Fail()
	}
	if res != 3 {
		t.Fail()
	}
	res, _ = GetPath(testData, "values.4")
	if res != "55" {
		t.Fail()
	}
}
func TestGetPathWithNestedMap(t *testing.T) {
	testData := getTestData()
	res, err := GetPath(testData, "mapValues.a")
	if err != nil {
		t.Fail()
	}
	if res != 1 {
		t.Fail()
	}

}
