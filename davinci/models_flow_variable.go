package davinci

import "encoding/json"

type _FlowVariable FlowVariable
type FlowVariable struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	Context              string                 `json:"context"`
	CreatedDate          int64                  `json:"createdDate"`
	CustomerID           string                 `json:"customerId"`
	Fields               FlowVariableFields     `json:"fields"`
	FlowID               string                 `json:"flowId"`
	Type                 string                 `json:"type"`
	UpdatedDate          *int64                 `json:"updatedDate,omitempty"`
	Visibility           string                 `json:"visibility"`
	Name                 string                 `json:"name"`
	CompanyID            string                 `json:"companyId"`
}

func (o FlowVariable) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o FlowVariable) ToMap() (map[string]interface{}, error) {

	// Marshal and unmarshal the metadata
	jsonFlowVariable, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err = json.Unmarshal(jsonFlowVariable, &result); err != nil {
		return nil, err
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *FlowVariable) UnmarshalJSON(bytes []byte) (err error) {
	varFlowVariable := _FlowVariable{}

	if err = json.Unmarshal(bytes, &varFlowVariable); err == nil {
		*o = FlowVariable(varFlowVariable)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "context")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "flowId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updatedDate")
		delete(additionalProperties, "visibility")
		delete(additionalProperties, "name")
		delete(additionalProperties, "companyId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
