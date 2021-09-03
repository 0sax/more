package more

import (
"fmt"
"reflect"
"strings"
)

type MoreData struct {
	More map[string]*ExtensionData
}

type ExtensionData struct {
	Data      interface{}
	ValueType string
}

func (md *MoreData) Add(
	key string, value interface{}) error {

	if value == nil {
		fmt.Errorf("nil value")
		// Add new
	}
	if key == "il" {
		fmt.Errorf("no key provided")
		// Add new
	}

	// convert key to lowecase
	keyLower := strings.ToLower(key)

	// Check if Map has not been initialised
	if md.More == nil {
		md.More = make(map[string]*ExtensionData)
	}

	// Check if item doesn't exist
	if _, ok := md.More[keyLower]; !ok {
		// create empty struct
		md.More[keyLower] = &ExtensionData{}
	}

	// update item
	t := reflect.TypeOf(value)               // get type
	md.More[keyLower].Data = value           //update value
	md.More[keyLower].ValueType = t.String() //update type

	return nil

}

func (md *MoreData) Get(key string) (v interface{}, vType string, err error) {

	if md.More == nil {
		err = fmt.Errorf("no additional data in this object")
		return
	}

	if key == "" {
		err = fmt.Errorf("no key entered")
		return
	}

	keyLower := strings.ToLower(key)

	if _, ok := md.More[keyLower]; !ok {
		err = fmt.Errorf("no value found for key '%v'", keyLower)
		return
	}

	d := md.More[keyLower]

	v = d.Data
	vType = d.ValueType
	return

}

func (md *MoreData) GetIfString(key string) (string, error) {

	v, vt, err := md.Get(key)
	if err != nil {
		return "", err
	}

	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf(
			"key: '%v' returned value '%v', which is of type %v, and not a string",
			strings.ToLower(key), v, vt)
	}

	return s, nil
}
func (md *MoreData) GetIfInt(key string) (int, error) {

	v, vt, err := md.Get(key)
	if err != nil {
		return 0, err
	}

	i, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf(
			"key: '%v' returned value '%v', which is of type %v, and not an int",
			strings.ToLower(key), v, vt)
	}

	return i, nil
}
