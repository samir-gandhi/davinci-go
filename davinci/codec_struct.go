package davinci

import (
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"
)

var _ ValueEncoder = &StructDecoder{}
var _ ValueDecoder = &StructDecoder{}

type StructDecoder struct {
	dCtx *DecoderContext
}

func NewStructDecoder(dCtx *DecoderContext) StructDecoder {
	return StructDecoder{
		dCtx: dCtx,
	}
}

func (StructDecoder) String() string {
	return "davinci.StructDecoder"
}

func (d StructDecoder) DecodeValue(data []byte, v reflect.Value) error {
	if !v.IsValid() || !v.CanSet() || v.Kind() != reflect.Struct {
		return fmt.Errorf("invalid struct value to decode")
	}

	mappedFields := make([]string, 0)

	typ := v.Type()

	// Unmarshal the data into a map[string]interface{} to work with.
	var tempMap map[string]interface{}
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if typ.Field(i).Anonymous {
			// decode the value into the struct field
			if err := d.dCtx.Decode(data, field.Addr().Interface()); err != nil {
				return err
			}

			continue
		}

		fieldType := typ.Field(i)
		tagValue := fieldType.Tag.Get("davinci")

		// Split the tag value into JSON field name and purpose
		tagParts := strings.Split(tagValue, ",")
		if len(tagParts) < 2 {
			continue // Skip fields without a proper tag
		}

		jsonFieldName, fieldPurpose := tagParts[0], tagParts[1]

		mappedFields = append(mappedFields, jsonFieldName)

		if !slices.Contains([]string{"*", "-", "designercue", "environmentmetadata", "config", "flowmetadata", "versionmetadata"}, fieldPurpose) {
			return fmt.Errorf("davinci export field purpose %s is not recognised", fieldPurpose)
		}

		if fieldPurpose == "-" {
			continue
		}

		if (fieldPurpose == "designercue" && !d.dCtx.Opts.IgnoreDesignerCues) ||
			(fieldPurpose == "environmentmetadata" && !d.dCtx.Opts.IgnoreEnvironmentMetadata) ||
			(fieldPurpose == "config" && !d.dCtx.Opts.IgnoreConfig) ||
			(fieldPurpose == "flowmetadata" && !d.dCtx.Opts.IgnoreFlowMetadata) ||
			(fieldPurpose == "versionmetadata" && !d.dCtx.Opts.IgnoreVersionMetadata) ||
			fieldPurpose == "*" {

			if jsonValue, ok := tempMap[jsonFieldName]; ok {
				// Convert the map value to a JSON byte slice
				jsonValueBytes, err := json.Marshal(jsonValue)
				if err != nil {
					return err
				}

				if len(tagParts) > 2 && tagParts[2] == "omitempty" && string(jsonValueBytes) == "null" {
					continue
				}

				// Decode the value into the struct field
				if err := d.dCtx.Decode(jsonValueBytes, field.Addr().Interface()); err != nil {
					return err
				}
			}

			continue
		}
	}

	//Deal with additional / unmapped properties
	if !d.dCtx.Opts.IgnoreUnmappedProperties {
		for _, mappedField := range mappedFields {
			delete(tempMap, mappedField)
		}

		if additionalPropsField := v.FieldByName("AdditionalProperties"); additionalPropsField.IsValid() {
			if additionalPropsField.Kind() == reflect.Map && additionalPropsField.Type().Key().Kind() == reflect.String && additionalPropsField.Type().Elem().Kind() == reflect.Interface {
				additionalPropsField.Set(reflect.ValueOf(tempMap))
			}
		}

	}

	return nil
}

func (d StructDecoder) EncodeValue(data interface{}, v reflect.Value) error {
	return json.Unmarshal(data.([]byte), v.Addr().Interface())
}
