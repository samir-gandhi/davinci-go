package davinci

import "encoding/json"

type _SubFlowProperties SubFlowProperties
type SubFlowProperties struct {
	AdditionalProperties map[string]interface{} `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	SubFlowID            *SubFlowID             `json:"subFlowId,omitempty" davinci:"subFlowId,*,omitempty"`
	SubFlowProperties    *SubFlowProperties     `json:"subFlowVersionId,omitempty" davinci:"subFlowVersionId,*,omitempty"`
}

func (o SubFlowProperties) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o SubFlowProperties) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.SubFlowID != nil {
		result["subFlowId"] = o.SubFlowID
	}

	if o.SubFlowProperties != nil {
		result["subFlowVersionId"] = o.SubFlowProperties
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *SubFlowProperties) UnmarshalJSON(bytes []byte) (err error) {
	varSubFlowProperties := _SubFlowProperties{}

	if err = json.Unmarshal(bytes, &varSubFlowProperties); err == nil {
		*o = SubFlowProperties(varSubFlowProperties)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "subFlowId")
		delete(additionalProperties, "subFlowVersionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
