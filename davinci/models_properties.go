package davinci

import "encoding/json"

type _Properties Properties
type Properties struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	Form                 *string                `json:"form,omitempty"`
	SubFlowID            *SubFlowID             `json:"subFlowId,omitempty"`
	SubFlowVersionID     *SubFlowVersionID      `json:"subFlowVersionId,omitempty"`
	SaveFlowVariables    *SaveFlowVariables     `json:"saveFlowVariables,omitempty"`
}

func (o Properties) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Properties) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Form != nil {
		result["form"] = o.Form
	}

	if o.SubFlowID != nil {
		result["subFlowId"] = o.SubFlowID
	}

	if o.SubFlowVersionID != nil {
		result["subFlowVersionId"] = o.SubFlowVersionID
	}

	if o.SaveFlowVariables != nil {
		result["saveFlowVariables"] = o.SaveFlowVariables
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Properties) UnmarshalJSON(bytes []byte) (err error) {
	varProperties := _Properties{}

	if err = json.Unmarshal(bytes, &varProperties); err == nil {
		*o = Properties(varProperties)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "form")
		delete(additionalProperties, "subFlowId")
		delete(additionalProperties, "subFlowVersionId")
		delete(additionalProperties, "saveFlowVariables")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
