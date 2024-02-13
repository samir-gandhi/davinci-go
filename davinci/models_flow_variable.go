package davinci

import "encoding/json"

type _FlowVariable FlowVariable
type FlowVariable struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	CompanyID            string                 `json:"companyId"`
	Context              string                 `json:"context"`
	CreatedDate          int32                  `json:"createdDate"`
	CustomerID           string                 `json:"customerId"`
	Fields               FlowVariableFields     `json:"fields"`
	FlowID               string                 `json:"flowId"`
	Key                  float64                `json:"key"`
	Label                string                 `json:"label"`
	Name                 string                 `json:"name"`
	Type                 string                 `json:"type"`
	UpdatedDate          *int32                 `json:"updatedDate,omitempty"`
	Value                string                 `json:"value"`
	Visibility           string                 `json:"visibility"`
}

func (o FlowVariable) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o FlowVariable) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	result["companyId"] = o.CompanyID
	result["context"] = o.Context
	result["createdDate"] = o.CreatedDate
	result["customerId"] = o.CustomerID
	result["fields"] = o.Fields
	result["flowId"] = o.FlowID
	result["key"] = o.Key
	result["label"] = o.Label
	result["name"] = o.Name
	result["type"] = o.Type

	if o.UpdatedDate != nil {
		result["updatedDate"] = o.UpdatedDate
	}

	result["value"] = o.Value
	result["visibility"] = o.Visibility

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
		delete(additionalProperties, "companyId")
		delete(additionalProperties, "context")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "flowId")
		delete(additionalProperties, "key")
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updatedDate")
		delete(additionalProperties, "value")
		delete(additionalProperties, "visibility")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
