package davinci

import "encoding/json"

type _SubFlowVersionID SubFlowVersionID
type SubFlowVersionID struct {
	AdditionalProperties map[string]interface{} `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Value                *SubFlowVersionIDValue `json:"value,omitempty" davinci:"value,*,omitempty"`
}

func (o SubFlowVersionID) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o SubFlowVersionID) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Value != nil {
		result["value"] = o.Value
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *SubFlowVersionID) UnmarshalJSON(bytes []byte) (err error) {
	varSubFlowVersionID := _SubFlowVersionID{}

	if err = json.Unmarshal(bytes, &varSubFlowVersionID); err == nil {
		*o = SubFlowVersionID(varSubFlowVersionID)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
