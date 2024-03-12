package conv

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// Get a reflect.Value for the struct
	val := reflect.ValueOf(obj)

	// Ensure it's a struct
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, errors.New("obj must be a struct or a pointer to a struct")
	}

	// Iterate through the struct's fields
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// Get the field name
		name := val.Type().Field(i).Name

		// Check if the field is exported (public)
		if field.CanInterface() {
			result[name] = field.Interface()
		}
	}

	return result, nil
}

func MapToStruct(data map[string]any, result any) error {

	// Marshal the map into a JSON byte slice
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling map to JSON: %w", err)
	}

	// Unmarshal the JSON byte slice into the struct
	if err := json.Unmarshal(jsonData, result); err != nil {
		return fmt.Errorf("error unmarshalling JSON to User struct: %w", err)
	}

	return nil
}
