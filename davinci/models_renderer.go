package davinci

import "encoding/json"

type _Renderer Renderer
type Renderer struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	Name                 string                 `json:"name,omitempty"`
}

func (o Renderer) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Renderer) ToMap() (map[string]interface{}, error) {

	// Marshal and unmarshal the metadata
	jsonRenderer, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err = json.Unmarshal(jsonRenderer, &result); err != nil {
		return nil, err
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Renderer) UnmarshalJSON(bytes []byte) (err error) {
	varRenderer := _Renderer{}

	if err = json.Unmarshal(bytes, &varRenderer); err == nil {
		*o = Renderer(varRenderer)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
