package davinci

import "encoding/json"

type _FormObj FormObj
type FormObj struct {
	AdditionalProperties map[string]interface{} `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Value                *string                `json:"value,omitempty" davinci:"value,config,omitempty"`
}

func (o FormObj) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o FormObj) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Value != nil {
		result["value"] = o.Value
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *FormObj) UnmarshalJSON(bytes []byte) (err error) {
	varFormObj := _FormObj{}

	if err = json.Unmarshal(bytes, &varFormObj); err == nil {
		*o = FormObj(varFormObj)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
