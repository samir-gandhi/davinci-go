package davinci

import "encoding/json"

type _FlowVariable FlowVariable
type FlowVariable struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	CompanyID            *string                `json:"companyId,omitempty"`
	Context              *string                `json:"context,omitempty"`
	CreatedDate          *EpochTime             `json:"createdDate,omitempty"`
	CustomerID           *string                `json:"customerId,omitempty"`
	Fields               *FlowVariableFields    `json:"fields,omitempty"`
	FlowID               *string                `json:"flowId,omitempty"`
	Key                  *float64               `json:"key,omitempty"`
	Label                *string                `json:"label,omitempty"`
	Name                 string                 `json:"name"`
	Type                 string                 `json:"type"`
	UpdatedDate          *EpochTime             `json:"updatedDate,omitempty"`
	Value                *string                `json:"value,omitempty"`
	Visibility           *string                `json:"visibility,omitempty"`
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

	if o.CompanyID != nil {
		result["companyId"] = o.CompanyID
	}

	if o.Context != nil {
		result["context"] = o.Context
	}

	if o.CreatedDate != nil {
		result["createdDate"] = o.CreatedDate
	}

	if o.CustomerID != nil {
		result["customerId"] = o.CustomerID
	}

	if o.Fields != nil {
		result["fields"] = o.Fields
	}

	if o.FlowID != nil {
		result["flowId"] = o.FlowID
	}

	if o.Key != nil {
		result["key"] = o.Key
	}

	if o.Label != nil {
		result["label"] = o.Label
	}

	result["name"] = o.Name
	result["type"] = o.Type

	if o.UpdatedDate != nil {
		result["updatedDate"] = o.UpdatedDate
	}

	if o.Value != nil {
		result["value"] = o.Value
	}

	if o.Visibility != nil {
		result["visibility"] = o.Visibility
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
