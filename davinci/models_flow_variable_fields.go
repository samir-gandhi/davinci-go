package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = FlowVariableFields{}
)

type _FlowVariableFields FlowVariableFields
type FlowVariableFields struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmapped"` // used to capture all other properties that are not explicitly defined in the model
	Type                 *string                `davinci:"type,config,omitempty"`
	DisplayName          *string                `davinci:"displayName,config,omitempty"`
	Mutable              *bool                  `davinci:"mutable,config,omitempty"`
	Value                *string                `davinci:"value,config,omitempty"`
	Min                  *int32                 `davinci:"min,config,omitempty"`
	Max                  *int32                 `davinci:"max,config,omitempty"`
}

func (o FlowVariableFields) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o FlowVariableFields) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Type != nil {
		result["type"] = o.Type
	}

	if o.DisplayName != nil {
		result["displayName"] = o.DisplayName
	}

	if o.Mutable != nil {
		result["mutable"] = o.Mutable
	}

	if o.Value != nil {
		result["value"] = o.Value
	}

	if o.Min != nil {
		result["min"] = o.Min
	}

	if o.Max != nil {
		result["max"] = o.Max
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *FlowVariableFields) UnmarshalJSON(bytes []byte) (err error) {
	varFlowVariableFields := _FlowVariableFields{}

	if err = json.Unmarshal(bytes, &varFlowVariableFields); err == nil {
		*o = FlowVariableFields(varFlowVariableFields)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "mutable")
		delete(additionalProperties, "value")
		delete(additionalProperties, "min")
		delete(additionalProperties, "max")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// DesignerCuesFields implements DaVinciExportModel.
func (o FlowVariableFields) DesignerCuesFields() []string {
	return []string{}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o FlowVariableFields) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o FlowVariableFields) FlowConfigFields() []string {
	return []string{
		"Type",
		"DisplayName",
		"Mutable",
		"Value",
		"Min",
		"Max",
	}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o FlowVariableFields) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o FlowVariableFields) VersionMetadataFields() []string {
	return []string{}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o FlowVariableFields) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
