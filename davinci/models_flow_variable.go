package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = FlowVariable{}
)

type _FlowVariable FlowVariable
type FlowVariable struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmapped"` // used to capture all other properties that are not explicitly defined in the model
	CompanyID            *string                `davinci:"companyId,environmentmetadata,omitempty"`
	Context              *string                `davinci:"context,config,omitempty"`
	CreatedDate          *EpochTime             `davinci:"createdDate,versionmetadata,omitempty"`
	CustomerID           *string                `davinci:"customerId,environmentmetadata,omitempty"`
	Fields               *FlowVariableFields    `davinci:"fields,unmapped,omitempty"`
	FlowID               *string                `davinci:"flowId,environmentmetadata,omitempty"`
	Key                  *float64               `davinci:"key,flowmetadata,omitempty"`
	Label                *string                `davinci:"label,config,omitempty"`
	Name                 string                 `davinci:"name,config"`
	Type                 string                 `davinci:"type,config"`
	UpdatedDate          *EpochTime             `davinci:"updatedDate,versionmetadata,omitempty"`
	Value                *string                `davinci:"value,config,omitempty"`
	Visibility           *string                `davinci:"visibility,flowmetadata,omitempty"`
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

// DesignerCuesFields implements DaVinciExportModel.
func (o FlowVariable) DesignerCuesFields() []string {
	return []string{}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o FlowVariable) EnvironmentMetadataFields() []string {
	return []string{
		"CompanyID",
		"CustomerID",
		"FlowID",
	}
}

// FlowConfigFields implements DaVinciExportModel.
func (o FlowVariable) FlowConfigFields() []string {
	return []string{
		"Context",
		"Label",
		"Name",
		"Type",
		"Value",
	}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o FlowVariable) FlowMetadataFields() []string {
	return []string{
		"Key",
		"Visibility",
	}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o FlowVariable) VersionMetadataFields() []string {
	return []string{
		"CreatedDate",
		"UpdatedDate",
	}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o FlowVariable) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
