package pather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPather_getValues(t *testing.T) {
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
	assert.Equal(t, 2, res)
	res = pather.K("values").I(32).Value
	assert.Equal(t, nil, res)
	assert.NotEqual(t, nil, pather.K("values").I(32).err)

	res, err := pather.K("values").I(1).AsInt()
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, res)
}

func TestPather_getValues_err(t *testing.T) {
	testData := map[string]any{
		"values": []any{
			1, 2, 3, 4, "55",
		},
		"b": 123,
	}
	pather := Pather{
		Value: testData,
	}
	err := pather.K("valuesss").
		K("values").
		I(32).err
	assert.NotEqual(t, nil, err)
}
