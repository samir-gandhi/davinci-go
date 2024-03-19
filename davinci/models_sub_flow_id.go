package davinci

import "encoding/json"

type _SubFlowID SubFlowID
type SubFlowID struct {
	AdditionalProperties map[string]interface{} `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Value                *SubFlowValue          `json:"value,omitempty" davinci:"value,*,omitempty"`
}

func (o SubFlowID) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o SubFlowID) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Value != nil {
		result["value"] = o.Value
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *SubFlowID) UnmarshalJSON(bytes []byte) (err error) {
	varSubFlowID := _SubFlowID{}

	if err = json.Unmarshal(bytes, &varSubFlowID); err == nil {
		*o = SubFlowID(varSubFlowID)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
