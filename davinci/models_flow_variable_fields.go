package davinci

import "encoding/json"

type _FlowVariableFields FlowVariableFields
type FlowVariableFields struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	Type                 string                 `json:"type,omitempty"`
	DisplayName          string                 `json:"displayName,omitempty"`
	Mutable              bool                   `json:"mutable,omitempty"`
	Value                *string                `json:"value,omitempty"`
	Min                  int                    `json:"min,omitempty"`
	Max                  int                    `json:"max,omitempty"`
}

func (o FlowVariableFields) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o FlowVariableFields) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	result["type"] = o.Type
	result["displayName"] = o.DisplayName
	result["mutable"] = o.Mutable

	if o.Value != nil {
		result["value"] = o.Value
	}

	result["min"] = o.Min
	result["max"] = o.Max

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *FlowVariableFields) UnmarshalJSON(bytes []byte) (err error) {
	varFlowVariableFields := _FlowVariableFields{}

	if err = json.Unmarshal(bytes, &varFlowVariableFields); err == nil {
		*o = FlowVariableFields(varFlowVariableFields)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "mutable")
		delete(additionalProperties, "value")
		delete(additionalProperties, "min")
		delete(additionalProperties, "max")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
