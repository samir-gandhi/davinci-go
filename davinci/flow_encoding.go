package davinci

import (
	"encoding/json"
	"log"
	"reflect"
	"strings"
)

type MarshalOpts ExportCmpOpts
type UnmarshalOpts ExportCmpOpts

func Unmarshal(bytes []byte, v any, opts ExportCmpOpts) (err error) {

	_, ok := v.(DaVinciExportModel)
	log.Printf("HERE!!!1 %s, %v", reflect.TypeOf(v).String(), ok)
	ok = reflect.TypeOf(v).Elem().Implements(reflect.TypeOf((*DaVinciExportModel)(nil)).Elem())
	log.Printf("HERE!!!1 %s, %v", reflect.TypeOf(v).Elem(), ok)

	// Test if v implements the DaVinciExportModel interface.  If it does not, let's just use json encoding.  Expected to be used for base types and custom types that don't have special meaning.
	if ok := reflect.TypeOf(v).Elem().Implements(reflect.TypeOf((*DaVinciExportModel)(nil)).Elem()); !ok {
		log.Printf("HERE!!!2")
		return json.Unmarshal(bytes, v)
	}

	log.Printf("HERE!!!3")
	// Unmarshal the data into a map[string]interface{} to work with.
	var tempMap map[string]interface{}
	if err := json.Unmarshal(bytes, &tempMap); err != nil {
		return err
	}

	// Use reflection to iterate over the fields of the struct
	val := reflect.ValueOf(v).Elem()
	typ := val.Type()

	mappedFields := make([]string, 0)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		tagValue := fieldType.Tag.Get("davinci")

		// Split the tag value into JSON field name and purpose
		tagParts := strings.Split(tagValue, ",")
		if len(tagParts) < 2 {
			continue // Skip fields without a proper tag
		}

		jsonFieldName, fieldPurpose := tagParts[0], tagParts[1]

		mappedFields = append(mappedFields, jsonFieldName)

		if (fieldPurpose == "designercue" && !opts.IgnoreDesignerCues) ||
			(fieldPurpose == "environmentmetadata" && !opts.IgnoreEnvironmentMetadata) ||
			(fieldPurpose == "config" && !opts.IgnoreConfig) ||
			(fieldPurpose == "flowmetadata" && !opts.IgnoreFlowMetadata) ||
			(fieldPurpose == "versionmetadata" && !opts.IgnoreVersionMetadata) ||
			fieldPurpose == "unmapped" {

			if jsonValue, ok := tempMap[jsonFieldName]; ok {
				// Convert the map value to a JSON byte slice
				jsonValueBytes, err := json.Marshal(jsonValue)
				if err != nil {
					return err
				}

				// Unmarshal the value into the struct field
				if err := Unmarshal(jsonValueBytes, field.Addr().Interface(), opts); err != nil {
					return err
				}
			}

			continue
		}
	}

	log.Printf("HERE!!!4")
	// Deal with additional / unmapped properties
	if !opts.IgnoreUnmappedProperties {
		for _, mappedField := range mappedFields {
			delete(tempMap, mappedField)
		}

		val.FieldByName("AdditionalProperties").Set(reflect.ValueOf(tempMap))
	}

	return nil
}

func Marshal(v any, opts ExportCmpOpts) ([]byte, error) {
	return json.Marshal(v)
}
