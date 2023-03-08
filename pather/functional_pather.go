package pather

import (
	"fmt"
	"reflect"
)

type Pather struct {
	Value  any
	parent *Pather
	Path   any
	err    error
}

func (pather *Pather) generateError(key any, err error) *Pather {
	return &Pather{
		parent: pather,
		Value:  nil,
		err:    err,
		Path:   key,
	}
}

func (pather *Pather) K(key string) *Pather {
	if pather.err != nil {
		return pather.generateError(key, pather.err)
	}
	m, ok := pather.Value.(map[string]any)
	if !ok {
		return pather.generateError(key, fmt.Errorf("map expected, %s given", reflect.TypeOf(pather.Value).Name()))
	}
	val, ok := m[key]
	if !ok {
		return pather.generateError(key, fmt.Errorf("key not found: %s", key))
	}
	return &Pather{
		parent: pather,
		Value:  val,
		Path:   key,
		err:    nil,
	}
}

func (pather *Pather) I(key int) *Pather {
	if pather.err != nil {
		return pather.generateError(key, pather.err)
	}
	m, ok := pather.Value.([]any)
	if !ok {
		return pather.generateError(key, fmt.Errorf("list expected, %s given", reflect.TypeOf(pather.Value).Name()))
	}
	if len(m) <= key {
		return pather.generateError(key, fmt.Errorf("no such index: %d", key))
	}
	return &Pather{
		parent: pather,
		Value:  m[key],
		err:    nil,
		Path:   key,
	}
}

func (pather *Pather) AsInt() (int, error) {
	if pather.parent.err != nil {
		return 0, pather.parent.err
	}
	val, ok := pather.Value.(int)
	if !ok {
		return 0, fmt.Errorf("error converting value: %v, from %v, of type %s", pather.Value, pather.Value, reflect.ValueOf(pather.Value).Type().Name())
	}
	return val, nil
}

func (pather *Pather) AsString() (string, error) {
	if pather.parent.err != nil {
		return "", pather.parent.err
	}
	val, ok := pather.Value.(string)
	if !ok {
		return "", fmt.Errorf("error converting value: %v", pather.Value)
	}
	return val, nil
}

func (pather *Pather) AsFloat64() (float64, error) {
	if pather.parent.err != nil {
		return 0, pather.parent.err
	}
	val, ok := pather.Value.(float64)
	if !ok {
		return 0, fmt.Errorf("error converting value: %v", pather.Value)
	}
	return val, nil
}
func (pather *Pather) AsBool() (bool, error) {
	if pather.parent.err != nil {
		return false, pather.parent.err
	}
	val, ok := pather.Value.(bool)
	if !ok {
		return false, fmt.Errorf("error converting value: %v", pather.Value)
	}
	return val, nil
}
