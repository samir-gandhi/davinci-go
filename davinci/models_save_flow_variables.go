package davinci

import "encoding/json"

type _SaveFlowVariables SaveFlowVariables
type SaveFlowVariables struct {
	AdditionalProperties map[string]interface{} `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Value                []SaveFlowVariable     `json:"value,omitempty" davinci:"value,*,omitempty"`
}

func (o SaveFlowVariables) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o SaveFlowVariables) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Value != nil {
		result["value"] = o.Value
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *SaveFlowVariables) UnmarshalJSON(bytes []byte) (err error) {
	varSaveFlowVariables := _SaveFlowVariables{}

	if err = json.Unmarshal(bytes, &varSaveFlowVariables); err == nil {
		*o = SaveFlowVariables(varSaveFlowVariables)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
