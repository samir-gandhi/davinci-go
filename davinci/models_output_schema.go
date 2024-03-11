package davinci

import "encoding/json"

type _OutputSchema OutputSchema
type OutputSchema struct {
	AdditionalProperties map[string]interface{} `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Output               interface{}            `json:"output,omitempty" davinci:"output,config,omitempty"`
}

func (o OutputSchema) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o OutputSchema) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Output != nil {
		result["output"] = o.Output
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *OutputSchema) UnmarshalJSON(bytes []byte) (err error) {
	varOutputSchema := _OutputSchema{}

	if err = json.Unmarshal(bytes, &varOutputSchema); err == nil {
		*o = OutputSchema(varOutputSchema)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "output")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
