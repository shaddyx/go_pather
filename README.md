# Go pather

The idea of this library is to provide a simple interface for unstructured data traversal

## How to use

### Path way
the GetPath function receives 2 arguments, the object to traverce and the path:
```golang
func GetPath(obj any, path string) (any, *PatherError) 
```

the path should be splitted with a dot and contain map keys or array indices, for sample:

```golang
testData :=map[string]any{
		"values": []any{
			1, 22, 3, 4, "55",
		},
}
res, err := GetPath(testData, "values.2")
```
will return 22

Example in the test:

```golang
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

```

### functional way

```golang
	testData := map[string]any{
		"values": []any{
			1, 2, 3, 4, "55",
		},
		"b": 123,
	}
	pather := Pather{
		Value: testData,
	}

	res := pather.K("values").I(1).Value

```

where:

- pather.K - select value by key
- pather.I - select value by index